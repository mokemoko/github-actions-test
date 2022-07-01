package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"log"
	"net/http"
	"os"
)

func logEvent(data interface{}) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("Event: %#v\nError: %v\n\n", data, err)
	}
	fmt.Printf("%T: %s\n\n", data, string(bytes))
}
func logError(r *http.Request, err error) {
	log.Printf("Headers: %v\nError: %v\n\n", r.Header, err)
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := github.ValidatePayload(r, []byte(os.Getenv("GHA_SECRET")))
	defer r.Body.Close()
	if err != nil {
		logError(r, err)
		return
	}

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		logError(r, err)
		return
	}
	logEvent(event)
}

func main() {
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
