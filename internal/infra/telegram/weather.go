package telegram

import (
	"context"
	"fmt"

	"github.com/bruli/rasberryTelegram/internal/app"
	"github.com/bruli/rasberryTelegram/internal/domain/weather"
	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Weather(ctx context.Context, qh cqs.QueryHandler, chatID int64, msgs *Messages) {
	msg := tgbotapi.NewMessage(chatID, "")
	result, err := qh.Handle(ctx, app.WeatherQuery{})
	if err != nil {
		buildMessage(msgs, msg, fmt.Sprintf("failed getting weather: %s", err.Error()))
		return
	}
	weath, _ := result.(weather.Weather)
	buildMessage(msgs, msg, fmt.Sprintf("Current temperature: %v *C", weath.Temperature()))
	buildMessage(msgs, msg, fmt.Sprintf("Current humidity: %v", weath.Humidity()))
	buildMessage(msgs, msg, fmt.Sprintf("Is raining:  %v", weath.IsRaining()))
}
