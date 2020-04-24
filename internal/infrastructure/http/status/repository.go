package http_status

import (
	"encoding/json"
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/jsontime"
	status2 "github.com/bruli/rasberryTelegram/internal/status"
	"io/ioutil"
	"net/http"
)

type rain struct {
	IsRaining bool   `json:"is_raining"`
	Value     uint16 `json:"value"`
}
type status struct {
	SystemStarted jsontime.JsonTime `json:"system_started"`
	Temperature   float32           `json:"temperature"`
	Humidity      float32           `json:"humidity"`
	OnWater       bool              `json:"on_water"`
	Rain          *rain             `json:"rain"`
}

type Repository struct {
	serverUrl string
}

func NewRepository(serverUrl string) *Repository {
	return &Repository{serverUrl: serverUrl}
}

func (r *Repository) Get() (*status2.Status, error) {
	req, err := http.NewRequest(http.MethodGet, r.serverUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request to get status failed: %w", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	status := status{}
	if err := json.Unmarshal(body, &status); err != nil {
		return nil, fmt.Errorf("failed to unmarshal status body: %w", err)
	}
	return r.buildStatus(status), nil
}

func (r *Repository) buildStatus(s status) *status2.Status {
	return status2.NewStatus(
		s.SystemStarted.ToTime(),
		s.Temperature,
		s.Humidity,
		s.OnWater,
		status2.NewRain(s.Rain.IsRaining, s.Rain.Value))
}
