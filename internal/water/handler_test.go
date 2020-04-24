package water_test

import (
	"errors"
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/logger"
	"github.com/bruli/rasberryTelegram/internal/water"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHandler(t *testing.T) {
	tests := map[string]struct {
		zone              string
		seconds           uint8
		err, formattedErr error
	}{
		"it should return error when repository returns error": {
			zone:         "a",
			seconds:      uint8(20),
			err:          errors.New("error"),
			formattedErr: fmt.Errorf("failed to execute water in zone a: %w", errors.New("error")),
		},
		"it should execute": {
			zone:    "a",
			seconds: uint8(20),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo := water.RepositoryMock{}
			logg := logger.LoggerMock{}
			handler := water.NewHandler(&repo, &logg)

			repo.ExecuteFunc = func(zone string, seconds uint8) error {
				return tt.err
			}
			logg.FatalfFunc = func(format string, v ...interface{}) {
			}

			err := handler.Handle(tt.zone, tt.seconds)

			assert.Equal(t, tt.formattedErr, err)
		})
	}
}
