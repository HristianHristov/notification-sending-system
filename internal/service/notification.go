package service

import (
	"context"
	"fmt"
	"log"
	"notification/internal/channels"
	"time"
)

type NotificationService struct {
	channels []channels.NotificationChannel
}

func NewNotificationService(channels ...channels.NotificationChannel) *NotificationService {
	return &NotificationService{channels: channels}
}

func (n *NotificationService) SendNotification(ctx context.Context, message string, params []string, targetChannels ...string) {
	for _, channel := range n.channels {
		for _, target := range targetChannels {
			if channel.GetType() == target {
				go func(ctx context.Context, ch channels.NotificationChannel) {
					// Use context for timeouts, cancellations, etc.
					err := n.sendWithRetry(ctx, ch, message, params, 3) // Retry 3 times
					if err != nil {
						fmt.Printf("Error sending through %s channel: %s\n", ch.GetType(), err)
					}
				}(ctx, channel)
				//break // Send through only one channel of each type
			} else {
				log.Println("No such channel!")
			}
		}
	}
}

func (n *NotificationService) sendWithRetry(ctx context.Context, ch channels.NotificationChannel, message string, params []string, maxRetries int) error {
	for retries := 0; retries < maxRetries; retries++ {
		err := ch.SendNotification(ctx, message, params) // Pass ctx and any required parameters
		if err == nil {
			return nil // Sent successfully
		}

		// If context canceled or deadline exceeded, stop retrying
		if ctx.Err() != nil {
			return ctx.Err()
		}

		// Wait before retrying
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Duration(retries+1) * time.Second):
			continue
		}
	}

	return fmt.Errorf("failed after %d retries", maxRetries)
}
