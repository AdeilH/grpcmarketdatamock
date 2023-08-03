package main

import (
	"encoding/json"
	"log"

	marketdata "github.com/adeilh/grpcmarketdatamock/proto" // Update with your package path
	"google.golang.org/protobuf/proto"
)

func main() {
	// Create a sample market data
	data := &marketdata.MarketData{
		Symbol:          "AAPL",
		BidPrice:        150.0,
		AskPrice:        151.0,
		LastTradedPrice: 150.5,
		OpenPrice:       149.0,
		HighPrice:       152.0,
		LowPrice:        148.0,
		ClosePrice:      150.0,
		Volume:          1000000,
		Exchange:        "NASDAQ",
		Timestamp:       1627868276,
	}

	// Marshal data using ProtoBuf
	protoBytes, err := proto.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to encode market data with ProtoBuf: %v", err)
	}
	protoSize := len(protoBytes)
	log.Printf("Size of ProtoBuf encoded data: %d bytes (%f kilobytes)", protoSize, float64(protoSize)/1024)

	// Marshal data using JSON
	jsonData := map[string]interface{}{
		"Symbol":          data.Symbol,
		"BidPrice":        data.BidPrice,
		"AskPrice":        data.AskPrice,
		"LastTradedPrice": data.LastTradedPrice,
		"OpenPrice":       data.OpenPrice,
		"HighPrice":       data.HighPrice,
		"LowPrice":        data.LowPrice,
		"ClosePrice":      data.ClosePrice,
		"Volume":          data.Volume,
		"Exchange":        data.Exchange,
		"Timestamp":       data.Timestamp,
	}
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatalf("Failed to encode market data with JSON: %v", err)
	}
	jsonSize := len(jsonBytes)
	log.Printf("Size of JSON encoded data: %d bytes (%f kilobytes)", jsonSize, float64(jsonSize)/1024)

	log.Printf("ProtoBuf is %d bytes (%f kilobytes) smaller than JSON", jsonSize-protoSize, float64(jsonSize-protoSize)/1024)
}
