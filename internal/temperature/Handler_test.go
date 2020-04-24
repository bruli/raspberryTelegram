package temperature_test

import (
	"errors"
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/log"
	"github.com/bruli/rasberryTelegram/internal/temperature"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTemperatureHandler(t *testing.T) {
	tests := map[string]struct {
		temp                *temperature.Temperature
		error, formattedErr error
	}{
		"it should return error when repository returns error": {error: errors.New("error"),
			formattedErr: fmt.Errorf("error getting temperature: %w", errors.New("error"))},
		"it should return temperature": {temp: temperature.NewTemperature(40, 20)},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo := temperature.RepositoryMock{}
			logger := log.LoggerMock{}
			handler := temperature.NewHandler(&repo, &logger)

			repo.GetFunc = func() (temperature.Temperature, error) {
				if tt.temp == nil {
					return temperature.Temperature{}, tt.error
				}
				return *tt.temp, tt.error
			}

			logger.FatalfFunc = func(format string, v ...interface{}) {
			}
			temp, err := handler.Handle()
			assert.Equal(t, tt.temp, temp)
			assert.Equal(t, tt.formattedErr, err)
		})
	}
}
