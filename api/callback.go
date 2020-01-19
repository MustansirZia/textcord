package api

import (
	"fmt"
	"net/http"

	"github.com/mustansirzia/simcord/handler"
	"github.com/mustansirzia/simcord/parser"
)

var instructionParser parser.InstructionParser

func init() {
	instructionParser = parser.NewInstructionParser()
}

// Callback - Single HTTP POST endpoint which will be called by SMS reception webhooks.
func Callback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	instruction, err := instructionParser.Parse(r)
	if err != nil {
		if err == parser.ErrInvalidRequest {
			http.Error(
				w,
				fmt.Sprintf("Invalid. Correct Pattern - %s", instructionParser.Pattern()),
				http.StatusBadRequest,
			)
			return
		}
		handleError(w, err)
		return
	}
	response, err := handler.HandleInstruction(instruction)
	if err != nil {
		handleError(w, err)
		return
	}
	sendResponse(w, response)
}

func handleError(w http.ResponseWriter, err error) {
	fmt.Println(err.Error())
	w.Header().Add("content-type", "text/plain; charset=utf-8")
	// Sending a 200 and not a 500 because we want the response to be
	// sent back to the client.
	fmt.Fprint(w, err.Error())
	// http.Error(w, err.Error(), http.StatusInternalServerError)
}

func sendResponse(w http.ResponseWriter, response string) {
	w.Header().Add("content-type", "text/plain; charset=utf-8")
	fmt.Fprint(w, response)
}
