package app_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bruli/rasberryTelegram/internal/app"
	"github.com/bruli/rasberryTelegram/internal/domain/log"
	"github.com/stretchr/testify/require"
)

func TestLogsHandle(t *testing.T) {
	errTest := errors.New("")
	logs := []log.Log{
		{},
		{},
	}
	tests := []struct {
		name                 string
		repoErr, expectedErr error
		logs                 []log.Log
	}{
		{
			name:        "and repository returns an error, then it returns same error",
			repoErr:     errTest,
			expectedErr: errTest,
		},
		{
			name: "and repository returns logs, then it returns a valid result",
			logs: logs,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(`Given a Logs query handler,
		when Handle method is called `+tt.name, func(t *testing.T) {
			t.Parallel()
			lr := &LogsRepositoryMock{
				FindLogsFunc: func(ctx context.Context, number int) ([]log.Log, error) {
					return tt.logs, tt.repoErr
				},
			}
			handler := app.NewLogs(lr)
			result, err := handler.Handle(context.Background(), app.LogsQuery{})
			if err != nil {
				require.IsType(t, tt.expectedErr, err)
				return
			}
			_, ok := result.([]string)
			require.Equal(t, tt.expectedErr, err)
			require.True(t, ok)
		})
	}
}
