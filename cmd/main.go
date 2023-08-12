package main

import (
	"notification/internal/channels"
	"notification/internal/service"
)

func main() {
	// Create instances of notification channels
	emailChannel := channels.NewEmailChannel("smtp.example.com", 587, "your@email.com", "your_password")
	smsChannel := channels.NewSMSChannel("your_twilio_account_sid", "your_twilio_auth_token", "your_twilio_phone_number")
	slackChannel := channels.NewSlackChannel("your_slack_api_token")

	// Create the notification service and register channels
	notificationService := service.NewNotificationService(emailChannel, smsChannel, slackChannel)

	// Send a notification
	message := "Hello, this is a notification!"
	toEmail := "recipient@example.com"

	notificationService.SendNotification(message, toEmail)
}
