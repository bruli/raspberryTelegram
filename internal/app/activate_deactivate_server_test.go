package app_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bruli/rasberryTelegram/internal/app"
	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
	"github.com/stretchr/testify/require"
)

func TestActivateDeactivateServer_Handle(t *testing.T) {
	errTest := errors.New("")
	cmd := app.ActivateDeactivateServerCmd{Activate: true}
	tests := []struct {
		name                 string
		cmd                  cqs.Command
		expectedErr, repoErr error
	}{
		{
			name:        "with an invalid command, then it returns an invalid command error",
			cmd:         &invalidCommand{},
			expectedErr: cqs.InvalidCommandError{},
		},
		{
			name:        "with a valid command and repository returns an error, then it returns same error",
			cmd:         cmd,
			repoErr:     errTest,
			expectedErr: errTest,
		},
		{
			name: "with a valid command and repository returns nil, then it returns nil",
			cmd:  cmd,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(`Given a ActivateDeactivateServer command handler,
		when Handle method is called `+tt.name, func(t *testing.T) {
			t.Parallel()
			repo := &StatusRepositoryMock{}
			repo.ActivateDeactivateFunc = func(ctx context.Context, activate bool) error {
				return tt.repoErr
			}
			handler := app.NewActivateDeactivateServer(repo)
			_, err := handler.Handle(context.Background(), tt.cmd)
			if err != nil {
				require.ErrorAs(t, err, &tt.expectedErr)
				return
			}
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

type invalidCommand struct{}

func (i invalidCommand) Name() string {
	return "invalid"
}
