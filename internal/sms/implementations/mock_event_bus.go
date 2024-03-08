package implementations

import "github.com/Michael-kyalo/sms_exchange/internal/sms/interfaces"

// EventBusImpl represents the concrete implementation of the EventBus.
type MockEventBus struct {
	Logger         interfaces.Logger
	Started        bool
	Stopped        bool
	EventPublished bool
}

// creates a new EventBus instance
func NewMockEventBus() *MockEventBus {
	return &MockEventBus{
		Logger:         &MockLogger{},
		EventPublished: true,
	}
}

// starts a new EventBus
func (m *MockEventBus) StartBus() error {
	m.Started = true
	m.Logger.Info("EventBus started", "Starting")
	return nil
}

func (m *MockEventBus) StopBus() error {
	m.Stopped = true
	m.Logger.Info("EventBus stopped", "Stopping")
	return nil
}

func (m *MockEventBus) PublishEvent(event interfaces.SmsSentEvent) {
	m.EventPublished = true
	m.Logger.Info("Publishing event", "Event published")
}

func (m *MockEventBus) Reset() {
	m.Started = false
	m.Stopped = false
	m.EventPublished = false
}
