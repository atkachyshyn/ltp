package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHandleLTP tests the LTP handler
func TestHandleLTP(t *testing.T) {
	// Create a server to mock Kraken API responses
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := `{"error":[],"result":{"XXBTZUSD":{"c":["52000.12","1"]},"XXBTZEUR":{"c":["50000.12","1"]}}}`
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer ts.Close()

	// Create the handler with the mocked server URL
	handler := HandleLTP(ts.URL)

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/api/v1/ltp", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(handler)

	// Dispatch the request
	handlerFunc.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := `{"ltp":[{"pair":"BTC/USD","amount":"52000.12"},{"pair":"BTC/EUR","amount":"50000.12"}]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
