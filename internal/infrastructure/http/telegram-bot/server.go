package telegram_bot

import (
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/infrastructure/http/temperature"
	"github.com/bruli/rasberryTelegram/internal/infrastructure/log/logger"
	temperature2 "github.com/bruli/rasberryTelegram/internal/temperature"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Config struct {
	token, waterSystemUrl string
}

type messages []tgbotapi.MessageConfig

func (m *messages) addMessage(msg tgbotapi.MessageConfig) {
	*m = append(*m, msg)
}

func NewConfig(token, waterSystemUrl string) *Config {
	return &Config{token: token, waterSystemUrl: waterSystemUrl}
}

type Server struct {
	config *Config
	mess   messages
	temp   *temperature2.Handler
}

func NewServer(config *Config) *Server {
	logger := logger.NewLogError()
	return &Server{config: config,
		mess: messages{},
		temp: temperature2.NewHandler(temperature.NewRepository(config.waterSystemUrl), logger)}
}

func (s *Server) Run() error {
	bot, err := tgbotapi.NewBotAPI(s.config.token)
	if err != nil {
		return fmt.Errorf("error running bot: %w", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		return fmt.Errorf("error updating channel: %w", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = "Type /temp, /status or /water [zone] [seconds]."
				s.mess.addMessage(msg)
			case "temp":
				tmp, err := s.temp.Handle()
				if err != nil {
					msg.Text = err.Error()
					s.mess.addMessage(msg)
				} else {
					msg.Text = fmt.Sprintf("Current temperature %v *C", tmp.Temperature())
					s.mess.addMessage(msg)
					msg.Text = fmt.Sprintf("Current humidity %v %%", tmp.Humidity())
					s.mess.addMessage(msg)
				}
			}
		}

		for _, j := range s.mess {
			_, _ = bot.Send(j)
		}
	}
	return nil
}
