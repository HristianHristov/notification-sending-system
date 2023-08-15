package channels

import (
	"context"
	"net/smtp"
	"testing"
	"time"
)

func TestEmailChannel_SendNotification(t *testing.T) {
	// Create an instance of the EmailChannel with a mock SMTP sending function
	mockSMTPSend := func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		return nil
	}
	emailChannel := NewEmailChannel("smtp.example.com", 587, "sender@example.com", "password")
	emailChannel.AddRecepients("recipient@example.com")
	emailChannel.smtpSendFunc = mockSMTPSend

	// Define the context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Test successful email sending
	err := emailChannel.SendNotification(ctx, "Test message")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Test context cancellation
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel() // Immediately cancel the context

	err = emailChannel.SendNotification(cancelledCtx, "Cancelled message")
	if err == nil || err != context.Canceled {
		t.Errorf("Expected context cancellation error, but got: %v", err)
	}
}

func TestEmailChannel_GetType(t *testing.T) {
	// Create an instance of the EmailChannel
	emailChannel := NewEmailChannel("smtp.example.com", 587, "sender@example.com", "password")

	// Get the channel type
	typ := emailChannel.GetType()

	// Check if the returned channel type matches the expected value
	if typ != "email" {
		t.Errorf("Expected channel type 'email', but got '%s'", typ)
	}
}
