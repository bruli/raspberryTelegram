package status

import "github.com/bruli/raspberryRainSensor/pkg/common/vo"

type Status struct {
	humidity        int
	raining         bool
	systemStartedAt vo.Time
	temperature     int
	updatedAt       *vo.Time
}

func (s Status) Humidity() int {
	return s.humidity
}

func (s Status) Raining() bool {
	return s.raining
}

func (s Status) SystemStartedAt() vo.Time {
	return s.systemStartedAt
}

func (s Status) Temperature() int {
	return s.temperature
}

func (s Status) UpdatedAt() *vo.Time {
	return s.updatedAt
}

func (s *Status) Hydrate(humidity int, raining bool, startedAt vo.Time, temperature int, updatedAt *vo.Time) {
	s.humidity = humidity
	s.raining = raining
	s.systemStartedAt = startedAt
	s.temperature = temperature
	s.updatedAt = updatedAt
}
