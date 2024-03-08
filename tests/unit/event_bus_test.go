package tests

import (
	"testing"
	"time"

	"github.com/Michael-kyalo/sms_exchange/internal/sms/implementations"
	"github.com/Michael-kyalo/sms_exchange/internal/sms/interfaces"
)

func TestEventBus_StartBus(t *testing.T) {
	logger := &implementations.MockLogger{}
	eventBus := &implementations.MockEventBus{Logger: logger}

	err := eventBus.StartBus()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !logger.InfoCalled {
		t.Errorf("Expected Info method to be called")
	}

}
func TestEventBus_StopBus(t *testing.T) {
	logger := &implementations.MockLogger{}
	eventBus := &implementations.MockEventBus{Logger: logger}

	err := eventBus.StopBus()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !logger.InfoCalled {
		t.Errorf("Expected Info method to be called")
	}
}

func TestEventBus_PublishEvent(t *testing.T) {
	logger := &implementations.MockLogger{}
	eventBus := &implementations.MockEventBus{Logger: logger}

	eventBus.PublishEvent(interfaces.SmsSentEvent{
		PhoneNumber: "1233445",
		Timestamp:   time.Now(),
	})

	if !logger.InfoCalled {
		t.Errorf("Expected Info method to be called")
	}
}
