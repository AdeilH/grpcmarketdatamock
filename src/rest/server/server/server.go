package server

import (
	"encoding/json"
	"fmt"
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

func GetMarketData(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")

	// Just returning some mock market data for the given symbol
	data := MarketData{
		Symbol:          symbol,
		BidPrice:        100.0,
		AskPrice:        101.0,
		LastTradedPrice: 102.0,
		OpenPrice:       98.0,
		HighPrice:       105.0,
		LowPrice:        95.0,
		ClosePrice:      103.0,
		Volume:          500000,
		Exchange:        "NASDAQ",
		Timestamp:       1628114400, // Example timestamp
	}

	encodedData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		return
	}

	// Get the size of the encoded data
	size := len(encodedData)
	fmt.Println("Size of encoded data: ", size)

	w.Write(encodedData)
}
