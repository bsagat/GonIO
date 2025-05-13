package main

import (
	"GonIO/internal/app"
	"GonIO/internal/domain"
	"log"
	"net/http"
)

func main() {
	mux := app.SetMux()

	log.Println("Starting listening...")
	if err := http.ListenAndServe(domain.URLDomain+":"+domain.Port, mux); err != nil {
		log.Fatalf("Server listening error: %s", err.Error())
	}
}
