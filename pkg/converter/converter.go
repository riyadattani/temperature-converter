package converter

import "fmt"

type TempConverter struct {
	temp   int
	choice string
}

func New(choice string, temp int) *TempConverter {
	return &TempConverter{choice: choice, temp: temp}
}

func (t TempConverter) Convert() (int, error) {
	switch t.choice {
	case "F":
		return t.celsiusToFahrenheit(), nil
	case "C":
		return t.fahrenheitToCelsius(), nil
	default:
		// todo: test me
		return 0, fmt.Errorf("invalid choice %s", t.choice)
	}
}

func (t TempConverter) celsiusToFahrenheit() int {
	return (t.temp * (9 / 5)) + 32
}

func (t TempConverter) fahrenheitToCelsius() int {
	return (t.temp - 32) * (5 / 9)
}
