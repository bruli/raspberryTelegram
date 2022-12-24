package telegram

import (
	"context"
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/app"
	"github.com/bruli/rasberryTelegram/internal/domain/status"
	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Status(ctx context.Context, qh cqs.QueryHandler, chatID int64, msgs *Messages) {
	msg := tgbotapi.NewMessage(chatID, "")
	result, err := qh.Handle(ctx, app.StatusQuery{})
	if err != nil {
		msg.Text = fmt.Sprintf("failed getting status: %s", err.Error())
		return
	}
	st, _ := result.(status.Status)
	buildMessage(msgs, msg, fmt.Sprintf("System started at: %s", st.SystemStartedAt().Date()))
	buildMessage(msgs, msg, fmt.Sprintf("Current temperature: %v *C", st.Temperature()))
	buildMessage(msgs, msg, fmt.Sprintf("Current humidity: %v", st.Humidity()))
	buildMessage(msgs, msg, fmt.Sprintf("Is raining: %v", st.Raining()))
	if st.UpdatedAt() != nil {
		buildMessage(msgs, msg, fmt.Sprintf("System updated at: %s", st.UpdatedAt().Date()))
	}
}

func buildMessage(msgs *Messages, msg tgbotapi.MessageConfig, text string) {
	msg.Text = text
	msgs.AddMessage(msg)
}
