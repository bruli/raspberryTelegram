package water

import (
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/logger"
)

type Handler struct {
	repository Repository
	logger     logger.Logger
}

func NewHandler(repository Repository, logger logger.Logger) *Handler {
	return &Handler{repository: repository, logger: logger}
}

func (h *Handler) Handle(zone string, seconds uint8) error {
	if err := h.repository.Execute(zone, seconds); err != nil {
		h.logger.Fatalf("failed to execute water in zone %s: %w", zone, err)
		return fmt.Errorf("failed to execute water in zone %s: %w", zone, err)
	}
	return nil
}
