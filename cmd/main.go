package main

import (
	"context"
	"log"
	"notification/internal/channels"
	"notification/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create instances of notification channels
	emailChannel := channels.NewEmailChannel("smtp.example.com", 587, "your@email.com", "your_password")
	emailChannel.AddRecepients("recipient1@example.com", "recipient2@example.com")

	smsChannel := channels.NewSMSChannel("your_twilio_account_sid", "your_twilio_auth_token")
	smsChannel.AddRecepients("+1234567890")

	slackChannel := channels.NewSlackChannel("your_slack_api_token")
	slackChannel.AddRecepients("a1")

	ctx := context.Background()

	// Create the notification service and register channels
	notificationService := service.NewNotificationService(emailChannel, smsChannel, slackChannel)

	// Initialize the Gin router
	r := gin.Default()

	// Define a POST endpoint to send notifications
	r.POST("/send-notification", func(c *gin.Context) {
		type PostNotificationDTO struct {
			Message        string
			TargetChannels []string
		}
		var dto PostNotificationDTO
		err := c.ShouldBindJSON(&dto)

		if err != nil {
			log.Printf("Error sending notification: %v", err)
			c.JSON(500, gin.H{"error": "Failed to send notification"})
			return
		}

		notificationService.SendNotification(ctx, dto.Message, dto.TargetChannels...)

		c.JSON(200, gin.H{"message": "Notification is scheduled!"})
	})

	// Run the server
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
