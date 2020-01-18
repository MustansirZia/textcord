package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mustansirzia/simcord/api"
)

func main() {
	http.HandleFunc("/callback", api.Callback)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	fmt.Printf("Listening on :%s!\n", port)
	print(http.ListenAndServe(":"+port, nil).Error())
}
