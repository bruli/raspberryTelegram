package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	logger *logrus.Logger
}

func (l Logger) Fatal(v ...interface{}) {
	l.logger.Warn(v)
}

func (l Logger) Fatalf(format string, v ...interface{}) {
	l.logger.Warnf(format, v)
}

func NewLogger() *Logger {
	return &Logger{logger: logrus.New()}
}
