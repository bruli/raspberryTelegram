package telegram

import (
	"context"

	"github.com/bruli/rasberryTelegram/internal/app"
	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ActivateDeactivate(ctx context.Context, ch cqs.CommandHandler, activate bool, chatID int64, msgs *Messages) {
	if _, err := ch.Handle(ctx, app.ActivateDeactivateServerCmd{Activate: activate}); err != nil {
		msg := tgbotapi.NewMessage(chatID, "")
		buildMessage(msgs, msg, err.Error())
		return
	}
	var txt string
	switch {
	case activate:
		txt = "Activated!!"
	default:
		txt = "Deactivated!!"
	}
	msg := tgbotapi.NewMessage(chatID, txt)
	msgs.AddMessage(msg)
}
