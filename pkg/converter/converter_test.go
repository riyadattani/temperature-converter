package converter_test

import (
	"testing"

	"temperature-converter/pkg/converter"

	"github.com/alecthomas/assert/v2"
)

func TestTemperatureConverter_CelsiusToFahrenheit(t *testing.T) {
	t.Run("convert celsius to fahrenheit", func(t *testing.T) {
		tempConverter := converter.New()
		temp, err := tempConverter.Convert("F", 32)
		assert.NoError(t, err)
		assert.Equal(t, 64, temp)
	})

	t.Run("convert fahrenheit to celsius", func(t *testing.T) {
		tempConverter := converter.New()
		temp, err := tempConverter.Convert("C", 32)
		assert.NoError(t, err)
		assert.Equal(t, 0, temp)
	})
}
