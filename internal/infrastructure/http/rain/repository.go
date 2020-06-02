package http_rain

import (
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/rain"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"time"
)

type rainBody struct {
	IsRain bool   `json:"is_rain"`
	Value  uint16 `json:"value"`
}

type Repository struct {
	serverUrl, authToken string
}

func (r *Repository) Get() (*rain.Rain, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/rain", r.serverUrl), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", r.authToken)

	cl := http.DefaultClient
	cl.Timeout = 5 * time.Second
	res, err := cl.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting rain from water system: %w", err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	data := rainBody{}
	err = jsoniter.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error getting rain from water system: %w", err)
	}

	ra := rain.New(data.IsRain, data.Value)
	return ra, nil
}

func NewRepository(serverUrl, authToken string) *Repository {
	return &Repository{serverUrl: serverUrl, authToken: authToken}
}
