package main

import (
	telegram_bot "github.com/bruli/rasberryTelegram/internal/infrastructure/http/telegram-bot"
	"log"
	"os"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	serverURL := os.Getenv("WATER_SYSTEM_URL")
	authToken := os.Getenv("AUTH_TOKEN")
	conf := telegram_bot.NewConfig(token, serverURL, authToken)
	t := telegram_bot.NewServer(conf)
	if err := t.Run(); err != nil {
		log.Fatal(err)
	}
}
