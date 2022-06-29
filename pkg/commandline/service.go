package commandline

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type TempConverter interface {
	Convert(choice string, temp int) (int, error)
}

func Convert(in io.Reader, service TempConverter) (int, error) {
	s := bufio.NewScanner(in)
	s.Split(bufio.ScanWords)

	readWord := func() string {
		s.Scan()
		return s.Text()
	}

	fmt.Printf("\nPress C to convert from Fahrenheit to Celsius.\nPress F to convert from Celsius to Fahrenheit.\nYour choice: \n")
	choice := readWord()

	fmt.Printf("Enter the temperature in Fahrenheit: ")
	tempToConvert := readWord()

	tempNum, err := strconv.Atoi(tempToConvert)
	if err != nil {
		return 0, err
	}

	return service.Convert(choice, tempNum)
}
