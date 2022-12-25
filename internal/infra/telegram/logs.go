package telegram

import (
	"context"
	"fmt"

	"github.com/bruli/rasberryTelegram/internal/app"
	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Logs(ctx context.Context, qh cqs.QueryHandler, chatID int64, msgs *Messages, number int) {
	msg := tgbotapi.NewMessage(chatID, "")
	result, err := qh.Handle(ctx, app.LogsQuery{Number: number})
	if err != nil {
		buildMessage(msgs, msg, fmt.Sprintf("failed to find logs: %s", err.Error()))
		return
	}
	logs, _ := result.([]string)
	for _, lo := range logs {
		buildMessage(msgs, msg, lo)
	}
}
