package main

import (
	"context"
	"fmt"
	"io"
	"log"
	notifications "notification/api" // Import the generated package

	"google.golang.org/grpc"
)

func main() {
	serverAddr := "localhost:50051" // Change to your server's address

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := notifications.NewNotificationServiceClient(conn)

	// Define your notification message and channel configs here
	message := &notifications.Message{Text: "Hello, world!"}
	channelConfigs := []*notifications.ChannelConfig{
		{
			Type:       "email",
			Recipients: []string{"recipient1@example.com", "recipient2@example.com"},
		},
		{
			Type:       "sms",
			Recipients: []string{"+1234567890", "+9876543210"},
		},
		//Add more channel configs as needed
	}

	req := &notifications.NotificationRequest{
		Message:        message,
		ChannelConfigs: channelConfigs,
	}

	stream, err := client.SendNotifications(context.Background(), req)
	if err != nil {
		log.Fatalf("Error sending notifications: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Error receiving response: %v", err)
			break
		}
		fmt.Printf("Notification sent: %v\n", resp.Success)
	}
}
