package commandline

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
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

	switch choice {
	case "F":
		fahrenheit := (tempNum * (9 / 5)) + 32
		return fahrenheit, nil
	case "C":
		celsius := (tempNum - 32) * (5 / 9)
		return celsius, nil
	default:
		return 0, fmt.Errorf("incorrect choice %s", choice)
	}
}
