package log

import (
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/logger"
)

type Logs []string

type Getter struct {
	repository Repository
	logger     logger.Logger
}

func NewGetter(repository Repository, logger logger.Logger) *Getter {
	return &Getter{repository: repository, logger: logger}
}

func (g *Getter) Get(limit uint16) ([]string, error) {
	l, err := g.repository.Get()
	if err != nil {
		g.logger.Fatalf("failed to get logs: %w", err)
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}

	var data []string
	for _, j := range l[:limit] {
		data = append(data, j)
	}

	return data, nil
}
