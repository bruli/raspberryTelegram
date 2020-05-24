package water

import (
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/logger"
)

type Getter struct {
	repository Repository
	logger     logger.Logger
}

func NewGetter(repository Repository, logger logger.Logger) *Getter {
	return &Getter{repository: repository, logger: logger}
}

func (g *Getter) Get(zone string, seconds uint8) error {
	if err := g.repository.Execute(zone, seconds); err != nil {
		g.logger.Fatalf("failed to execute water in zone %s: %w", zone, err)
		return fmt.Errorf("failed to execute water in zone %s: %w", zone, err)
	}
	return nil
}
