package channels

import (
	"fmt"

	"github.com/slack-go/slack"
)

// SlackChannel represents a Slack notification channel.
type SlackChannel struct {
	apiToken string
}

// NewSlackChannel creates a new instance of the SlackChannel.
func NewSlackChannel(apiToken string) *SlackChannel {
	return &SlackChannel{
		apiToken: apiToken,
	}
}

// SendNotification sends a message to a Slack channel with the provided name, if exists
func (s *SlackChannel) SendNotification(message, channelName string) error {
	api := slack.New(s.apiToken)

	channelID, err := s.getChannelID(api, channelName)
	//TODO: Improve error handling
	if err != nil {
		return fmt.Errorf("failed to get channel ID: %w", err)
	}

	_, _, err = api.PostMessage(channelID,
		slack.MsgOptionText(message, false),
	)
	if err != nil {
		return fmt.Errorf("failed to send Slack message: %w", err)
	}

	return nil
}

// getChannelID fetches the ID of a channel based on the name
// TODO: Provide a better implementation, as this approach is not efficient
func (s *SlackChannel) getChannelID(api *slack.Client, channelName string) (string, error) {
	conversations, _, err := api.GetConversations(&slack.GetConversationsParameters{
		ExcludeArchived: true,
	})
	if err != nil {
		return "", err
	}

	for _, convo := range conversations {
		if convo.Name == channelName {
			return convo.ID, nil
		}
	}
	//TODO: Improve error handling
	return "", fmt.Errorf("channel not found: %s", channelName)
}
