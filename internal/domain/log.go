package domain

//go:generate moq -out loggerMock.go . Logger
type Logger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}
