package temperature

import (
	"fmt"
	temperature2 "github.com/bruli/rasberryTelegram/internal/temperature"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"time"
)

type temperature struct {
	Humidity    float32 `json:"humidity"`
	Temperature float32 `json:"temperature"`
}

type Repository struct {
	serverUrl string
}

func NewRepository(serverUrl string) *Repository {
	return &Repository{serverUrl: serverUrl}
}

func (r *Repository) Get() (temperature2.Temperature, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/temperature", r.serverUrl), nil)
	if err != nil {
		return temperature2.Temperature{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	cl := http.DefaultClient
	cl.Timeout = 5 * time.Second
	res, err := cl.Do(req)
	if err != nil {
		return temperature2.Temperature{}, fmt.Errorf("error getting temperature from water system: %w", err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	data := temperature{}
	err = jsoniter.Unmarshal(body, &data)
	if err != nil {
		return temperature2.Temperature{}, fmt.Errorf("error reading temperature response body: %w", err)
	}

	temp := temperature2.NewTemperature(data.Humidity, data.Temperature)

	return *temp, err
}
