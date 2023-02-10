package main

import (
	"flag"

	"wallet-service"
	"wallet-service/service"
)

var broker = flag.String("broker", "localhost:9092", "bootstrap Kafka broker")

func main() {
	flag.Parse()
	service.Run([]string{*broker}, wallet.DepositStream)
}
