package api

import (
	"net/url"

	"github.com/bruli/raspberryWaterSystem/pkg/ws"
)

type api struct {
	pkg ws.Handlers
}

func newApi(serverURL url.URL, client ws.HTTPClient, token string) api {
	return api{pkg: ws.New(serverURL, client, token)}
}
