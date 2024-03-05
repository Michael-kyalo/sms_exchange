package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// represents client for interacting with third-party SMS service providers
type HTTPClient struct {
	baseURL string
	client  *http.Client
}

// create a new instance of the HTTPClient
func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// SendSMS sends an SMS to the specified phone number via a third-party SMS service provider
func (c *HTTPClient) SendSMS(phoneNumber, smsText string) error {
	requestPayload := map[string]string{
		"PhoneNumber": phoneNumber,
		"SmsText":     smsText,
	}

	requestBytes, err := json.Marshal(requestPayload)
	if err != nil {
		return err
	}

	response, err := c.client.Post(c.baseURL, "application/json", bytes.NewBuffer(requestBytes))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)

	}

	return nil
}
