package main

import (
	"context"
	"log"
	"net"
	"time"

	marketdata "github.com/adeilh/grpcmarketdatamock/proto" // update with your package path
	"google.golang.org/grpc"
)

type server struct {
	marketdata.UnimplementedMarketDataServiceServer
}

func (s *server) GetMarketData(ctx context.Context, req *marketdata.MarketDataRequest) (*marketdata.MarketDataResponse, error) {
	data := &marketdata.MarketData{
		Symbol:          req.Symbol,
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
	return &marketdata.MarketDataResponse{Data: data}, nil
}

func startServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	marketdata.RegisterMarketDataServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startClient() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := marketdata.NewMarketDataServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	start := time.Now()
	for i := 0; i < 10000; i++ {
		_, err := c.GetMarketData(ctx, &marketdata.MarketDataRequest{Symbol: "AAPL"})
		if err != nil {
			log.Fatalf("could not get market data: %v", err)
		}
	}
	elapsed := time.Since(start)
	// log.Printf("Data: %v", r)
	log.Printf("Time taken: %v", elapsed)
}

func main() {
	go startServer()
	time.Sleep(2 * time.Second)
	go startClient()
	time.Sleep(2 * time.Second)
}
