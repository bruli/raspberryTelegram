package main

import (
	telegram_bot "github.com/bruli/rasberryTelegram/internal/infrastructure/http/telegram-bot"
	"log"
	"os"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	serverUrl := os.Getenv("WATER_SYSTEM_URL")
	conf := telegram_bot.NewConfig(token, serverUrl)
	t := telegram_bot.NewServer(conf)
	if err := t.Run(); err != nil {
		log.Fatal(err)
	}
}
