package temperature_test

import (
	"errors"
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/logger"
	"github.com/bruli/rasberryTelegram/internal/temperature"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTemperatureGetter(t *testing.T) {
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
			logg := logger.LoggerMock{}
			getter := temperature.NewGetter(&repo, &logg)

			repo.GetFunc = func() (temperature.Temperature, error) {
				if tt.temp == nil {
					return temperature.Temperature{}, tt.error
				}
				return *tt.temp, tt.error
			}

			logg.FatalfFunc = func(format string, v ...interface{}) {
			}
			temp, err := getter.Get()
			assert.Equal(t, tt.temp, temp)
			assert.Equal(t, tt.formattedErr, err)
		})
	}
}
