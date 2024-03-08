package tests

import (
	"testing"

	"github.com/Michael-kyalo/sms_exchange/internal/sms/implementations"
)

func TestMockMessageQueue_StartMessageQueue(t *testing.T) {
	queue := implementations.NewMockMessageQueue()

	err := queue.StartMessageQueue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !queue.Started {
		t.Error("Expected StartMessageQueue to be called")
	}
}

func TestMockMessageQueue_StopMessageQueue(t *testing.T) {
	queue := implementations.NewMockMessageQueue()

	err := queue.StopMessageQueue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !queue.Stopped {
		t.Error("Expected StopMessageQueue to be called")
	}
}

func TestMockMessageQueue_RecieveCommand(t *testing.T) {
	queue := implementations.NewMockMessageQueue()

	_, err := queue.RecieveCommand()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !queue.ReceiveInvoked {
		t.Error("Expected RecieveCommand to be called")
	}
}
