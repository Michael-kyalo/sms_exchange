package tests

import (
	"github.com/Michael-kyalo/sms_exchange/internal/sms/implementations"

	"testing"
)

func TestLoggerInfo(t *testing.T) {
	mock := &implementations.MockLogger{
		InfoCalled: false,
	}
	message := "Test Info message"
	data := "Test Info data"

	mock.Info(message, data)

	if !mock.InfoCalled {
		t.Errorf("Info method not called")
	}

	if mock.InfoMessage != message {
		t.Errorf("Expected message '%s', got '%s'", message, mock.InfoMessage)
	}

	if mock.InfoData != data {
		t.Errorf("Expected data '%v', got '%v'", data, mock.InfoData)
	}
}

func TestLoggerError(t *testing.T) {
	mock := &implementations.MockLogger{
		InfoCalled: false,
	}
	message := "Test Error message"
	data := "Test Error data"

	mock.Error(message, data)

	if !mock.ErrorCalled {
		t.Errorf("Error method not called")
	}

	if mock.ErrorMessage != message {
		t.Errorf("Expected message '%s', got '%s'", message, mock.ErrorMessage)
	}

	if mock.ErrorData != data {
		t.Errorf("Expected data '%v', got '%v'", data, mock.ErrorData)
	}
}
