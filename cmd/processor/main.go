package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"wallet-service"

	"wallet-service/above_threshold"
	"wallet-service/balance"
	"wallet-service/topic_init"

	"golang.org/x/sync/errgroup"
)

var (
	brokers           = []string{"localhost:9092"}
	runBalance        = flag.Bool("balance", false, "run balance processor")
	runAboveThreshold = flag.Bool("above_threshold", false, "run above_threshold processor")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)

	if *runBalance {
		balance.PrepareTopics(brokers)
	}
	if *runAboveThreshold {
		above_threshold.PrepareTopics(brokers)
	}

	if *runBalance {
		log.Println("starting balance processor")
		grp.Go(balance.Run(ctx, brokers))
	}
	if *runAboveThreshold {
		log.Println("starting above_threshold processor")
		grp.Go(above_threshold.Run(ctx, brokers))
	}

	// Wait for SIGINT/SIGTERM
	waiter := make(chan os.Signal, 1)
	signal.Notify(waiter, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-waiter:
	case <-ctx.Done():
	}
	cancel()
	if err := grp.Wait(); err != nil {
		log.Println(err)
	}
	log.Println("done")
}
