package interfaces

import (
	"time"
)

// Handler handles all the requests and events
type Handler interface {
	Logger
	EventBus
	MessageQueue
	HTTPClient
	//Start starts the handler
	Start() error
	//Stop stops the handler
	Stop() error

	HandleMessage(command SendSMSCommand) error
}

// represents an event indicating that an SMS has been sent
type SmsSentEvent struct {
	PhoneNumber string    `json:"phoneNumber"`
	Timestamp   time.Time `json:"timestamp"`
}
