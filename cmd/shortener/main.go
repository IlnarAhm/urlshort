package main

import (
	"net/http"

	"github.com/IlnarAhm/urlshort/internal/config"
	"github.com/IlnarAhm/urlshort/internal/logger"
)

func main() {
	// init config
	config := config.New()

	// init logger
	logger := logger.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
