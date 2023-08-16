package service

import (
	"context"
	"errors"
	"testing"
)

// MockNotificationChannel is a mock implementation of the NotificationChannel interface.
type MockNotificationChannel struct {
	getTypeFunc          func() string
	addRecipientsFunc    func(...string)
	sendNotificationFunc func(context.Context, string) error
}

func (m *MockNotificationChannel) GetType() string {
	return m.getTypeFunc()
}

func (m *MockNotificationChannel) AddRecepients(recipients ...string) {
	m.addRecipientsFunc(recipients...)
}

func (m *MockNotificationChannel) SendNotification(ctx context.Context, message string) error {
	return m.sendNotificationFunc(ctx, message)
}

func TestSendNotifications(t *testing.T) {
	// Create mock implementations of NotificationChannels for testing
	mockEmailChannel := &MockNotificationChannel{
		getTypeFunc: func() string { return "email" },
		sendNotificationFunc: func(ctx context.Context, message string) error {
			return nil
		},
	}

	mockSMSChannel := &MockNotificationChannel{
		getTypeFunc: func() string { return "sms" },
		sendNotificationFunc: func(ctx context.Context, message string) error {
			return errors.New("failed to send")
		},
	}

	mockSlackChannel := &MockNotificationChannel{
		getTypeFunc: func() string { return "slack" },
		sendNotificationFunc: func(ctx context.Context, message string) error {
			return nil
		},
	}

	// Create a NotificationService with the mock channels
	notificationService := NewNotificationService(mockEmailChannel, mockSMSChannel, mockSlackChannel)

	// Call the function under test and get the error channel
	errChan := notificationService.SendNotifications(context.Background(), "Test Message")

	// Collect errors from the errChan
	var errorsReceived []error
	for err := range errChan {
		errorsReceived = append(errorsReceived, err)
	}

	// Assert the results
	if len(errorsReceived) != 1 || errorsReceived[0].Error() != "failed after 3 retries" {
		t.Errorf("Expected 1 error with message 'failed after 3 retries', but got: %v", errorsReceived)
	}
}
