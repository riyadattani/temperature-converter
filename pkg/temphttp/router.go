package temphttp

import (
	"fmt"
	"net/http"
	"strconv"
)

type TempConverter interface {
	Convert(choice string, temp int) (int, error)
}

type server struct {
	tempConverter TempConverter
}

func NewRouter(converterService TempConverter) http.Handler {
	svr := server{tempConverter: converterService}

	mux := http.NewServeMux()
	mux.HandleFunc("/temp-converter", svr.converter)
	return mux
}

func (s server) converter(w http.ResponseWriter, r *http.Request) {
	choice := r.URL.Query().Get("choice")
	if choice == "" {
		http.Error(w, "must have a choice", http.StatusBadRequest)
		return
	}
	temp := r.URL.Query().Get("temp")
	if temp == "" {
		http.Error(w, "must have a temp", http.StatusBadRequest)
		return
	}

	t, _ := strconv.Atoi(temp)

	convertedTemp, _ := s.tempConverter.Convert(choice, t)

	fmt.Fprintf(w, "%v", convertedTemp)
}
