package commandline

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"temperature-converter/pkg/converter"
)

func Convert(in io.Reader) (int, error) {
	s := bufio.NewScanner(in)
	s.Split(bufio.ScanWords)

	fmt.Printf("\nPress C to convert from Fahrenheit to Celsius.\nPress F to convert from Celsius to Fahrenheit.\nYour choice: \n")
	var choice string

	if s.Scan() {
		choice = s.Text()
	}
	if err := s.Err(); err != nil {
		return 0, fmt.Errorf("unable to scan choice: %v", err)
	}

	var tempToConvert string
	fmt.Printf("Enter the temperature in Fahrenheit: ")
	if s.Scan() {
		tempToConvert = s.Text()
	}
	if err := s.Err(); err != nil {
		return 0, fmt.Errorf("unable to scan temperature to convert: %v", err)
	}

	tempNum, err := strconv.Atoi(tempToConvert)
	if err != nil {
		return 0, err
	}

	tempConverter := converter.New(tempNum)

	switch choice {
	case "F":
		return tempConverter.CelsiusToFahrenheit(), nil
	case "C":
		return tempConverter.FahrenheitToCelsius(), nil
	default:
		return 0, fmt.Errorf("incorrect choice %s", choice)
	}
}
