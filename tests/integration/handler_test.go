package tests

import (
	"testing"

	"github.com/Michael-kyalo/sms_exchange/internal/sms/implementations"
)

func TestHandler_Start(t *testing.T) {
	logger := &implementations.MockLogger{}
	queue := &implementations.MockMessageQueue{}
	bus := &implementations.MockEventBus{Logger: logger}
	client := &implementations.MockHttpClient{}

	handler := implementations.NewHandler(
		logger,
		bus,
		queue,
		client,
	)

	err := handler.Start()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !logger.InfoCalled {
		t.Errorf("Expected Info method to be called")
	}
}
