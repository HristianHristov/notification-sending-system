package channels

import "context"

type NotificationChannel interface {
	SendNotification(ctx context.Context, message string, params []string) error
	GetType() string
}
