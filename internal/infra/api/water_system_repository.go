package api

import (
	"context"
	"fmt"
	"github.com/bruli/rasberryTelegram/internal/domain/status"
	"github.com/bruli/raspberryWaterSystem/pkg/ws"
	"net/url"
)

type WaterSystemRepository struct {
	api
}

func (s WaterSystemRepository) FindStatus(ctx context.Context) (status.Status, error) {
	st, err := s.pkg.GetStatus(ctx)
	if err != nil {
		return status.Status{}, fmt.Errorf("failed to find status: %w", err)
	}
	return buildStatusDomain(st), nil
}

func buildStatusDomain(s ws.Status) status.Status {
	var st status.Status
	st.Hydrate(s.Humidity, s.IsRaining, s.SystemStartedAt, s.Temperature, s.UpdatedAt)
	return st
}

func NewWaterSystemRepository(serverURL url.URL, client ws.HTTPClient, token string) WaterSystemRepository {
	return WaterSystemRepository{api: newApi(serverURL, client, token)}
}
