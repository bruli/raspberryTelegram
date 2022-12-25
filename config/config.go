package config

import (
	"github.com/bruli/raspberryRainSensor/pkg/common/env"
	"net/url"
)

const (
	ProjectPrefix = "TELEGRAM_"
	ServerURL     = ProjectPrefix + "SERVER_URL"
	Token         = ProjectPrefix + "TOKEN"
	AuthToken     = ProjectPrefix + "AUTH_TOKEN"
	WSServerURl   = ProjectPrefix + "WS_URL"
	WSServerToken = ProjectPrefix + "WS_TOKEN"
)

type Config struct {
	serverUrl,
	telegramToken,
	authToken,
	wsServerToken string
	wsServerURL url.URL
}

func (c Config) WsServerURL() url.URL {
	return c.wsServerURL
}

func (c Config) WsServerToken() string {
	return c.wsServerToken
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
	serverURL, err := env.Value(ServerURL)
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
		serverUrl:     serverURL,
		telegramToken: token,
		authToken:     auth,
		wsServerURL:   *wsServerURL,
		wsServerToken: wsToken,
	}, nil
}
