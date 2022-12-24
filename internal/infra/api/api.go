package api

import (
	"github.com/bruli/raspberryWaterSystem/pkg/ws"
	"net/url"
)

type api struct {
	pkg ws.Handlers
}

func newApi(serverURL url.URL, client ws.HTTPClient, token string) api {
	return api{pkg: ws.New(serverURL, client, token)}
}
