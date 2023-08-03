package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

func main() {
	symbol := "AAPL"
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/getmarketdata?symbol=%s", symbol))
	if err != nil {
		log.Fatalf("could not get market data: %v", err)
	}
	defer resp.Body.Close()

	var data MarketData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("could not decode response: %v", err)
	}
	fmt.Printf("Market Data: %+v\n", data)
}
