package main

import (
	"log"
	"net/http"

	"temperature-converter/pkg/converter"
	"temperature-converter/pkg/temphttp"
)

func main() {
	tempConverter := converter.New()
	router := temphttp.NewRouter(tempConverter)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
