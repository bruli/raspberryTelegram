package acceptance

import (
	http_status "github.com/bruli/rasberryTelegram/internal/infrastructure/http/status"
	"github.com/bruli/rasberryTelegram/internal/infrastructure/log/logger"
	"github.com/bruli/rasberryTelegram/internal/status"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatus(t *testing.T) {
	authToken, err := getAuthToken()
	assert.NoError(t, err)
	getter := status.NewGetter(http_status.NewRepository(serverUrl, authToken), logger.NewLogger())
	st, err := getter.Get()

	assert.Nil(t, err)
	assert.NotNil(t, st)
	assert.NotNil(t, st.SystemStarted)
	assert.NotNil(t, st.OnWater)
	assert.NotNil(t, st.Humidity)
	assert.NotNil(t, st.Temperature)
	assert.NotNil(t, st.Rain)

}
