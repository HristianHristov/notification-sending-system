package channels

import (
	"context"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSClient interface {
	CreateMessage(params *twilioApi.CreateMessageParams) (*twilioApi.ApiV2010Message, error)
}

// SMSChannel represents an SMS notification channel using the Twilio API.
type SMSChannel struct {
	client     SMSClient
	recepients []string
}

// NewSMSChannel creates a new instance of the SMSChannel.
func NewSMSChannel(accountSid, authToken string) *SMSChannel {
	client := twilio.NewRestClient().Api
	return &SMSChannel{
		client:     client,
		recepients: []string{},
	}
}

func (s *SMSChannel) AddRecepients(recepients ...string) {
	s.recepients = append(s.recepients, recepients...)
}

// SendNotification sends an SMS notification to the specified phone number.
func (s *SMSChannel) SendNotification(ctx context.Context, message string) error {

	// Send SMS using Twilio API
	params := &twilioApi.CreateMessageParams{}
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody(message)

	var err error
	for _, recepient := range s.recepients {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
			params.SetTo(recepient)
			_, err = s.client.CreateMessage(params)
		}
	}
	return err
}

func (s *SMSChannel) GetType() string {
	return "sms"
}
