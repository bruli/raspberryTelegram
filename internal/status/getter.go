package status

import (
	"github.com/bruli/rasberryTelegram/internal/logger"
	"time"
)

type Rain struct {
	IsRaining bool
	Value     uint16
}

func NewRain(isRaining bool, value uint16) *Rain {
	return &Rain{IsRaining: isRaining, Value: value}
}

type Status struct {
	SystemStarted         time.Time
	Temperature, Humidity float32
	OnWater               bool
	Rain                  *Rain
}

func NewStatus(systemStarted time.Time, temperature float32, humidity float32, onWater bool, rain *Rain) *Status {
	return &Status{
		SystemStarted: systemStarted,
		Temperature:   temperature,
		Humidity:      humidity,
		OnWater:       onWater,
		Rain:          rain}
}

type Getter struct {
	repository Repository
	logger     logger.Logger
}

func NewGetter(repository Repository, logger logger.Logger) *Getter {
	return &Getter{repository: repository, logger: logger}
}

func (g *Getter) Get() (*Status, error) {
	st, err := g.repository.Get()
	if err != nil {
		g.logger.Fatal(err)
		return nil, err
	}

	return st, nil
}
