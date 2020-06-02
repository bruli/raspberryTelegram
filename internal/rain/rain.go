package rain

type Rain struct {
	IsRain bool
	Value  uint16
}

func New(isRain bool, value uint16) *Rain {
	return &Rain{IsRain: isRain, Value: value}
}
