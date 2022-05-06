package bot

import (
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
	handler.HandlerUpdate(controller)
	callback("Telegram bot has started.")
}
