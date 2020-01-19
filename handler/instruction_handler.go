package handler

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mustansirzia/simcord/discord"
	"github.com/mustansirzia/simcord/parser"
)

var discordClient discord.Client

func init() {
	discordConfig, err := discord.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	discordClient, err = discord.NewClient(discordConfig)
	if err != nil {
		log.Fatal(err)
	}
}

// HandleInstruction - Function which does the actual handling of the instruction
// and resolves an instruction to a string response which can be sent to the client.
func HandleInstruction(instruction *parser.Instruction) (string, error) {
	if instruction.Type == parser.GET {
		return handleGetInstruction(instruction)
	}
	if instruction.Type == parser.SEND {
		return "", handleSendInstruction(instruction)
	}
	return "", nil
}

func handleGetInstruction(instruction *parser.Instruction) (string, error) {
	switch instruction.Name {
	case parser.GUILDS:
		if len(instruction.ResourceID) > 0 {
			channels, err := discordClient.Channels(instruction.ResourceID)
			if err != nil {
				return "", err
			}
			return convertResourceSliceToString(channels), nil
		}
		guilds, err := discordClient.Guilds()
		if err != nil {
			return "", err
		}
		return convertResourceSliceToString(guilds), nil
	case parser.CHANNELS:
		count := 10
		if len(instruction.Arg) > 0 {
			countInt, err := strconv.Atoi(instruction.Arg)
			if err == nil {
				count = countInt
			}
		}
		messages, err := discordClient.Messages(instruction.ResourceID, count)
		if err != nil {
			return "", err
		}
		return convertResourceSliceToString(messages), nil
	}
	return "", nil
}

func handleSendInstruction(instruction *parser.Instruction) error {
	switch instruction.Name {
	case parser.CHANNELS:
		err := discordClient.SendMessage(instruction.ResourceID, instruction.Arg)
		if err != nil {
			return err
		}
	}
	return nil
}

func convertResourceSliceToString(resources interface{}) string {
	var resourcesAsString []string
	switch resourcesSlice := resources.(type) {
	case []discord.Guild:
		resourcesAsString = make([]string, 0, len(resourcesSlice))
		for _, guild := range resourcesSlice {
			resourcesAsString = append(resourcesAsString, fmt.Sprintf("%s", guild))
		}
	case []discord.Channel:
		resourcesAsString = make([]string, 0, len(resourcesSlice))
		for _, channel := range resourcesSlice {
			resourcesAsString = append(resourcesAsString, fmt.Sprintf("%s", channel))
		}
	case []discord.Message:
		resourcesAsString = make([]string, 0, len(resourcesSlice))
		for _, message := range resourcesSlice {
			resourcesAsString = append(resourcesAsString, fmt.Sprintf("%s", message))
		}
	default:
		log.Fatal("Unknown argument given to convertResourceSliceToString")
	}
	return strings.Join(resourcesAsString, "\n")

}
