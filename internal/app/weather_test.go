package app_test

import (
	"context"
	"errors"
	"github.com/bruli/rasberryTelegram/internal/app"
	"github.com/bruli/rasberryTelegram/internal/domain/weather"
	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWeatherHandle(t *testing.T) {
	errTest := errors.New("")
	weat := weather.Weather{}
	tests := []struct {
		name                 string
		weather              weather.Weather
		repoErr, expectedErr error
		expectedResult       cqs.QueryResult
	}{
		{
			name:        "and repository returns an error, then it returns same error",
			repoErr:     errTest,
			expectedErr: errTest,
		},
		{
			name:           "and repository returns a weather struct, then it returns a valid result",
			weather:        weat,
			expectedResult: weat,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(`Given a Weather query handler,
		when Handle method is called `+tt.name, func(t *testing.T) {
			t.Parallel()
			wr := &WeatherRepositoryMock{
				FindWeatherFunc: func(ctx context.Context) (weather.Weather, error) {
					return tt.weather, tt.repoErr
				},
			}
			handler := app.NewWeather(wr)
			result, err := handler.Handle(context.Background(), app.WeatherQuery{})
			if err != nil {
				require.IsType(t, tt.expectedErr, err)
				return
			}
			require.Equal(t, tt.expectedErr, err)
			require.Equal(t, tt.expectedResult, result)
		})
	}
}
