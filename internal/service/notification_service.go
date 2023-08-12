package service

import "fmt"

type NotificationChannel interface {
	SendNotification(message, param string) error
}

type NotificationService struct {
	channels []NotificationChannel
}

func NewNotificationService(channels ...NotificationChannel) *NotificationService {
	return &NotificationService{channels: channels}
}

func (n *NotificationService) SendNotification(message, param string) {
	for _, channel := range n.channels {
		go func(ch NotificationChannel) {
			err := ch.SendNotification(message, param)
			if err != nil {
				// Handle error, possibly retry
				fmt.Println(err)
			}
		}(channel)
	}
}
