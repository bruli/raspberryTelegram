package rain

import (
	"errors"
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/logger"
)

type Getter struct {
	log  logger.Logger
	repo Repository
}

func NewGetter(log logger.Logger, repo Repository) *Getter {
	return &Getter{log: log, repo: repo}
}

func (g *Getter) Get() (*Rain, error) {
	r, err := g.repo.Get()
	if err != nil {
		mesg := fmt.Sprintf("error getting rain: %s", err)
		g.log.Fatal(mesg)
		return nil, errors.New(mesg)
	}
	return r, nil
}
