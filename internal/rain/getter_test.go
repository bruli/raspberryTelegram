package rain

import (
	"errors"
	"github.com/bruli/rasberryTelegram/internal/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetter_Get(t *testing.T) {
	tests := map[string]struct {
		ra               *Rain
		err, expectedErr error
	}{
		"it should return error when repository returns error": {
			err:         errors.New("error"),
			expectedErr: errors.New("error getting rain: error"),
		},
		"it should return rain": {ra: New(true, 200)},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			log := logger.LoggerMock{}
			repo := RepositoryMock{}
			get := &Getter{
				log:  &log,
				repo: &repo,
			}
			log.FatalFunc = func(v ...interface{}) {
			}
			repo.GetFunc = func() (*Rain, error) {
				return tt.ra, tt.err
			}

			ra, err := get.Get()
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.ra, ra)
		})
	}
}
