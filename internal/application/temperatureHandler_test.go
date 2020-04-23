package application_test

import (
	"errors"
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/application"
	"github.com/bruli/rasberryTelegram/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTemperatureHandler(t *testing.T) {
	tests := map[string]struct {
		temp                domain.Temperature
		error, formattedErr error
	}{
		"it should return error when repository returns error": {temp: domain.Temperature{}, error: errors.New("error"),
			formattedErr: fmt.Errorf("error getting temperature: %w", errors.New("error"))},
		"it should return temperature": {temp: *domain.NewTemperature(40, 20)},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo := domain.TemperatureRepositoryMock{}
			logger := domain.LoggerMock{}
			handler := application.NewTemperatureHandler(&repo, &logger)

			repo.GetFunc = func() (domain.Temperature, error) {
				return tt.temp, tt.error
			}

			logger.FatalfFunc = func(format string, v ...interface{}) {
			}
			temp, err := handler.Handle()
			assert.Equal(t, tt.temp, temp)
			assert.Equal(t, tt.formattedErr, err)
		})
	}
}
