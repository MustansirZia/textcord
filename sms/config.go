package sms

import (
	"errors"
	"os"
	"strings"
)

// SenderConfig - To configure sms sender.
type SenderConfig interface {
	// Phone number used for sending.
	PhoneNumber() string

	// Used to identity account.
	AccountID() string

	// Used to authenticate account.
	AuthToken() string
}

type senderEnvConfig struct {
	phoneNumber string
	accountID   string
	authToken   string
}

func (e *senderEnvConfig) PhoneNumber() string {
	return e.phoneNumber
}

func (e *senderEnvConfig) AccountID() string {
	return e.accountID
}

func (e *senderEnvConfig) AuthToken() string {
	return e.authToken
}

// NewEnvConfig - Provides a new config from environment variables
// for sms sender.
func NewEnvConfig() (SenderConfig, error) {
	errorsSlice := make([]string, 0, 3)
	phoneNumber, found := os.LookupEnv("PHONE_NUMBER")
	if !found {
		errorsSlice = append(errorsSlice, "PHONE_NUMBER environment variable missing")
	}
	accountID, found := os.LookupEnv("ACCOUNT_ID")
	if !found {
		errorsSlice = append(errorsSlice, "ACCOUNT_ID environment variable missing")
	}
	authToken, found := os.LookupEnv("AUTH_TOKEN")
	if !found {
		errorsSlice = append(errorsSlice, "AUTH_TOKEN environment variable missing")
	}
	if len(errorsSlice) > 0 {
		return nil, errors.New(strings.Join(errorsSlice, "\n"))
	}
	return &senderEnvConfig{
		phoneNumber,
		accountID,
		authToken,
	}, nil
}
