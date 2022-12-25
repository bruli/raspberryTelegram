package app

import (
	"context"

	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
)

const StatusQueryName = "status"

type StatusQuery struct{}

func (s StatusQuery) Name() string {
	return StatusQueryName
}

type Status struct {
	sr StatusRepository
}

func (s Status) Handle(ctx context.Context, query cqs.Query) (cqs.QueryResult, error) {
	return s.sr.FindStatus(ctx)
}

func NewStatus(sr StatusRepository) Status {
	return Status{sr: sr}
}
