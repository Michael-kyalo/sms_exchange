package implementations

import (
	"github.com/Michael-kyalo/sms_exchange/internal/sms/interfaces"
)

type MockHttpClient struct {
	logger interfaces.Logger
}

func NewMockHttpClient() *MockHttpClient {
	mock := &MockHttpClient{
		logger: &MockLogger{},
	}
	return mock
}

func (m *MockHttpClient) Get(url string) <-chan *interfaces.HTTPResponse {
	respChan := make(chan *interfaces.HTTPResponse, 1)

	// Simulate a successful response
	go func() {
		respChan <- &interfaces.HTTPResponse{
			StatusCode: 200,
			Body:       []byte("Mock GET response"),
			Error:      nil,
		}
	}()
	return respChan
}

func (m *MockHttpClient) Post(url string, body []byte) <-chan *interfaces.HTTPResponse {

	respChan := make(chan *interfaces.HTTPResponse, 1)

	go func() {
		//simulate successful response
		respChan <- &interfaces.HTTPResponse{
			StatusCode: 200,
			Body:       []byte("Mock POST response"),
			Error:      nil,
		}
	}()

	return respChan
}
