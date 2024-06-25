package main

import (
	"WeekendPOS/app/config"
	"WeekendPOS/app/delivery/messaging"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	logger.Info("Starting worker service")

	ctx, cancel := context.WithCancel(context.Background())

	logger.Info("setup user consumer")
	userConsumer := config.NewKafkaConsumer(viperConfig, logger)
	userHandler := messaging.NewUserConsumer(logger)
	go messaging.ConsumeTopic(ctx, userConsumer, "users", logger, userHandler.Consume)

	logger.Info("setup category consumer")
	categoryConsumer := config.NewKafkaConsumer(viperConfig, logger)
	categoryHandler := messaging.NewCategoryConsumer(logger)
	go messaging.ConsumeTopic(ctx, categoryConsumer, "categories", logger, categoryHandler.Consume)

	logger.Info("Worker is running")

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	stop := false
	for !stop {
		select {
		case s := <-terminateSignals:
			logger.Info("Got one of stop signals, shutting down worker gracefully, SIGNAL NAME :", s)
			cancel()
			stop = true
		}
	}

	time.Sleep(5 * time.Second)
}
