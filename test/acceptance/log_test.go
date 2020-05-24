package acceptance

import (
	http_log "github.com/bruli/rasberryTelegram/internal/infrastructure/http/log"
	"github.com/bruli/rasberryTelegram/internal/infrastructure/log/logger"
	"github.com/bruli/rasberryTelegram/internal/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLog(t *testing.T) {
	authToken, err := getAuthToken()
	assert.NoError(t, err)
	limit := uint16(2)
	handler := log.NewGetter(http_log.NewRepository(serverUrl, authToken), logger.NewLogger())
	logs, err := handler.Get(limit)

	assert.Nil(t, err)
	assert.Equal(t, limit, uint16(len(logs)))
}
