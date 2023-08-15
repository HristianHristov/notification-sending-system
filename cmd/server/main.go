package main

import (
	"context"
	"log"
	"net"
	"notification/internal/channels"

	notifications "notification/api" // Import the generated package
	"notification/internal/service"

	"google.golang.org/grpc"
)

type NotificationServer struct {
	notifications.UnimplementedNotificationServiceServer
}

func NewNotificationServer() *NotificationServer {
	return &NotificationServer{}
}

func (s *NotificationServer) SendNotifications(req *notifications.NotificationRequest, stream notifications.NotificationService_SendNotificationsServer) error {
	// Create a new notification service
	notificationService := service.NewNotificationService()

	// Create channels and add recipients
	targetChannels := make([]string, 0)
	for _, config := range req.ChannelConfigs {
		switch config.Type {
		case "email":
			emailChannel := channels.NewEmailChannel("smtp.example.com", 587, "smtp_username", "smtp_password")
			emailChannel.AddRecepients(config.Recipients...)
			notificationService.AddNotificationChannels(emailChannel)
			targetChannels = append(targetChannels, "email")
		case "sms":
			smsChannel := channels.NewSMSChannel("twilio_account_sid", "twilio_auth_token")
			smsChannel.AddRecepients(config.Recipients...)
			notificationService.AddNotificationChannels(smsChannel)
			targetChannels = append(targetChannels, "sms")
		case "slack":
			slackChannel := channels.NewSlackChannel("slack_api_token")
			slackChannel.AddRecepients(config.Recipients...)
			notificationService.AddNotificationChannels(slackChannel)
			targetChannels = append(targetChannels, "slack")
		default:
			log.Println("No such channel type")
		}
	}

	// Send notifications and wait for goroutines to finish
	err := notificationService.SendNotifications(context.Background(), req.Message.Text, targetChannels...)
	if err != nil {
		log.Printf("Error sending notifications: %v", err)
	}

	_ = stream.Send(&notifications.NotificationResponse{Success: true})

	return nil
}
func main() {
	// Create and start the gRPC server
	server := grpc.NewServer()
	notificationServer := NewNotificationServer()
	notifications.RegisterNotificationServiceServer(server, notificationServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Server listening on port 50051")
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
