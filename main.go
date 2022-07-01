package main

import (
	"github.com/google/go-github/github"
	"log"
	"net/http"
	"os"
)

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
	log.Printf("Event: %#v\n\n", event)
}

func main() {
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
