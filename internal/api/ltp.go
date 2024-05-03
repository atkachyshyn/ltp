package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// TickerResponse structure to unmarshal JSON data from Kraken
type TickerResponse struct {
	Error  []string              `json:"error"`
	Result map[string]TickerData `json:"result"`
}

// TickerData structure for nested result
type TickerData struct {
	LastTradeClosed []string `json:"c"`
}

// LTP structure to send final response
type LTP struct {
	Pair   string `json:"pair"`
	Amount string `json:"amount"`
}

// HandleLTP retrieves and sends the LTP data
func HandleLTP(krakenAPIURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(krakenAPIURL)
		if err != nil {
			http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}

		var apiResp TickerResponse
		if err := json.Unmarshal(body, &apiResp); err != nil {
			http.Error(w, "Failed to parse JSON", http.StatusInternalServerError)
			return
		}

		ltps := make([]LTP, 0)
		for key, data := range apiResp.Result {
			var pair string
			switch key {
			case "XXBTZUSD":
				pair = "BTC/USD"
			case "XXBTZEUR":
				pair = "BTC/EUR"
			case "XXBTZCHF":
				pair = "BTC/CHF"
			}
			ltps = append(ltps, LTP{Pair: pair, Amount: data.LastTradeClosed[0]})
		}

		response, _ := json.Marshal(map[string][]LTP{"ltp": ltps})
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}
