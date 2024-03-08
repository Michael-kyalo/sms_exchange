package interfaces

// defines the interface for interacting with a message queue.
type MessageQueue interface {
	Logger
	//Start starts the message queue
	StartMessageQueue() error
	//Stop stops the message queue
	StopMessageQueue() error
	//RecieveCommand handles incoming SendSMS commands
	RecieveCommand() (*SendSMSCommand, error)
}

// represents the command to send to an SMS
type SendSMSCommand struct {
	PhoneNumber string `json:"phoneNumber"`
	SmsText     string `json:"smsText"`
}
