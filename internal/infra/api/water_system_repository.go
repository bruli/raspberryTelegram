package api

import (
	"context"
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/domain/log"
	"github.com/bruli/rasberryTelegram/internal/domain/status"
	"github.com/bruli/rasberryTelegram/internal/domain/weather"
	"github.com/bruli/raspberryWaterSystem/pkg/ws"
	"net/url"
)

type WaterSystemRepository struct {
	api
}

func (s WaterSystemRepository) FindLogs(ctx context.Context, number int) ([]log.Log, error) {
	logs, err := s.pkg.GetLogs(ctx, number)
	if err != nil {
		return nil, fmt.Errorf("failed to find logs, %w", err)
	}
	return buildLogs(logs), nil
}

func buildLogs(lo []ws.Log) []log.Log {
	logs := make([]log.Log, len(lo))
	for i, n := range lo {
		var l log.Log
		l.Hydrate(n.ExecutedAt, n.Seconds, n.ZoneName)
		logs[i] = l
	}
	return logs
}

func (s WaterSystemRepository) FindWeather(ctx context.Context) (weather.Weather, error) {
	weat, err := s.pkg.GetWeather(ctx)
	if err != nil {
		return weather.Weather{}, fmt.Errorf("failed to find weather, %w", err)
	}
	return buildWeather(weat), nil
}

func buildWeather(weat ws.Weather) weather.Weather {
	var w weather.Weather
	w.Hydrate(int(weat.Temperature), int(weat.Humidity), weat.IsRaining)
	return w
}

func (s WaterSystemRepository) FindStatus(ctx context.Context) (status.Status, error) {
	st, err := s.pkg.GetStatus(ctx)
	if err != nil {
		return status.Status{}, fmt.Errorf("failed to find status: %w", err)
	}
	return buildStatus(st), nil
}

func buildStatus(s ws.Status) status.Status {
	var st status.Status
	st.Hydrate(int(s.Humidity), s.IsRaining, s.SystemStartedAt, int(s.Temperature), s.UpdatedAt)
	return st
}

func NewWaterSystemRepository(serverURL url.URL, client ws.HTTPClient, token string) WaterSystemRepository {
	return WaterSystemRepository{api: newApi(serverURL, client, token)}
}
