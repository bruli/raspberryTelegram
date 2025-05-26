package telegram

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Help(ctx context.Context, chatID int64, msgs *Messages) {
	select {
	case <-ctx.Done():
		return
	default:
		msg := tgbotapi.NewMessage(chatID, "Type: /weather, /status, /log [limit], /water [zone] [seconds], /activate, /deactivate.")
		msgs.AddMessage(msg)
	}
}
