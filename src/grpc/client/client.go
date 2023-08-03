package main

import (
	"context"
	"log"
	"time"

	marketdata "github.com/adeilh/grpcmarketdatamock/proto" // Update with your package path
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := marketdata.NewMarketDataServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &marketdata.MarketDataRequest{Symbol: "AAPL"}
	r, err := c.GetMarketData(ctx, req)
	if err != nil {
		log.Fatalf("could not get market data: %v", err)
	}
	log.Printf("Market Data: %v", r.GetData())
}
