package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// MarketData struct to represent market data
type MarketData struct {
	Symbol          string  `json:"symbol"`
	BidPrice        float64 `json:"bid_price"`
	AskPrice        float64 `json:"ask_price"`
	LastTradedPrice float64 `json:"last_traded_price"`
	OpenPrice       float64 `json:"open_price"`
	HighPrice       float64 `json:"high_price"`
	LowPrice        float64 `json:"low_price"`
	ClosePrice      float64 `json:"close_price"`
	Volume          int64   `json:"volume"`
	Exchange        string  `json:"exchange"`
	Timestamp       int64   `json:"timestamp"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := &MarketData{
		Symbol:          "AAPL",
		BidPrice:        100.0,
		AskPrice:        101.0,
		LastTradedPrice: 150.5,
		OpenPrice:       149.0,
		HighPrice:       152.0,
		LowPrice:        148.0,
		ClosePrice:      150.0,
		Volume:          1000000,
		Exchange:        "NASDAQ",
		Timestamp:       1627868276,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func startServer() {
	http.HandleFunc("/marketdata", handler)
	log.Printf("server listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func startClient() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		resp, err := http.Get("http://localhost:8080/marketdata")
		if err != nil {
			log.Fatalf("Error fetching data: %v", err)
		}
		defer resp.Body.Close()

		data := &MarketData{}
		if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
			log.Fatalf("Error decoding response: %v", err)
		}
	}
	elapsed := time.Since(start)
	// log.Printf("Data: %v", data)
	log.Printf("Time taken: %v", elapsed)
}

func main() {
	// Start server
	go startServer()

	// Give some time for the server to start
	time.Sleep(2 * time.Second)

	// Start client
	go startClient()

	// Wait for client to complete
	time.Sleep(2 * time.Second)
}
