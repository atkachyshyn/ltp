package config

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestLoadConfig tests the loading of the config file
func TestLoadConfig(t *testing.T) {
	// Create a temporary config file
	content := []byte("serverAddress: \":8080\"\nkrakenAPIURL: \"https://api.kraken.com/0/public/Ticker?pair=XXBTZUSD,XXBTZEUR\"")
	tmpfile, err := ioutil.TempFile("", "config.*.yml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Load the config
	config, err := LoadConfig(tmpfile.Name())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check the values
	if config.ServerAddress != ":8080" {
		t.Errorf("Expected server address ':8080', got '%s'", config.ServerAddress)
	}
	if config.KrakenAPIURL != "https://api.kraken.com/0/public/Ticker?pair=XXBTZUSD,XXBTZEUR" {
		t.Errorf("Expected Kraken API URL 'https://api.kraken.com/0/public/Ticker?pair=XXBTZUSD,XXBTZEUR', got '%s'", config.KrakenAPIURL)
	}
}
