package service

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"notification/internal/channels"
)

type NotificationService struct {
	channels []channels.NotificationChannel
}

func NewNotificationService(channels ...channels.NotificationChannel) *NotificationService {
	return &NotificationService{channels: channels}
}

func (n *NotificationService) AddNotificationChannels(channels ...channels.NotificationChannel) {
	n.channels = append(n.channels, channels...)
}

func (n *NotificationService) SendNotifications(ctx context.Context, message string, targetChannels ...string) <-chan (error) {
	var wg sync.WaitGroup // Create a WaitGroup to wait for all goroutines
	errChan := make(chan error, len(targetChannels))

	for _, channel := range n.channels {
		wg.Add(1) // Increment the WaitGroup counter

		go func(ctx context.Context, ch channels.NotificationChannel) {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine finishes

			// Use context for timeouts, cancellations, etc.
			err := n.sendWithRetry(ctx, ch, message, 3) // Retry 3 times
			if err != nil {
				log.Printf("Error sending through %s channel: %s\n", ch.GetType(), err)
				errChan <- err
				log.Printf("After Adding to chan")
			}
		}(ctx, channel)
	}

	go func() {
		defer close(errChan)
		wg.Wait() // Wait for all goroutines to finish before returning
	}()

	return errChan
}

func (n *NotificationService) sendWithRetry(ctx context.Context, ch channels.NotificationChannel, message string, maxRetries int) error {
	for retries := 0; retries < maxRetries; retries++ {
		err := ch.SendNotification(ctx, message) // Pass ctx and any required parameters
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
