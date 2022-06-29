package converter

import "fmt"

type TempConverter struct {
}

func New() *TempConverter {
	return &TempConverter{}
}

func (t TempConverter) Convert(choice string, temp int) (int, error) {
	switch choice {
	case "F":
		return t.celsiusToFahrenheit(temp), nil
	case "C":
		return t.fahrenheitToCelsius(temp), nil
	default:
		// todo: test me
		return 0, fmt.Errorf("invalid choice %s", choice)
	}
}

func (t TempConverter) celsiusToFahrenheit(temp int) int {
	return (temp * (9 / 5)) + 32
}

func (t TempConverter) fahrenheitToCelsius(temp int) int {
	return (temp - 32) * (5 / 9)
}
