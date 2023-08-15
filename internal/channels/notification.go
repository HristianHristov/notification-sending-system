package channels

import "context"

type NotificationChannel interface {
	SendNotification(ctx context.Context, message string) error
	GetType() string
}
