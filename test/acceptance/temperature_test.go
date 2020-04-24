package acceptance

import (
	"github.com/bruli/rasberryTelegram/internal/infrastructure/http/temperature"
	"github.com/bruli/rasberryTelegram/internal/infrastructure/log/logger"
	"github.com/bruli/rasberryTelegram/internal/temperature"
	"github.com/stretchr/testify/assert"
	"testing"
)

const serverUrl = "http://192.168.1.10"

func TestTemperature(t *testing.T) {
	handler := temperature.NewHandler(http_temperature.NewRepository(serverUrl), logger.NewLogger())
	temp, err := handler.Handle()

	assert.NoError(t, err)
	assert.NotNil(t, temp)
	assert.NotEqual(t, float32(0), temp.Humidity())
	assert.NotEqual(t, float32(0), temp.Temperature())
}
