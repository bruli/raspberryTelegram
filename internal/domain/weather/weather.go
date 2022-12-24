package weather

type Weather struct {
	humidity, temperature int
	raining               bool
}

func (w Weather) Humidity() int {
	return w.humidity
}

func (w Weather) Temperature() int {
	return w.temperature
}

func (w Weather) IsRaining() bool {
	return w.raining
}

func (w *Weather) Hydrate(temperature, humidity int, raining bool) {
	w.temperature = temperature
	w.humidity = humidity
	w.raining = raining
}
