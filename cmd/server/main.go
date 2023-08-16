package main

import (
	"context"
	"log"
	"net"
	notifications "notification/api" // Import the generated package
	"notification/internal/channels"
	"notification/internal/service"
	"os"

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
	// To improve this implementation we could implement a factory
	for _, config := range req.ChannelConfigs {
		switch config.Type {
		case "email":
			emailChannel := channels.NewEmailChannel(os.Getenv("SMTP_SERVER"), 25, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))
			emailChannel.AddRecepients(config.Recipients...)
			notificationService.AddNotificationChannels(emailChannel)
		case "sms":
			smsChannel := channels.NewSMSChannel(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"))
			smsChannel.AddRecepients(config.Recipients...)
			notificationService.AddNotificationChannels(smsChannel)
		case "slack":
			slackChannel := channels.NewSlackChannel(os.Getenv("SLACK_API_TOKEN"))
			slackChannel.AddRecepients(config.Recipients...)
			notificationService.AddNotificationChannels(slackChannel)
		default:
			log.Println("No such channel type")
		}
	}

	// Send notifications and wait for goroutines to finish
	errChan := notificationService.SendNotifications(context.Background(), req.Message.Text)

	// Iterate over the error channel and send responses to the client
	for err := range errChan {
		errResp := &notifications.NotificationResponse{
			Success: false,
			Error:   err.Error(),
		}
		if sendErr := stream.Send(errResp); sendErr != nil {
			log.Printf("Error sending response to client: %v", sendErr)
		}
	}

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
