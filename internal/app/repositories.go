package app

import (
	"context"

	"github.com/bruli/rasberryTelegram/internal/domain/log"
	"github.com/bruli/rasberryTelegram/internal/domain/status"
	"github.com/bruli/rasberryTelegram/internal/domain/weather"
)

//go:generate go tool moq -out zmock_repositories_test.go -pkg app_test . StatusRepository WeatherRepository LogsRepository ExecutionRepository

type StatusRepository interface {
	FindStatus(ctx context.Context) (status.Status, error)
	ActivateDeactivate(ctx context.Context, activate bool) error
}

type WeatherRepository interface {
	FindWeather(ctx context.Context) (weather.Weather, error)
}

type LogsRepository interface {
	FindLogs(ctx context.Context, number int) ([]log.Log, error)
}

type ExecutionRepository interface {
	ExecuteZone(ctx context.Context, zone string, seconds int) error
}
