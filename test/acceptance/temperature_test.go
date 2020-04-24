package acceptance

import (
	"github.com/bruli/rasberryTelegram/internal/infrastructure/http/temperature"
	"github.com/bruli/rasberryTelegram/internal/infrastructure/log/logger"
	"github.com/bruli/rasberryTelegram/internal/temperature"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemperature(t *testing.T) {
	repo := temperature.NewHandler(http_temperature.NewRepository("http://192.168.1.10"), logger.NewLogError())
	temp, err := repo.Handle()

	assert.NoError(t, err)
	assert.NotNil(t, temp)
	assert.NotEqual(t, float32(0), temp.Humidity())
	assert.NotEqual(t, float32(0), temp.Temperature())
}
