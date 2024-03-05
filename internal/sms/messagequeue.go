package sms

import (
	"encoding/json"
	"log"
)

// represents a concrete implentation of message queue for receiving commands
type MessageQueue struct {
	logger *log.Logger
}

// defines the interface for interacting with a message queue.
type MessageQueueInterface interface {
	RecieveCommand() (*SendSMSCommand, error)
}

// creates a new MessageQueue instance
func NewMessageQueue(logger *log.Logger) *MessageQueue {
	return &MessageQueue{
		logger: logger,
	}
}

// starts a listening to the message queue for sending commands
func (mq *MessageQueue) Start() error {
	//TODO: start listening to the message queue
	return nil
}

// stops listening to the message queue
func (mq *MessageQueue) Stop() error {
	//TODO: stop listening to the message queue
	return nil
}

// RecieveCommand handles incoming SendSMS commands.

func (mq *MessageQueue) RecieveCommand() (*SendSMSCommand, error) {

	//simulating recieve command from the message queue
	commandBytes := []byte(`{"PhoneNumber": "123456789", "SmsText": "Test SMS"}`)

	// Unmarshal the received message into a SendSmsCommand
	var command SendSMSCommand
	if err := json.Unmarshal(commandBytes, &command); err != nil {
		return nil, err
	}

	mq.logger.Printf("Received command from message queue: %+v", command)

	return &command, nil

}
