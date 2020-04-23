package log

import (
	"github.com/bruli/rasberryTelegram/internal/domain"
	"log"
	"os"
)

func NewLogError() domain.Logger {
	return log.New(os.Stdout, "ERROR", 1)
}
