package channels

import (
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// SMSChannel represents an SMS notification channel using the Twilio API.
type SMSChannel struct {
	accountSid string
	authToken  string
	fromNumber string
}

// NewSMSChannel creates a new instance of the SMSChannel.
func NewSMSChannel(accountSid, authToken, fromNumber string) *SMSChannel {
	return &SMSChannel{
		accountSid: accountSid,
		authToken:  authToken,
		fromNumber: fromNumber,
	}
}

// SendNotification sends an SMS notification to the specified phone number.
func (s *SMSChannel) SendNotification(message, toNumber string) error {

	// Initialize Twilio client
	client := twilio.NewRestClient()

	// Send SMS using Twilio API
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(toNumber)
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)

	return err
}
