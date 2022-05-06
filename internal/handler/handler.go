package handler

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/youpps/ruler/internal/controller"
)

type Handler struct {
	bot *tg.BotAPI
}

func NewHandler(bot *tg.BotAPI) *Handler {
	return &Handler{bot}
}

func (h *Handler) HandlerUpdate(controller *controller.Controller) {
	updates := h.getUpdatesChan(h.bot)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			controller.OnCommand(update.Message)
			continue
		}

		if update.Message.Photo != nil {
			controller.OnPhoto(update.Message)
			continue
		}

		if update.Message.Video != nil {
			controller.OnVideo(update.Message)
			continue
		}
		
		controller.OnMessage(update.Message)
	}
}

func (h *Handler) getUpdatesChan(bot *tg.BotAPI) tg.UpdatesChannel {
	return bot.GetUpdatesChan(tg.UpdateConfig{
		Offset:  1,
		Timeout: 60,
	})
}
