package channels

import (
	"context"
	"errors"
	"testing"

	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// Define the Twilio client interface
type TwilioClient interface {
	CreateMessage(params *twilioApi.CreateMessageParams) (string, string, error)
}

// MockTwilioClient is a mock implementation of the Twilio client interface.
type MockTwilioClient struct {
	// createMessageFunc holds a function that simulates the CreateMessage behavior.
	createMessageFunc func(params *twilioApi.CreateMessageParams) (string, string, error)
}

func (m *MockTwilioClient) CreateMessage(params *twilioApi.CreateMessageParams) (string, string, error) {
	// Call the mock function to simulate behavior.
	return m.createMessageFunc(params)
}

func TestSMSChannel_SendNotification_Success(t *testing.T) {
	// Create a mock client with a function that returns success.
	mockClient := &MockTwilioClient{
		createMessageFunc: func(params *twilioApi.CreateMessageParams) (string, string, error) {
			return "", "", nil
		},
	}

	// Create an SMSChannel instance with the mock client.
	channel := &SMSChannel{
		accountSid: "test_account_sid",
		authToken:  "test_auth_token",
		fromNumber: "test_from_number",
	}

	// Attempt to send a notification.
	err := channel.SendNotification(context.Background(), "Test message", "test_to_number")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestSMSChannel_SendNotification_Error(t *testing.T) {
	// Create a mock client with a function that returns an error.
	mockClient := &MockTwilioClient{
		createMessageFunc: func(params *twilioApi.CreateMessageParams) (string, string, error) {
			return "", "", errors.New("mock error")
		},
	}

	// Create an SMSChannel instance with the mock client.
	channel := &SMSChannel{
		accountSid: "test_account_sid",
		authToken:  "test_auth_token",
		fromNumber: "test_from_number",
	}

	// Attempt to send a notification.
	err := channel.SendNotification(context.Background(), "Test message", "test_to_number")
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
}
