package discord

import (
	"errors"
	"os"
	"strings"
)

// ClientConfig - To configure discord client.
type ClientConfig interface {
	// Token - Used to authenticate.
	Token() string

	// IsBot - Used to identify if token belongs to bot.
	IsBot() bool
}

type clientEnvConfig struct {
	token string
}

func (e *clientEnvConfig) Token() string {
	return e.token
}

func (e *clientEnvConfig) IsBot() bool {
	return strings.Contains(e.token, "Bot ")
}

// NewEnvConfig - Provides a new config from environment variables
// for discord client.
func NewEnvConfig() (ClientConfig, error) {
	discordToken, found := os.LookupEnv("DISCORD_TOKEN")
	if !found {
		return nil, errors.New("DISCORD_TOKEN environment variable missing")
	}

	return &clientEnvConfig{
		discordToken,
	}, nil
}
