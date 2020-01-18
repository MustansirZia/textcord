package sms

import (
	"github.com/sfreiberg/gotwilio"
)

// Sender - Implementations of this
// will be used to send SMSes. (If required)
type Sender interface {
	// Send - To send a new SMS.
	Send(phoneNumber, message string) error
}

type twilioSender struct {
	twilio      *gotwilio.Twilio
	phoneNumber string
}

func (s *twilioSender) Send(phoneNumber, message string) error {
	_, exception, err := s.twilio.SendSMS(
		s.phoneNumber,
		phoneNumber,
		message,
		"",
		"",
	)
	if exception != nil {
		return exception
	}
	if err != nil {
		return err
	}
	return nil
}

// NewSender - Gets a new SMS sender.
func NewSender(config SenderConfig) Sender {
	twilio := gotwilio.NewTwilioClient(
		config.AccountID(),
		config.AuthToken(),
	)
	return &twilioSender{
		twilio:      twilio,
		phoneNumber: config.PhoneNumber(),
	}
}
