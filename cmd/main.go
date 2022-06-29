package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	fmt.Printf("\nPress C to convert from Fahrenheit to Celsius.\nPress F to convert from Celsius to Fahrenheit.\nYour choice: \n")
	var choice string

	if s.Scan() {
		choice = s.Text()
	}
	if err := s.Err(); err != nil {
		log.Fatalf("unable to scan choice: %v", err)
	}

	var tempToConvert string
	fmt.Printf("Enter the temperature in Fahrenheit: ")
	if s.Scan() {
		tempToConvert = s.Text()
	}
	if err := s.Err(); err != nil {
		log.Fatalf("unable to scan temperature to convert: %v", err)
	}

	tempNum, err := strconv.Atoi(tempToConvert)
	if err != nil {
		log.Fatal(err)
	}

	switch choice {
	case "F":
		fahrenheit := (tempNum * (9 / 5)) + 32
		fmt.Println(fahrenheit)
	case "C":
		celsius := (tempNum - 32) * (5 / 9)
		fmt.Println(celsius)
	default:
		fmt.Println("incorrect choice")
	}
}
