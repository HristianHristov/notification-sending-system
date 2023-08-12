package channels

import (
	"fmt"
	"net/smtp"
)

// EmailChannel represents an email notification channel.
type EmailChannel struct {
	smtpServer   string
	smtpPort     int
	smtpUsername string
	smtpPassword string
}

// NewEmailChannel creates a new instance of the EmailChannel.
func NewEmailChannel(smtpServer string, smtpPort int, smtpUsername, smtpPassword string) *EmailChannel {
	return &EmailChannel{
		smtpServer:   smtpServer,
		smtpPort:     smtpPort,
		smtpUsername: smtpUsername,
		smtpPassword: smtpPassword,
	}
}

// SendNotification sends a notification email to the specified email address.
func (e *EmailChannel) SendNotification(message string, toEmail string) error {
	// Authenticate with the SMTP server using the provided credentials
	auth := smtp.PlainAuth("", e.smtpUsername, e.smtpPassword, e.smtpServer)

	// Construct the email message
	msg := fmt.Sprintf("To: %s\r\nSubject: Notification\r\n\r\n%s", toEmail, message)

	// Send the email using the SMTP server
	err := smtp.SendMail(fmt.Sprintf("%s:%d", e.smtpServer, e.smtpPort), auth, e.smtpUsername, []string{toEmail}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}