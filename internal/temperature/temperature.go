package temperature

type Temperature struct {
	humidity, temperature float32
}

func (t Temperature) Temperature() float32 {
	return t.temperature
}

func (t Temperature) Humidity() float32 {
	return t.humidity
}

func NewTemperature(humidity float32, temperature float32) *Temperature {
	return &Temperature{humidity: humidity, temperature: temperature}
}
