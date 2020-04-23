package logger

import (
	log2 "github.com/bruli/rasberryTelegram/internal/log"
	"log"
	"os"
)

func NewLogError() log2.Logger {
	return log.New(os.Stdout, "ERROR", 1)
}
