package channels

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"
)

type SlackClient interface {
	PostMessageContext(ctx context.Context, channelID string, options ...slack.MsgOption) (string, string, error)
}

// SlackChannel represents a Slack notification channel.
type SlackChannel struct {
	client     SlackClient
	recepients []string
}

// NewSlackChannel creates a new instance of the SlackChannel.
func NewSlackChannel(apiToken string) *SlackChannel {
	client := slack.New(apiToken)
	return &SlackChannel{
		client:     client,
		recepients: []string{},
	}
}

// AddRecepients adds a list of recepients to the channel
func (s *SlackChannel) AddRecepients(channels ...string) {
	s.recepients = append(s.recepients, channels...)
}

// SendNotification sends a message to a Slack channel with the provided name, if exists
func (s *SlackChannel) SendNotification(ctx context.Context, message string) error {

	var err error
	for _, recepient := range s.recepients {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			_, _, err = s.client.PostMessageContext(ctx, recepient,
				slack.MsgOptionText(message, false),
			)
			if err != nil {
				return fmt.Errorf("failed to send Slack message: %w", err)
			}
		}
	}

	return nil
}

func (s *SlackChannel) GetType() string {
	return "slack"
}
