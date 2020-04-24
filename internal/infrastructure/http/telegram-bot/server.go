package telegram_bot

import (
	"fmt"
	http_log "github.com/bruli/rasberryTelegram/internal/infrastructure/http/log"
	http_status "github.com/bruli/rasberryTelegram/internal/infrastructure/http/status"
	http_temperature "github.com/bruli/rasberryTelegram/internal/infrastructure/http/temperature"
	http_water "github.com/bruli/rasberryTelegram/internal/infrastructure/http/water"
	"github.com/bruli/rasberryTelegram/internal/infrastructure/log/logger"
	"github.com/bruli/rasberryTelegram/internal/log"
	"github.com/bruli/rasberryTelegram/internal/status"
	temperature2 "github.com/bruli/rasberryTelegram/internal/temperature"
	"github.com/bruli/rasberryTelegram/internal/water"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
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
	logs   *log.Handler
	water  *water.Handler
}

func NewServer(config *Config) *Server {
	logger := logger.NewLogger()
	return &Server{config: config,
		mess:   messages{},
		temp:   temperature2.NewHandler(http_temperature.NewRepository(config.waterSystemUrl), logger),
		status: status.NewHandler(http_status.NewRepository(config.waterSystemUrl), logger),
		logs:   log.NewHandler(http_log.NewRepository(config.waterSystemUrl), logger),
		water:  water.NewHandler(http_water.NewRepository(config.waterSystemUrl), logger)}
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
				msg.Text = "Type /temp /status /log [limit] /water [zone] [seconds]."
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
			case "log":
				ar := update.Message.CommandArguments()
				args := strings.Fields(ar)
				var limit uint16
				if 0 == len(args) {
					limit = 2
				} else {
					a, err := strconv.ParseUint(args[0], 10, 8)
					if err != nil {
						msg.Text = fmt.Sprintf("invalid limit: %s", err.Error())
						s.mess.addMessage(msg)
					} else {
						limit = uint16(a)
						logs, err := s.logs.Handle(limit)
						if err != nil {
							msg.Text = fmt.Sprintf("Error: %s", err.Error())
							s.mess.addMessage(msg)
						}
						for _, j := range logs {
							msg.Text = j
							s.mess.addMessage(msg)
						}
					}
				}
			case "water":
				arg := update.Message.CommandArguments()
				args := strings.Fields(arg)
				if 2 != len(args) {
					msg.Text = "Invalid arguments. Required [zone][seconds]"
					s.mess.addMessage(msg)
				} else {
					seconds, _ := strconv.ParseUint(args[1], 10, 8)
					zone := args[0]

					if err := s.water.Handle(zone, uint8(seconds)); err != nil {
						msg.Text = fmt.Sprint("failed to execute water on zone %s: %w", zone, err)
						s.mess.addMessage(msg)
					}
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
