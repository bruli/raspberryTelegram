package temperature

import (
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/logger"
)

type Getter struct {
	repo   Repository
	logger logger.Logger
}

func NewGetter(repo Repository, logger logger.Logger) *Getter {
	return &Getter{repo: repo, logger: logger}
}

func (g *Getter) Get() (*Temperature, error) {
	t, err := g.repo.Get()
	if err != nil {
		g.logger.Fatalf("error getting temperature: %w", err)
		return nil, fmt.Errorf("error getting temperature: %w", err)
	}

	return &t, nil
}
