package discord

import (
	"fmt"
	"sort"
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
	Channels(guildID string) ([]Channel, error)

	// Messages - Returns the last "count" of messages for this
	// channel.
	Messages(channelID string, count int) ([]Message, error)

	// SendMessage - To send a message to a channel.
	SendMessage(channelID string, message string) error
}

type resource struct {
	// Snowflake ID of the resource.
	ID string

	// Name of the resource.
	Name string
}

func (e resource) String() string {
	return fmt.Sprintf("%s: %s", e.Name, e.ID)
}

// Guild - A server.
type Guild struct {
	resource
}

// Channel - A text channel inside a server.
type Channel struct {
	resource
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

func (m Message) String() string {
	return fmt.Sprintf("%s at %s: %s", m.Sender, m.SentAt.Format("02 Jan 06 15:04"), m.Text)
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
			resource: resource{
				ID:   userGuild.ID,
				Name: userGuild.Name,
			}})
	}
	return guilds, nil

}

func (c *client) Channels(guildID string) ([]Channel, error) {
	userChannels, err := c.session.GuildChannels(guildID)
	if err != nil {
		return nil, err
	}
	channels := make([]Channel, 0, len(userChannels))
	for _, userChannel := range userChannels {
		if userChannel.Type == discordgo.ChannelTypeGuildText {
			// Filter out non text channels.
			channels = append(channels, Channel{
				resource: resource{
					ID:   userChannel.ID,
					Name: userChannel.Name,
				}})
		}
	}
	return channels, nil
}

func (c *client) Messages(channelID string, count int) ([]Message, error) {
	userMessages, err := c.session.ChannelMessages(channelID, count, "", "", "")
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
	// Sorting so that latest messages appear at the end of the slice.
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].SentAt.Before(messages[j].SentAt)
	})
	return messages, nil
}

func (c *client) SendMessage(channelID string, message string) error {
	_, err := c.session.ChannelMessageSend(channelID, message)
	if err != nil {
		return err
	}
	return nil
}

// NewClient - Gets a new Discord client.
func NewClient(config ClientConfig) (Client, error) {
	s, err := discordgo.New(config.Token())
	if err != nil {
		return nil, err
	}
	return &client{
		s,
	}, nil
}
