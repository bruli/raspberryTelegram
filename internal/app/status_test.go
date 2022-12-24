package app_test

import (
	"context"
	"errors"
	"github.com/bruli/rasberryTelegram/internal/app"
	"github.com/bruli/rasberryTelegram/internal/domain/status"
	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStatusHandle(t *testing.T) {
	errTest := errors.New("")
	st := status.Status{}
	tests := []struct {
		name                 string
		status               status.Status
		findErr, expectedErr error
		expectedResult       cqs.QueryResult
	}{
		{
			name:        "and repository returns an error, then it returns same error",
			findErr:     errTest,
			expectedErr: errTest,
		},
		{
			name:           "and repository returns a result struct, then it returns a valid result",
			status:         st,
			expectedResult: st,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(`Given a Status query handler,
		when Handler is called `+tt.name, func(t *testing.T) {
			t.Parallel()
			sr := &StatusRepositoryMock{
				FindStatusFunc: func(ctx context.Context) (status.Status, error) {
					return tt.status, tt.findErr
				},
			}
			handler := app.NewStatus(sr)
			result, err := handler.Handle(context.Background(), app.StatusQuery{})
			if err != nil {
				require.IsType(t, tt.expectedErr, err)
				return
			}
			require.Equal(t, tt.expectedErr, err)
			require.Equal(t, tt.expectedResult, result)
		})
	}
}
