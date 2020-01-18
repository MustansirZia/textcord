package parser

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
)

// InstructionParser - Which unmarshalls an instruction
// from a *http.Request.
type InstructionParser interface {
	// Parse - Actual method which parses the request.
	Parse(r *http.Request) (*Instruction, error)
	// Pattern - Method which returns the format of the input.
	Pattern() string
}

// Instruction - An unmarshalled instruction.
type Instruction struct {
	Type instructionMethod

	Name instructionResource

	// ResourceID - If present, denotes a single resource.
	ResourceID string

	Arg string
}

// ErrInvalidRequest - This error will be returned from Parse if the request is malformed.
var ErrInvalidRequest = errors.New("Invalid Request")

type instructionMethod string

const (
	// GET - For getting a resource.
	GET instructionMethod = "GET"
	// SEND - For sending a message.
	SEND instructionMethod = "SEND"
)

type instructionResource string

const (
	// GUILDS - This signfies the instruction pertains to Guilds.
	GUILDS instructionResource = "GUILDS"
	// CHANNELS - This signfies the instruction pertains to Channels.
	CHANNELS instructionResource = "CHANNELS"
)

type twilioSMSParser struct{}

var twilioSMSParserRegex *regexp.Regexp

func init() {
	twilioSMSParserRegex = regexp.MustCompile("\\s+")
}

func (twilioSMSParser) Parse(r *http.Request) (*Instruction, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	body := r.FormValue("Body")
	if len(body) == 0 {
		return nil, ErrInvalidRequest
	}
	// Lets split the body using spaces between words.
	// One or more consecutive whitespaces are considered as one will spliting.
	// So "  word1   word2 word3   word4  " will be split
	// into []string{"word1", "word2", "word3", "wor4"}.
	// The maximum length of the slice will be 4.
	split := twilioSMSParserRegex.Split(strings.Trim(body, " "), 4)
	if len(split) < 2 {
		return nil, ErrInvalidRequest
	}

	instruction := new(Instruction)
	// Parsing the instruction type.
	switch strings.ToUpper(split[0]) {
	case string(GET):
		instruction.Type = GET
	case string(SEND):
		instruction.Type = SEND
	default:
		return nil, ErrInvalidRequest
	}

	// Parsing the instruction name.
	switch strings.ToUpper(split[1]) {
	case string(GUILDS):
		instruction.Name = GUILDS
	case string(CHANNELS):
		instruction.Name = CHANNELS
	default:
		return nil, ErrInvalidRequest
	}

	// Parsing the resource ID, if present.
	if len(split) > 2 {
		instruction.ResourceID = split[2]
	}

	// Adding Args, if any.
	if len(split) > 3 {
		instruction.Arg = split[3]
	}

	// Resource ID must be supplied in case of channels.
	if instruction.Name == CHANNELS && len(instruction.ResourceID) == 0 {
		return nil, ErrInvalidRequest
	}

	// Extra argument must be supplied in case of SEND.
	if instruction.Type == SEND && len(instruction.Arg) == 0 {
		return nil, ErrInvalidRequest
	}

	return instruction, nil
}

func (twilioSMSParser) Pattern() string {
	return "GET|SEND GUILDS|CHANNELS [GUILD_ID|CHANNEL_ID] [Extra Argument]"
}

// NewInstructionParser - Gets an instruction parser.
func NewInstructionParser() InstructionParser {
	return twilioSMSParser{}
}
