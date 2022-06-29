package main

import (
	"fmt"
	"log"
	"os"

	"temperature-converter/pkg/commandline"
	"temperature-converter/pkg/converter"
)

func main() {
	tempConverter := converter.New()
	temperature, err := commandline.Convert(os.Stdin, tempConverter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(temperature)
}
