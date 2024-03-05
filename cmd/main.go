package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// //initialize logger
	// logger := log.New(os.Stdout, "[SMS EXCHANGE]", log.LstdFlags)

	// //initialize sms handler
	// smsHandler := sms.NewHandler(logger)

	// //start the sms_exchange microservice
	// err := smsHandler.Start()
	// if err != nil {
	// 	logger.Fatalf("Failed to start sms_exchange microservice: %v", err)
	// }

	// //handle graceful shutdown
	// sigChan := make(chan os.Signal, 1)

	// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// <-sigChan

	// logger.Println("Shutting down sms_exchange...")
	// smsHandler.Stop()
	// logger.Println("sms_exchange stopped")

	logger := log.New(os.Stdout, "[SMS EXCHANGE]", log.LstdFlags)

	// Initialize sms handler
	smsHandler := sms.NewHandler(logger)

	// Create a context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the sms_exchange microservice
	//in a seprate goroutine to avoid blocking the main thread
	go func() {
		if err := smsHandler.Start(ctx); err != nil {
			logger.Fatalf("Failed to start sms_exchange microservice: %v", err)
		}
	}()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigChan:
		logger.Printf("Received signal: %v", sig)
		cancel() // Cancel context to initiate graceful shutdown
	case <-ctx.Done():
		// Context canceled
	}

	logger.Println("Shutting down sms_exchange...")
	if err := smsHandler.Stop(); err != nil {
		logger.Fatalf("Error while stopping sms_exchange: %v", err)
	}
	logger.Println("sms_exchange stopped")

}
