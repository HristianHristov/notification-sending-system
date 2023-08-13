package channels

import (
	"context"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// SMSChannel represents an SMS notification channel using the Twilio API.
type SMSChannel struct {
	client *twilio.RestClient
}

// NewSMSChannel creates a new instance of the SMSChannel.
func NewSMSChannel(accountSid, authToken, fromNumber string) *SMSChannel {
	client := twilio.NewRestClient()
	return &SMSChannel{
		client: client,
	}
}

// SendNotification sends an SMS notification to the specified phone number.
func (s *SMSChannel) SendNotification(ctx context.Context, message string, recepients []string) error {

	// Send SMS using Twilio API
	params := &twilioApi.CreateMessageParams{}
	//params.SetTo(recepient)
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody(message)

	var err error
	for _, recepient := range recepients {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
			p := *params
			p.SetTo(recepient)
			_, err = s.client.Api.CreateMessage(&p)
		}

	}
	return err
}

func (s *SMSChannel) GetType() string {
	return "sms"
}
