package main

import (
	"log"
	"net/http"

	"github.com/atkachyshyn/ltp/internal/api"
	"github.com/atkachyshyn/ltp/internal/config"
)

func main() {
	cfg, err := config.LoadConfig("./config/config.yml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	http.HandleFunc("/api/v1/ltp", api.HandleLTP(cfg.KrakenAPIURL))
	log.Printf("Server starting on %s", cfg.ServerAddress)
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, nil))
}
