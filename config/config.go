package config

import (
	"github.com/bruli/raspberryRainSensor/pkg/common/env"
)

const (
	ProjectPrefix = "TELEGRAM_"
	ServerURL     = ProjectPrefix + "SERVER_URL"
	Token         = ProjectPrefix + "TOKEN"
	AuthToken     = ProjectPrefix + "AUTH_TOKEN"
)

type Config struct {
	serverUrl,
	telegramToken,
	authToken string
}

func (c Config) ServerUrl() string {
	return c.serverUrl
}

func (c Config) TelegramToken() string {
	return c.telegramToken
}

func (c Config) AuthToken() string {
	return c.authToken
}

func NewConfig() (Config, error) {
	url, err := env.Value(ServerURL)
	if err != nil {
		return Config{}, err
	}
	token, err := env.Value(Token)
	if err != nil {
		return Config{}, err
	}
	auth, err := env.Value(AuthToken)
	if err != nil {
		return Config{}, err
	}
	return Config{
		serverUrl:     url,
		telegramToken: token,
		authToken:     auth,
	}, nil
}
