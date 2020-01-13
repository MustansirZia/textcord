package discord

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// Client - Custom wrapper to interact with discord API.
type Client interface {
	// Guilds - Returns the guilds or servers the discord user or bot
	// is added to.
	Guilds() ([]Guild, error)

	// Channels - Returns the channels inside a a particular
	// guild or server this user or bot is added to.
	Channels(guild Guild) ([]Channel, error)

	// Messages - Returns the last "count" of messages for this
	// channel.
	Messages(channel Channel, count int) ([]Message, error)

	// SendMessage - To send a message to a channel.
	SendMessage(channel Channel, message string) error
}

type entity struct {
	// Snowflake ID of the entity.
	ID string

	// Name of the entity.
	Name string
}

// Guild - A server.
type Guild struct {
	entity
}

// Channel - A text channel inside a server.
type Channel struct {
	entity
}

// Message - Represents a text message inside a channel.
type Message struct {
	// Text - Content of the message.
	Text string

	// Sender - Sender of the message.
	Sender string

	// SentAt - Instant this message was sent at.
	SentAt time.Time
}

type client struct {
	session *discordgo.Session
}

func (c *client) Guilds() ([]Guild, error) {
	userGuilds, err := c.session.UserGuilds(20, "", "")
	if err != nil {
		return nil, err
	}
	guilds := make([]Guild, 0, len(userGuilds))
	for _, userGuild := range userGuilds {
		guilds = append(guilds, Guild{
			entity: entity{
				ID:   userGuild.ID,
				Name: userGuild.Name,
			}})
	}
	return guilds, nil

}

func (c *client) Channels(guild Guild) ([]Channel, error) {
	userGuild, err := c.session.Guild(guild.ID)
	if err != nil {
		return nil, err
	}
	channels := make([]Channel, 0, len(userGuild.Channels))
	for _, userChannel := range userGuild.Channels {
		channels = append(channels, Channel{
			entity: entity{
				ID:   userChannel.ID,
				Name: userChannel.Name,
			}})
	}
	return channels, nil
}

func (c *client) Messages(channel Channel, count int) ([]Message, error) {
	userMessages, err := c.session.ChannelMessages(channel.ID, count, "", "", "")
	if err != nil {
		return nil, err
	}
	messages := make([]Message, 0, count)
	for _, userMessage := range userMessages {
		sentAt, _ := userMessage.Timestamp.Parse()
		messages = append(messages, Message{
			Text:   userMessage.ContentWithMentionsReplaced(),
			Sender: userMessage.Author.Username,
			SentAt: sentAt,
		})
	}
	return messages, nil
}

func (c *client) SendMessage(channel Channel, message string) error {
	_, err := c.session.ChannelMessageSend(channel.ID, message)
	if err != nil {
		return err
	}
	return nil
}

// NewClient - Gets a new Discord client.
func NewClient() (Client, error) {
	s, err := discordgo.New("Bot NjY1OTkyMzI0NzUyMzQzMDQw.XhySCA.iISga9ksgoAZwDzhcZWYAaDR5lM")
	if err != nil {
		return nil, err
	}
	return &client{
		s,
	}, nil
}
