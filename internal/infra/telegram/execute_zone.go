package telegram

import (
	"context"
	"github.com/bruli/rasberryTelegram/internal/app"
	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func ExecuteZone(ctx context.Context, ch cqs.CommandHandler, chatID int64, msgs *Messages, zone string, seconds int) {
	msg := tgbotapi.NewMessage(chatID, "")
	if _, err := ch.Handle(ctx, app.ExecuteZoneCommand{
		Zone:    zone,
		Seconds: seconds,
	}); err != nil {
		buildMessage(msgs, msg, err.Error())
	}
}
