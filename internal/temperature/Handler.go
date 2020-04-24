package temperature

import (
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/logger"
)

type Handler struct {
	repo   Repository
	logger logger.Logger
}

func NewHandler(repo Repository, logger logger.Logger) *Handler {
	return &Handler{repo: repo, logger: logger}
}

func (h *Handler) Handle() (*Temperature, error) {
	t, err := h.repo.Get()
	if err != nil {
		h.logger.Fatalf("error getting temperature: %w", err)
		return nil, fmt.Errorf("error getting temperature: %w", err)
	}

	return &t, nil
}
