package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/youpps/ruler/pkg/cmd"
)

type Controller struct {
	bot *tg.BotAPI
}

func NewController(bot *tg.BotAPI) *Controller {
	return &Controller{bot}
}

func (c *Controller) OnCommand(msg *tg.Message) {
	switch msg.Command() {
	case "destroy":
		c.BotDestroy(msg)
		return
	case "open_file":
		c.BotOpenFile(msg)
		return
	case "get_file_body":
		c.BotGetFileBody(msg)
		return
	case "ls":
		c.BotLs(msg)
		return
	case "shutdown":
		c.BotShutdownOs(msg)
	}
}

func (c *Controller) OnMessage(message *tg.Message) {
	output, err := cmd.ExecuteCommandWithOutput(message.Text)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}
	msg := tg.NewMessage(message.Chat.ID, string(output))
	c.bot.Send(msg)
}

func (c *Controller) OnPhoto(message *tg.Message) {
	file := message.Photo[len(message.Photo)-1]
	link, err := c.bot.GetFileDirectURL(file.FileID)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	resp, err := http.Get(link)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	filesDir := os.Getenv("FILES_DIRECTORY")

	err = os.MkdirAll(filesDir, os.ModeDir)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	extension := strings.Split(link[len(link)-10:], ".")[1]
	filepath := filesDir + fmt.Sprint(file.FileSize*time.Now().Second()) + "." + extension

	img, err := os.Create(filepath)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}
	defer img.Close()

	img.Write(bytes)

	msg := tg.NewMessage(message.Chat.ID, "Your filepath: "+filepath)
	c.bot.Send(msg)
}

func (c *Controller) OnVideo(message *tg.Message) {
	video := message.Video
	link, err := c.bot.GetFileDirectURL(video.FileID)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}
	resp, err := http.Get(link)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	filesDir := os.Getenv("FILES_DIRECTORY")

	err = os.MkdirAll(filesDir, os.ModeDir)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	extension := strings.Split(video.FileName, ".")[1]
	filepath := filesDir + fmt.Sprint(video.FileSize*time.Now().Second()) + "." + extension

	img, err := os.Create(filepath)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}
	defer img.Close()

	img.Write(bytes)

	msg := tg.NewMessage(message.Chat.ID, "Your filepath: "+filepath)
	c.bot.Send(msg)
}

func (c *Controller) BotLs(message *tg.Message) {
	args := message.CommandArguments()
	dirPath, err := filepath.Abs(args)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	files, err := ioutil.ReadDir(args)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	for _, file := range files {
		msg := tg.NewMessage(message.Chat.ID, dirPath+"\\"+file.Name())
		c.bot.Send(msg)
	}
}

func (c *Controller) BotGetFileBody(message *tg.Message) {
	filepath := message.CommandArguments()
	file, err := os.Open(filepath)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}
	defer file.Close()

	fileInfo, err := ioutil.ReadAll(file)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	msg := tg.NewMessage(message.Chat.ID, string(fileInfo))
	c.bot.Send(msg)
}

func (c *Controller) BotOpenFile(message *tg.Message) {
	filepath := message.CommandArguments()
	if err := cmd.ExecuteCommand(filepath); err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}
	msg := tg.NewMessage(message.Chat.ID, "The file has been opened!")
	c.bot.Send(msg)
}

func (c *Controller) BotDestroy(message *tg.Message) {
	fmt.Println("Bot has started destroying...")

	userDir := os.Getenv("FILES_DIRECTORY")
	os.RemoveAll(userDir)

	msg := tg.NewMessage(message.Chat.ID, "Bot has destroyed all its data.")
	c.bot.Send(msg)

	fmt.Println("Bot has destroyed all its data.")

	os.Exit(0)
}

func (c *Controller) BotShutdownOs(message *tg.Message) {
	if err := cmd.ExecuteCommand("shutdown", "/p"); err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}
	msg := tg.NewMessage(message.Chat.ID, "System has been closed!")
	c.bot.Send(msg)
}
