package app

import (
	"context"

	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
)

const ExecuteZoneCommandName = "executeZone"

type ExecuteZoneCommand struct {
	Zone    string
	Seconds int
}

func (e ExecuteZoneCommand) Name() string {
	return ExecuteZoneCommandName
}

type ExecuteZone struct {
	er ExecutionRepository
}

func (e ExecuteZone) Handle(ctx context.Context, cmd cqs.Command) ([]cqs.Event, error) {
	co, _ := cmd.(ExecuteZoneCommand)
	return nil, e.er.ExecuteZone(ctx, co.Zone, co.Seconds)
}

func NewExecuteZone(er ExecutionRepository) ExecuteZone {
	return ExecuteZone{er: er}
}
