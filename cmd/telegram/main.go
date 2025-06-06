package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bruli/rasberryTelegram/config"
	"github.com/bruli/rasberryTelegram/internal/app"
	"github.com/bruli/rasberryTelegram/internal/infra/api"
	telegram2 "github.com/bruli/rasberryTelegram/internal/infra/telegram"
	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
)

const (
	statusCommand     = "status"
	weatherCommand    = "weather"
	logCommand        = "log"
	executionCommand  = "water"
	deactivateCommand = "deactivate"
	activateCommand   = "activate"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed reading config")
	}
	client := http.Client{Timeout: 5 * time.Second}
	wsr := api.NewWaterSystemRepository(conf.WsServerURL(), &client, conf.WsServerToken())

	qhErrMdw := cqs.NewQueryHndErrorMiddleware(&log)
	chErrMdw := cqs.NewCommandHndErrorMiddleware(&log)

	statusQh := qhErrMdw(app.NewStatus(wsr))
	weatherQh := qhErrMdw(app.NewWeather(wsr))
	logsQh := qhErrMdw(app.NewLogs(wsr))

	executeCh := chErrMdw(app.NewExecuteZone(wsr))
	actDeactCh := chErrMdw(app.NewActivateDeactivateServer(wsr))

	ctx, cancel := context.WithCancel(context.Background())
	bot, err := tgbotapi.NewBotAPI(conf.TelegramToken())
	if err != nil {
		log.Fatal().Err(err).Msg("failed starting bot api")
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	/* signal handling */
	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-done
		fmt.Println("signal trapped")
		cancel()
	}()

	execute(updates, ctx, statusQh, actDeactCh, weatherQh, logsQh, executeCh, bot)
}

func execute(
	updates tgbotapi.UpdatesChannel,
	ctx context.Context,
	statusQh cqs.QueryHandler,
	actDeactCh cqs.CommandHandler,
	weatherQh, logsQh cqs.QueryHandler,
	executeCh cqs.CommandHandler,
	bot *tgbotapi.BotAPI,
) {
	msgs := telegram2.NewMessages()
	for update := range updates {
		if update.Message == nil {
			continue
		}
		chatID := update.Message.Chat.ID
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			default:
				telegram2.Help(ctx, chatID, &msgs)
			case statusCommand:
				telegram2.Status(ctx, statusQh, chatID, &msgs)
			case activateCommand:
				telegram2.ActivateDeactivate(ctx, actDeactCh, true, chatID, &msgs)
			case deactivateCommand:
				telegram2.ActivateDeactivate(ctx, actDeactCh, false, chatID, &msgs)
			case weatherCommand:
				telegram2.Weather(ctx, weatherQh, chatID, &msgs)
			case logCommand:
				number, err := strconv.Atoi(update.Message.CommandArguments())
				if err != nil {
					number = 2
				}
				telegram2.Logs(ctx, logsQh, chatID, &msgs, number)
			case executionCommand:
				arguments := strings.Fields(update.Message.CommandArguments())
				zone := arguments[0]
				seconds, err := strconv.Atoi(arguments[1])
				if err != nil {
					msgs.AddMessage(tgbotapi.NewMessage(chatID, fmt.Sprintf("invalid seconds number: %s", arguments[1])))
				} else {
					telegram2.ExecuteZone(ctx, executeCh, chatID, &msgs, zone, seconds)
				}
			}
			for _, j := range msgs.GetMessages() {
				_, _ = bot.Send(j)
			}

			msgs.Clean()
		}
	}
}
