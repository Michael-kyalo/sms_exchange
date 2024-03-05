package sms

import (
	"log"
	"time"
)

type Handler struct {
	logger *log.Logger
}

// represents the command to send to an SMS
type SendSMSCommand struct {
	PhoneNumber string `json:"phoneNumber"`
	SmsText     string `json:"smsText"`
}

// represents an event indicating that an SMS has been sent
type SmsSentEvent struct {
	PhoneNumber string    `json:"phoneNumber"`
	Timestamp   time.Time `json:"timestamp"`
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

// Start starts the sms handler
func (h *Handler) Start() error {

	//TODO: start listening to the message queue
	return nil
}

// Stop stops the sms handler
func (h *Handler) Stop() error {

	//TODO: stop listening to the message queue
	return nil
}

// handleMessage handles incoming SendSMS commands.
func (h *Handler) handleMessage(command SendSMSCommand) error {
	//log the command that was received
	h.logger.Printf("Received SendSMSCommand command: %v", command)

	//send the SMS
	err := h.sendSMS(command.PhoneNumber, command.SmsText)
	if err != nil {

		//can additionally store the failed messages
		h.logger.Printf("Failed to send SMS command: %v", err)
		return err
	}

	h.publishEvent(SmsSentEvent{
		PhoneNumber: command.PhoneNumber,
		Timestamp:   time.Now().UTC(),
	})

	//TODO: handle incoming message
	return nil
}

func (h *Handler) sendSMS(phoneNumber string, smsText string) error {
	//TODO: send the SMS
	return nil
}

func (h *Handler) publishEvent(event SmsSentEvent) {
	//TODO: publish the event
}
