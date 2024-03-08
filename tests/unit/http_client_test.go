package tests

import (
	"testing"

	"github.com/Michael-kyalo/sms_exchange/internal/sms/implementations"
)

func TestMockHttpClient_Get(t *testing.T) {
	client := implementations.NewMockHttpClient()
	url := "http://example.com"

	resp := <-client.Get(url)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
	expectBody := "Mock GET response"

	if string(resp.Body) != expectBody {
		t.Errorf("Expected body '%s', got '%s'", expectBody, string(resp.Body))
	}
}

func TestMockHTTPClient_Post(t *testing.T) {
	client := implementations.NewMockHttpClient()
	url := "http://example.com"
	body := []byte("test body")

	resp := <-client.Post(url, body)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 201, got %d", resp.StatusCode)
	}

	expectedBody := "Mock POST response"
	if string(resp.Body) != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, resp.Body)
	}
}
