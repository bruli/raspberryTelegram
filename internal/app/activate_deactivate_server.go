package app

import (
	"context"

	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
)

const ActivateDeactivateServerCmdName = "activateDeactivateServer"

type ActivateDeactivateServerCmd struct {
	Activate bool
}

func (a ActivateDeactivateServerCmd) Name() string {
	return ActivateDeactivateServerCmdName
}

type ActivateDeactivateServer struct {
	stRepo StatusRepository
}

func (a ActivateDeactivateServer) Handle(ctx context.Context, cmd cqs.Command) ([]cqs.Event, error) {
	co, ok := cmd.(ActivateDeactivateServerCmd)
	if !ok {
		return nil, cqs.NewInvalidCommandError(ActivateDeactivateServerCmdName, cmd.Name())
	}
	return nil, a.stRepo.ActivateDeactivate(ctx, co.Activate)
}

func NewActivateDeactivateServer(stRepo StatusRepository) *ActivateDeactivateServer {
	return &ActivateDeactivateServer{stRepo: stRepo}
}
