package temperature

import (
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/log"
)

type Handler struct {
	repo   Repository
	logger log.Logger
}

func NewHandler(repo Repository, logger log.Logger) *Handler {
	return &Handler{repo: repo, logger: logger}
}

func (h *Handler) Handle() (Temperature, error) {
	t, err := h.repo.Get()
	if err != nil {
		h.logger.Fatalf("error getting temperature: %w", err)
		return Temperature{}, fmt.Errorf("error getting temperature: %w", err)
	}

	return t, nil
}
