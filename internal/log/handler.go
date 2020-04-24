package log

import (
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/logger"
)

type Logs []string

type Handler struct {
	repository Repository
	logger     logger.Logger
}

func NewHandler(repository Repository, logger logger.Logger) *Handler {
	return &Handler{repository: repository, logger: logger}
}

func (h *Handler) Handle(limit uint16) ([]string, error) {
	l, err := h.repository.Get()
	if err != nil {
		h.logger.Fatalf("failed to get logs: %w", err)
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}
	max := uint16(len(l))
	init := max - limit

	var data []string
	for _, j := range l[init:max] {
		data = append(data, j)
	}

	return data, nil
}
