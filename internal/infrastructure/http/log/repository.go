package http_log

import (
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/jsontime"
	"github.com/bruli/rasberryTelegram/internal/log"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

type logData struct {
	Message   string            `json:"message"`
	CreatedAt jsontime.JsonTime `json:"created_at"`
}

type logs []*logData

type Repository struct {
	serverUrl string
}

func NewRepository(serverUrl string) *Repository {
	return &Repository{serverUrl: serverUrl}
}

func (r *Repository) Get() (log.Logs, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/log", r.serverUrl), nil)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending logs request: %w", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("errof reading log response body: %w", err)
	}
	logs := logs{}
	if err := jsoniter.Unmarshal(body, &logs); err != nil {
		return nil, err
	}

	var data log.Logs
	for _, l := range logs {
		data = append(data, fmt.Sprintf("%s, executed at %s", l.Message, l.CreatedAt.ToString()))
	}

	return data, nil
}
