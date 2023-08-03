package main

import (
	"log"
	"net/http"

	"github.com/adeilh/grpcmarketdatamock/src/rest/server/server"
)

func main() {
	http.HandleFunc("/getmarketdata", server.GetMarketData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
