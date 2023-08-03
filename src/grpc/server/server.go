package main

import (
	"context"
	"fmt"
	"log"
	"net"

	marketdata "github.com/adeilh/grpcmarketdatamock/proto" // Update with your package path
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type server struct {
	marketdata.UnimplementedMarketDataServiceServer
}

func (s *server) GetMarketData(ctx context.Context, req *marketdata.MarketDataRequest) (*marketdata.MarketDataResponse, error) {
	// Just returning some mock data
	data := &marketdata.MarketData{
		Symbol:          req.Symbol,
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

	encodedData, err := proto.Marshal(data)
	if err != nil {
		log.Fatalf("failed to encode data: %v", err)
	}
	size := len(encodedData)
	fmt.Println("Size of encoded data: ", size)
	return &marketdata.MarketDataResponse{Data: data}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	marketdata.RegisterMarketDataServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
