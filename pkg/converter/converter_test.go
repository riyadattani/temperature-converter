package converter_test

import (
	"testing"

	"temperature-converter/pkg/converter"

	"github.com/alecthomas/assert/v2"
)

func TestTemperatureConverter_CelsiusToFahrenheit(t *testing.T) {
	t.Run("convert celsius to fahrenheit", func(t *testing.T) {
		tempConverter := converter.New(32)
		assert.Equal(t, 64, tempConverter.CelsiusToFahrenheit())
	})

	t.Run("convert fahrenheit to celsius", func(t *testing.T) {
		tempConverter := converter.New(32)
		assert.Equal(t, 0, tempConverter.FahrenheitToCelsius())
	})
}
