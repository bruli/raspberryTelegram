package application

import (
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/domain"
)

type TemperatureHandler struct {
	repo   domain.TemperatureRepository
	logger domain.Logger
}

func NewTemperatureHandler(repo domain.TemperatureRepository, logger domain.Logger) *TemperatureHandler {
	return &TemperatureHandler{repo: repo, logger: logger}
}

func (h *TemperatureHandler) Handle() (domain.Temperature, error) {
	t, err := h.repo.Get()
	if err != nil {
		h.logger.Fatalf("error getting temperature: %w", err)
		return domain.Temperature{}, fmt.Errorf("error getting temperature: %w", err)
	}

	return t, nil
}
