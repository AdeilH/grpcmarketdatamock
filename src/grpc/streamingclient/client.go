package main

import (
	"context"
	"io"
	"log"

	marketdata "github.com/adeilh/grpcmarketdatamock/proto" // Update with your package path
	"google.golang.org/grpc"
)

func main() {
	// Dial the server
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := marketdata.NewMarketDataServiceClient(conn)

	// Request streaming updates for a specific symbol
	req := &marketdata.MarketDataRequest{Symbol: "AAPL"}
	stream, err := client.StreamMarketData(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to stream market data: %v", err)
	}

	// Receive updates until the stream is closed
	for {
		update, err := stream.Recv()
		if err == io.EOF {
			break // Stream closed
		}
		if err != nil {
			log.Fatalf("failed to receive update: %v", err)
		}
		log.Printf("Received update: %v\n", update)
	}
}
