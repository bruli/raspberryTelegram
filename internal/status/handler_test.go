package status_test

import (
	"errors"
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/log"
	"github.com/bruli/rasberryTelegram/internal/status"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewHandler(t *testing.T) {
	tests := map[string]struct {
		status *status.Status
		err    error
	}{
		"it should return status": {
			status: &status.Status{SystemStarted: time.Now(), Temperature: float32(20), Humidity: float32(40),
				Rain: &status.Rain{IsRaining: false, Value: 1023}, OnWater: false}},
		"it should return error when repository returns error": {
			err: fmt.Errorf("failed to get status: %w", errors.New("error"))},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo := status.RepositoryMock{}
			logger := log.LoggerMock{}
			handler := status.NewHandler(&repo, &logger)

			repo.GetFunc = func() (*status.Status, error) {
				return tt.status, tt.err
			}
			logger.FatalFunc = func(v ...interface{}) {

			}

			st, err := handler.Handle()

			assert.Equal(t, tt.status, st)
			assert.Equal(t, tt.err, err)
		})
	}
}
