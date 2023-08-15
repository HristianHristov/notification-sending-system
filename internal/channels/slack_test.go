package channels

import (
	"context"
	"errors"
	"testing"

	"github.com/slack-go/slack"
)

// MockSlackClient is a mock implementation of the SlackClient interface.
type MockSlackClient struct {
	postMessageFuncMock func(ctx context.Context, channelID string, options ...slack.MsgOption) (string, string, error)
}

func (m *MockSlackClient) PostMessageContext(ctx context.Context, channelID string, options ...slack.MsgOption) (string, string, error) {
	return m.postMessageFuncMock(ctx, channelID, options...)
}

func TestSlackChannel_SendNotification_Success(t *testing.T) {
	mockClient := &MockSlackClient{
		postMessageFuncMock: func(ctx context.Context, channelID string, options ...slack.MsgOption) (string, string, error) {
			return "", "", nil
		},
	}

	// Create a SlackChannel instance with the mock client
	channel := &SlackChannel{
		client:     mockClient,
		recepients: []string{},
	}

	channel.AddRecepients("test1", "test2")

	err := channel.SendNotification(context.Background(), "Test message")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestSlackChannel_SendNotification_Error(t *testing.T) {
	mockClient := &MockSlackClient{
		postMessageFuncMock: func(ctx context.Context, channelID string, options ...slack.MsgOption) (string, string, error) {
			return "", "", errors.New("Error")
		},
	}

	// Create a SlackChannel instance with the mock client
	channel := &SlackChannel{
		client:     mockClient,
		recepients: []string{},
	}

	channel.AddRecepients("test1", "test2")

	err := channel.SendNotification(context.Background(), "Test message")
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
}

func TestSlackChannel_AddRecepients(t *testing.T) {
	channel := &SlackChannel{}
	channel.AddRecepients("channel_1", "channel_2")

	if len(channel.recepients) != 2 {
		t.Errorf("Expected 2 recipients, but got %d", len(channel.recepients))
	}
}
