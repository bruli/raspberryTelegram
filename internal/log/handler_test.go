package log_test

import (
	"errors"
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/log"
	"github.com/bruli/rasberryTelegram/internal/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogHandler(t *testing.T) {
	tests := map[string]struct {
		limit            uint16
		err, formatedErr error
		messages         []string
		logs             log.Logs
	}{
		"it should return error when repository returns error": {
			limit:       2,
			err:         errors.New("error"),
			formatedErr: fmt.Errorf("failed to get logs: %w", errors.New("error")),
		},
		"it should return limited messages": {
			limit:    2,
			messages: []string{"message1", "message2"},
			logs:     log.Logs{"message1", "message2", "message3"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo := log.RepositoryMock{}
			logg := logger.LoggerMock{}
			handler := log.NewHandler(&repo, &logg)

			repo.GetFunc = func() (log.Logs, error) {
				return tt.logs, tt.err
			}
			logg.FatalfFunc = func(format string, v ...interface{}) {
			}

			l, err := handler.Handle(tt.limit)
			assert.Equal(t, tt.formatedErr, err)
			//assert.Equal(t, tt.messages, l)
			if tt.messages != nil {
				assert.Equal(t, tt.limit, uint16(len(l)))
			}
		})
	}
}
