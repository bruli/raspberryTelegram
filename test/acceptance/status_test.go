package acceptance

import (
	http_status "github.com/bruli/rasberryTelegram/internal/infrastructure/http/status"
	"github.com/bruli/rasberryTelegram/internal/infrastructure/log/logger"
	"github.com/bruli/rasberryTelegram/internal/status"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatus(t *testing.T) {
	handler := status.NewHandler(http_status.NewRepository(serverUrl), logger.NewLogger())
	st, err := handler.Handle()

	assert.Nil(t, err)
	assert.NotNil(t, st)
	assert.NotNil(t, st.SystemStarted)
	assert.NotNil(t, st.OnWater)
	assert.NotNil(t, st.Humidity)
	assert.NotNil(t, st.Temperature)
	assert.NotNil(t, st.Rain)

}
