package main

import (
	"fmt"
	"log"
	"os"

	"temperature-converter/pkg/commandline"
)

func main() {
	temperature, err := commandline.Convert(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(temperature)
}
