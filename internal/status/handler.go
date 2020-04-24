package status

import (
	"github.com/bruli/rasberryTelegram/internal/log"
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

type Handler struct {
	repository Repository
	logger     log.Logger
}

func NewHandler(repository Repository, logger log.Logger) *Handler {
	return &Handler{repository: repository, logger: logger}
}

func (h *Handler) Handle() (*Status, error) {
	st, err := h.repository.Get()
	if err != nil {
		h.logger.Fatal(err)
		return nil, err
	}

	return st, nil
}
