package bot

import (
	"os"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/youpps/ruler/internal/controller"
	"github.com/youpps/ruler/internal/handler"
)

type Bot struct {
	bot *tg.BotAPI
}

func NewBot(token string) (*Bot, error) {
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{bot}, nil
}

func (b *Bot) Run(callback func(string)) {
	handler := handler.NewHandler(b.bot)
	controller := controller.NewController(b.bot)
	callback("Telegram bot has been started.")
	handler.HandlerUpdate(controller)
}

func (b *Bot) Destroy(callback func(string)) {
	userDir := os.Getenv("FILES_DIRECTORY")
	os.RemoveAll(userDir)
	callback("Bot has destroyed all its data.")
}
