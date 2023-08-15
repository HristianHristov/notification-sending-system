package channels

import (
	"context"
	"errors"
	"testing"

	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// MockTwilioClient is a mock implementation of the Twilio client interface.
type MockTwilioClient struct {
	// createMessageFunc holds a function that simulates the CreateMessage behavior.
	createMessageFuncMock func(params *twilioApi.CreateMessageParams) (*twilioApi.ApiV2010Message, error)
}

func (m *MockTwilioClient) CreateMessage(params *twilioApi.CreateMessageParams) (*twilioApi.ApiV2010Message, error) {
	// Call the mock function to simulate behavior.
	return m.createMessageFuncMock(params)
}

func TestSMSChannel_SendNotification_Success(t *testing.T) {
	// Create a mock client with a function that returns success.
	mockTwilioClient := &MockTwilioClient{createMessageFuncMock: func(params *twilioApi.CreateMessageParams) (*twilioApi.ApiV2010Message, error) {
		return &twilioApi.ApiV2010Message{}, nil
	},
	}

	// Create an SMSChannel instance with the mock client.
	channel := &SMSChannel{
		client:     mockTwilioClient,
		recepients: []string{},
	}

	// Add recepients
	channel.AddRecepients("123456789", "88888888")

	// Attempt to send a notification.
	err := channel.SendNotification(context.Background(), "Test message")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestSMSChannel_SendNotification_Error(t *testing.T) {
	// Create a mock client with a function that returns success.
	mockTwilioClient := &MockTwilioClient{createMessageFuncMock: func(params *twilioApi.CreateMessageParams) (*twilioApi.ApiV2010Message, error) {
		return &twilioApi.ApiV2010Message{}, errors.New("Error")
	},
	}

	// Create an SMSChannel instance with the mock client.
	channel := &SMSChannel{
		client:     mockTwilioClient,
		recepients: []string{},
	}

	// Add recepients
	channel.AddRecepients("123456789", "88888888")

	// Attempt to send a notification.
	err := channel.SendNotification(context.Background(), "Test message")
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
}
