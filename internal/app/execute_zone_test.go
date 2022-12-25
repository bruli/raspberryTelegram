package app_test

import (
	"context"
	"errors"
	"github.com/bruli/rasberryTelegram/internal/app"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExecuteZoneHandle(t *testing.T) {
	errTest := errors.New("")
	tests := []struct {
		name                 string
		repoErr, expectedErr error
	}{
		{
			name:        "and repository returns an error, then it returns same error",
			expectedErr: errTest,
			repoErr:     errTest,
		},
		{
			name: "and repository returns nil, then it returns nil",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(`Given an ExecuteZone command handler,
		when Handle method is called `+tt.name, func(t *testing.T) {
			t.Parallel()
			er := &ExecutionRepositoryMock{
				ExecuteZoneFunc: func(ctx context.Context, zone string, seconds int) error {
					return tt.repoErr
				},
			}
			handler := app.NewExecuteZone(er)
			_, err := handler.Handle(context.Background(), app.ExecuteZoneCommand{})
			if err != nil {
				require.IsType(t, tt.expectedErr, err)
				return
			}
			require.Equal(t, tt.expectedErr, err)
		})
	}
}
