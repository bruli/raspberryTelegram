package api

import (
	"context"
	"github.com/bruli/rasberryTelegram/internal/domain/status"
)

type StatusRepository struct {
}

func (s StatusRepository) FindStatus(ctx context.Context) (status.Status, error) {
	//TODO implement me
	panic("implement me")
}
