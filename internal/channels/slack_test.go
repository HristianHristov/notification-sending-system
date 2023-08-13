package channels

import (
	"context"
	"errors"
	"testing"

	"github.com/slack-go/slack"
)

// SlackClient is an interface that represents the methods of slack.Client needed by SlackChannel.
type SlackClient interface {
	PostMessageContext(ctx context.Context, channelID string, options ...slack.MsgOption) (string, string, error)
}

// MockSlackClient is a mock implementation of the SlackClient interface.
type MockSlackClient struct {
	postMessageFunc func(channelID string, options ...slack.MsgOption) (string, string, error)
}

func (m *MockSlackClient) PostMessageContext(ctx context.Context, channelID string, options ...slack.MsgOption) (string, string, error) {
	return m.postMessageFunc(channelID, options...)
}

func TestSlackChannel_SendNotification_Success(t *testing.T) {
	mockClient := &MockSlackClient{
		postMessageFunc: func(channelID string, options ...slack.MsgOption) (string, string, error) {
			return "", "", nil
		},
	}

	// Create a SlackChannel instance with the mock client
	channel := &SlackChannel{
		client: mockClient,
	}

	err := channel.SendNotification(context.Background(), "Test message", "channel_id_1")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestSlackChannel_SendNotification_Error(t *testing.T) {
	mockClient := &MockSlackClient{
		postMessageFunc: func(channelID string, options ...slack.MsgOption) (string, string, error) {
			return "", "", errors.New("mock error")
		},
	}

	// Create a SlackChannel instance with the mock client
	channel := &SlackChannel{
		client: mockClient,
	}

	err := channel.SendNotification(context.Background(), "Test message", "channel_id_1")
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
