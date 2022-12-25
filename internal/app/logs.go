package app

import (
	"context"
	"fmt"

	"github.com/bruli/rasberryTelegram/internal/domain/log"
	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
)

const LogsQueryName = "logs"

type LogsQuery struct {
	Number int
}

func (l LogsQuery) Name() string {
	return LogsQueryName
}

type Logs struct {
	lr LogsRepository
}

func (l Logs) Handle(ctx context.Context, query cqs.Query) (cqs.QueryResult, error) {
	q, _ := query.(LogsQuery)
	logs, err := l.lr.FindLogs(ctx, q.Number)
	if err != nil {
		return nil, err
	}
	return buildLogs(logs), nil
}

func buildLogs(logs []log.Log) []string {
	result := make([]string, len(logs))
	for i, n := range logs {
		result[i] = fmt.Sprintf("Zone: %s, seconds: %v, executed at: %s", n.ZoneName(), n.Seconds(), n.ExecutedAt().Date())
	}
	return result
}

func NewLogs(lr LogsRepository) Logs {
	return Logs{lr: lr}
}
