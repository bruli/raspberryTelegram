package app

import (
	"context"
	"github.com/bruli/rasberryTelegram/internal/domain/status"
)

//go:generate moq -out zmock_repositories_test.go -pkg app_test . StatusRepository

type StatusRepository interface {
	FindStatus(ctx context.Context) (status.Status, error)
}
