package main

import (
	"log"
	"net"
	"time"

	marketdata "github.com/adeilh/grpcmarketdatamock/proto" // Update with your package path
	"google.golang.org/grpc"
)

type server struct {
	marketdata.UnimplementedMarketDataServiceServer
}

func (s *server) StreamMarketData(req *marketdata.MarketDataRequest, stream marketdata.MarketDataService_StreamMarketDataServer) error {
	// Simulate streaming updates for a specific symbol
	for {
		bidUpdate := &marketdata.BidUpdate{
			Symbol:    req.Symbol,
			BidPrice:  100 + float64(time.Now().Unix()%10),
			BidVolume: 1000,
			Timestamp: time.Now().Unix(),
		}

		if err := stream.Send(&marketdata.MarketDataUpdate{Symbol: req.Symbol, Update: &marketdata.MarketDataUpdate_BidUpdate{BidUpdate: bidUpdate}}); err != nil {
			return err
		}
		time.Sleep(time.Second) // Pause between updates
		askUpdate := &marketdata.AskUpdate{
			Symbol:    req.Symbol,
			AskPrice:  100 + float64(time.Now().Unix()%10),
			AskVolume: 1000,
			Timestamp: time.Now().Unix(),
		}

		if err := stream.Send(&marketdata.MarketDataUpdate{Symbol: req.Symbol, Update: &marketdata.MarketDataUpdate_AskUpdate{AskUpdate: askUpdate}}); err != nil {
			return err
		}
		time.Sleep(time.Second) // Pause between updates
	}
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
