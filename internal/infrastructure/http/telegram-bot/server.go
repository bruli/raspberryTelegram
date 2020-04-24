package telegram_bot

import (
	"fmt"
	http_status "github.com/bruli/rasberryTelegram/internal/infrastructure/http/status"
	http_temperature "github.com/bruli/rasberryTelegram/internal/infrastructure/http/temperature"
	"github.com/bruli/rasberryTelegram/internal/infrastructure/log/logger"
	"github.com/bruli/rasberryTelegram/internal/status"
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
	status *status.Handler
}

func NewServer(config *Config) *Server {
	logger := logger.NewLogger()
	return &Server{config: config,
		mess:   messages{},
		temp:   temperature2.NewHandler(http_temperature.NewRepository(config.waterSystemUrl), logger),
		status: status.NewHandler(http_status.NewRepository(config.waterSystemUrl), logger)}
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
				msg.Text = "Type /temp."
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
			case "status":
				st, err := s.status.Handle()
				if err != nil {
					msg.Text = err.Error()
					s.mess.addMessage(msg)
				} else {
					msg.Text = fmt.Sprintf("System started at %s", st.SystemStarted.Format("2006-01-02 15:04:05"))
					s.mess.addMessage(msg)
					msg.Text = fmt.Sprintf("Current temperature %v *C", st.Temperature)
					s.mess.addMessage(msg)
					msg.Text = fmt.Sprintf("Current humidity %v %%", st.Humidity)
					s.mess.addMessage(msg)
					var bBusy string
					if st.OnWater {
						bBusy = "Executing water system"
					} else {
						bBusy = "Water system not running"
					}
					msg.Text = bBusy
					s.mess.addMessage(msg)
					msg.Text = fmt.Sprintf("Is raining %v, raining value %v", st.Rain.IsRaining, st.Rain.Value)
					s.mess.addMessage(msg)
				}
			}
		}

		for _, j := range s.mess {
			_, _ = bot.Send(j)
		}

		s.mess = nil
	}
	return nil
}
