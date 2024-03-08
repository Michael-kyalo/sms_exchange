package interfaces

// represents an event bus for publishing events
type EventBus interface {

	//Start start handles the logic for connecting to the event bus
	StartBus() error
	//Stop stops the event bus
	StopBus() error
	//PublishEvent publishes an event to event bus
	PublishEvent(event SmsSentEvent)
}
