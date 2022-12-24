package main

import (
	"context"
	"fmt"
	"github.com/bruli/rasberryTelegram/config"
	telegram2 "github.com/bruli/rasberryTelegram/internal/infra/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
)

const (
	helpCommand = "help"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed reading config")
	}
	ctx, cancel := context.WithCancel(context.Background())
	bot, err := tgbotapi.NewBotAPI(conf.TelegramToken())
	if err != nil {
		log.Fatal().Err(err).Msg("failed starting bot api")
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal().Err(err).Msg("error updating telegram channel")
	}

	/* signal handling */
	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-done
		fmt.Println("signal trapped")
		cancel()
	}()

	msgs := telegram2.NewMessages()
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case helpCommand:
				telegram2.Help(ctx, update.Message.Chat.ID, &msgs)
			}
			for _, j := range msgs.GetMessages() {
				_, _ = bot.Send(j)
			}

			msgs.Clean()
		}
	}
}
