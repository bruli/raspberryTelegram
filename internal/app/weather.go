package app

import (
	"context"
	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
)

const WeatherQueryName = "weather"

type WeatherQuery struct{}

func (w WeatherQuery) Name() string {
	return WeatherQueryName
}

type Weather struct {
	wr WeatherRepository
}

func (w Weather) Handle(ctx context.Context, _ cqs.Query) (cqs.QueryResult, error) {
	return w.wr.FindWeather(ctx)
}

func NewWeather(wr WeatherRepository) Weather {
	return Weather{wr: wr}
}
