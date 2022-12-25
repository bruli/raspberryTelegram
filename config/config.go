package config

import (
	"net/url"

	"github.com/bruli/raspberryRainSensor/pkg/common/env"
)

const (
	ProjectPrefix = "TELEGRAM_"
	Token         = ProjectPrefix + "TOKEN"
	WSServerURl   = ProjectPrefix + "WS_URL"
	WSServerToken = ProjectPrefix + "WS_TOKEN"
)

type Config struct {
	telegramToken,
	wsServerToken string
	wsServerURL url.URL
}

func (c Config) WsServerURL() url.URL {
	return c.wsServerURL
}

func (c Config) WsServerToken() string {
	return c.wsServerToken
}

func (c Config) TelegramToken() string {
	return c.telegramToken
}

func NewConfig() (Config, error) {
	token, err := env.Value(Token)
	if err != nil {
		return Config{}, err
	}
	wsserver, err := env.Value(WSServerURl)
	if err != nil {
		return Config{}, err
	}
	wsServerURL, err := url.Parse(wsserver)
	if err != nil {
		return Config{}, err
	}
	wsToken, err := env.Value(WSServerToken)
	if err != nil {
		return Config{}, err
	}
	return Config{
		telegramToken: token,
		wsServerURL:   *wsServerURL,
		wsServerToken: wsToken,
	}, nil
}
