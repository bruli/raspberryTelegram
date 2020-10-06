package main

import (
	"flag"
	telegram_bot "github.com/bruli/rasberryTelegram/internal/infrastructure/http/telegram-bot"
	"github.com/spf13/viper"
	"log"
)

func main() {
	configFile := flag.String("config", "", "config file")
	flag.Parse()

	viper.SetConfigFile(*configFile)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("invalid config file: %s", err)
	}
	token := viper.GetString("telegram_token")
	serverURL := viper.GetString("water_system_url")
	authToken := viper.GetString("auth_token")
	conf := telegram_bot.NewConfig(token, serverURL, authToken)
	t := telegram_bot.NewServer(conf)
	if err := t.Run(); err != nil {
		log.Fatal(err)
	}
}
