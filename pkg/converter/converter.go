package converter

type TempConverter struct {
	temp int
}

func New(temp int) *TempConverter {
	return &TempConverter{temp: temp}
}

func (t TempConverter) CelsiusToFahrenheit() int {
	return (t.temp * (9 / 5)) + 32
}

func (t TempConverter) FahrenheitToCelsius() int {
	return (t.temp - 32) * (5 / 9)
}
