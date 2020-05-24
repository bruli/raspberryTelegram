package status_test

import (
	"errors"
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/logger"
	"github.com/bruli/rasberryTelegram/internal/status"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewGetter(t *testing.T) {
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
			logg := logger.LoggerMock{}
			getter := status.NewGetter(&repo, &logg)

			repo.GetFunc = func() (*status.Status, error) {
				return tt.status, tt.err
			}
			logg.FatalFunc = func(v ...interface{}) {

			}

			st, err := getter.Get()

			assert.Equal(t, tt.status, st)
			assert.Equal(t, tt.err, err)
		})
	}
}
