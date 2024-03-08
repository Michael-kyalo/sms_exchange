package implementations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Michael-kyalo/sms_exchange/internal/sms/interfaces"
)

type HandlerImpl struct {
	logger interfaces.Logger
	bus    interfaces.EventBus
	queue  interfaces.MessageQueue
	client interfaces.HTTPClient
}

func NewHandler(logger interfaces.Logger, bus interfaces.EventBus, queue interfaces.MessageQueue, client interfaces.HTTPClient) *HandlerImpl {
	return &HandlerImpl{
		logger: logger,
		bus:    bus,
		queue:  queue,
		client: client,
	}
}

// Start starts the sms HandlerImpl
func (h *HandlerImpl) Start() error {

	if err := h.queue.StartMessageQueue(); err != nil {
		h.logger.Info("Failed to start message queue: %v", err)
		return err
	}
	if err := h.bus.StartBus(); err != nil {
		h.logger.Info("Failed to start event bus: %v", err)
		return err
	}

	command, err := h.queue.RecieveCommand()
	if err != nil {
		h.logger.Info("Failed to Recieve message: %v", err)
		return err
	}
	h.logger.Info("Recieved message: %v", command)

	if err := h.HandleMessage(command); err != nil {
		h.logger.Info("Failed to handle message: %v", err)
		return err
	}

	return nil

}

// Stop stops the sms HandlerImpl
func (h *HandlerImpl) Stop() error {

	if err := h.queue.StopMessageQueue(); err != nil {
		h.logger.Info("Failed to stop message queue: %v", err)
		return err
	}
	if err := h.bus.StopBus(); err != nil {
		h.logger.Info("Failed to stop event bus: %v", err)
		return err
	}
	return nil
}

// handleMessage handles incoming SendSMS commands.
func (h *HandlerImpl) HandleMessage(command *interfaces.SendSMSCommand) error {
	//log the command that was received
	h.logger.Info("Received SendSMSCommand command: %v", command)

	//send the SMS
	err := h.sendSMS(command.PhoneNumber, command.SmsText)
	if err != nil {

		//can additionally store the failed messages in a database
		//h.logger.Info("Failed to send SMS command: %v", err)
		return err
	}

	h.publishEvent(interfaces.SmsSentEvent{
		PhoneNumber: command.PhoneNumber,
		Timestamp:   time.Now().UTC(),
	})

	return nil
}

func (h *HandlerImpl) sendSMS(phoneNumber, smsText string) error {
	//initialize the HTTP client with third party url
	thirdPartyURL := os.Getenv("THIRD_PARTY_URL")
	requestPayload := map[string]string{
		"PhoneNumber": phoneNumber,
		"SmsText":     smsText,
	}

	requestBytes, err := json.Marshal(requestPayload)
	if err != nil {
		return err
	}

	response := <-h.client.Post(thirdPartyURL, requestBytes)
	if response.Error != nil {
		return response.Error
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)

	}

	h.publishEvent(interfaces.SmsSentEvent{
		PhoneNumber: phoneNumber,
		Timestamp:   time.Now().UTC(),
	})

	return nil
}

func (h *HandlerImpl) publishEvent(event interfaces.SmsSentEvent) {
	h.bus.PublishEvent(event)
}
