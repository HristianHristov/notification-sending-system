package channels

import (
	"context"
	"fmt"
	"net/smtp"
)

// EmailChannel represents an email notification channel.
type EmailChannel struct {
	smtpServer   string
	smtpPort     int
	smtpUsername string
	smtpPassword string
	recepients   []string
	smtpSendFunc func(string, smtp.Auth, string, []string, []byte) error
}

// NewEmailChannel creates a new instance of the EmailChannel.
func NewEmailChannel(smtpServer string, smtpPort int, smtpUsername, smtpPassword string) *EmailChannel {
	return &EmailChannel{
		smtpServer:   smtpServer,
		smtpPort:     smtpPort,
		smtpUsername: smtpUsername,
		smtpPassword: smtpPassword,
		smtpSendFunc: smtp.SendMail,
	}
}

func (e *EmailChannel) AddRecepients(recepients ...string) {
	e.recepients = append(e.recepients, recepients...)
}

// SendNotification sends a notification email to the specified email address.
func (e *EmailChannel) SendNotification(ctx context.Context, message string) error {
	// Authenticate with the SMTP server using the provided credentials
	auth := smtp.PlainAuth("", e.smtpUsername, e.smtpPassword, e.smtpServer)

	var err error
	for _, recepient := range e.recepients {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
			// Construct the email message
			msg := fmt.Sprintf("To: %s\r\nSubject: Notification\r\n\r\n%s", recepient, message)
			err = e.smtpSendFunc(fmt.Sprintf("%s:%d", e.smtpServer, e.smtpPort), auth, e.smtpUsername, []string{recepient}, []byte(msg))
		}
	}

	// Send the email using the SMTP server
	return err
}

func (e *EmailChannel) GetType() string {
	return "email"
}
