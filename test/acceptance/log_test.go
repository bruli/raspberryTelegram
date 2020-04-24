package acceptance

import (
	http_log "github.com/bruli/rasberryTelegram/internal/infrastructure/http/log"
	"github.com/bruli/rasberryTelegram/internal/infrastructure/log/logger"
	"github.com/bruli/rasberryTelegram/internal/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLog(t *testing.T) {
	limit := uint16(2)
	handler := log.NewHandler(http_log.NewRepository(serverUrl), logger.NewLogger())
	logs, err := handler.Handle(limit)

	assert.Nil(t, err)
	assert.Equal(t, limit, uint16(len(logs)))
}
