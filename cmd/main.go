package main

import (
	"github.com/Michael-kyalo/sms_exchange/internal/sms/implementations"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Initialize environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize logger
	logger := &implementations.MockLogger{}

	// Initialize event bus, message queue, and HTTP client
	eventBus := &implementations.MockEventBus{
		Logger: logger,
	}
	messageQueue := &implementations.MockMessageQueue{
		Logger: logger,
	} // Implement MessageQueue interface
	httpClient := &implementations.MockHttpClient{} // Implement HTTPClient interface

	// Create a new instance of the HandlerImpl with initialized dependencies
	handler := implementations.NewHandler(logger, eventBus, messageQueue, httpClient)

	// Start the SMS exchange microservice
	go func() {
		if err := handler.Start(); err != nil {
			logger.Error("Failed to start SMS exchange microservice: %v", err)
		}
	}()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a termination signal
	sig := <-sigChan
	logger.Info("Received signal: %v", sig)

	// Stop the SMS exchange microservice
	if err := handler.Stop(); err != nil {
		logger.Error("Error while stopping SMS exchange: %v", err)
	}
	logger.Info("SMS exchange stopped", "success")
}
