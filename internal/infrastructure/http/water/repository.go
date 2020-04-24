package http_water

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type body struct {
	Seconds uint8    `json:"seconds"`
	Zones   []string `json:"zones"`
}

func newBody(seconds uint8, zone string) body {
	var zones []string
	zones = append(zones, zone)
	return body{Seconds: seconds, Zones: zones}
}

type Repository struct {
	serverUrl string
}

func (r *Repository) Execute(zone string, seconds uint8) error {
	body := newBody(seconds, zone)
	b, err := jsoniter.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/water", r.serverUrl), bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed sending request: %w", err)
	}
	if 202 != res.StatusCode {
		return fmt.Errorf("Error: %s", res.Status)
	}

	return nil
}

func NewRepository(serverUrl string) *Repository {
	return &Repository{serverUrl: serverUrl}
}
