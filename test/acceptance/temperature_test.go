package acceptance

import (
	"github.com/bruli/rasberryTelegram/internal/infrastructure/http/temperature"
	"github.com/bruli/rasberryTelegram/internal/infrastructure/log/logger"
	"github.com/bruli/rasberryTelegram/internal/temperature"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const serverUrl = "http://192.168.1.10"

func TestTemperature(t *testing.T) {
	authToken, err := getAuthToken()
	assert.NoError(t, err)
	handler := temperature.NewGetter(http_temperature.NewRepository(serverUrl, authToken), logger.NewLogger())
	temp, err := handler.Get()

	assert.NoError(t, err)
	assert.NotNil(t, temp)
	assert.NotEqual(t, float32(0), temp.Humidity())
	assert.NotEqual(t, float32(0), temp.Temperature())
}

func getAuthToken() (string, error) {
	data, err := ioutil.ReadFile("./assets/authToken.txt")
	if err != nil {
		return "", err
	}
	return string(data), err
}
