package implementations

import "github.com/Michael-kyalo/sms_exchange/internal/sms/interfaces"

// MockMessageQueue is a mock implementation of the MessageQueue interface
type MockMessageQueue struct {
	Logger         interfaces.Logger
	Started        bool
	Stopped        bool
	ReceiveInvoked bool
}

// NewMockMessageQueue creates a new mock MessageQueue
func NewMockMessageQueue() *MockMessageQueue {
	return &MockMessageQueue{
		Started:        false,
		Stopped:        false,
		ReceiveInvoked: false,
	}
}

func (m *MockMessageQueue) StartMessageQueue() error {
	m.Started = true
	return nil
}

func (m *MockMessageQueue) StopMessageQueue() error {
	m.Stopped = true
	return nil
}

func (m *MockMessageQueue) RecieveCommand() (*interfaces.SendSMSCommand, error) {
	m.ReceiveInvoked = true
	return &interfaces.SendSMSCommand{
		PhoneNumber: "123445",
		SmsText:     "Hello World",
	}, nil
}
