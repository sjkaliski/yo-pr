package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sjkaliski/go-yo"
)

var (
	yoClient *yo.Client
)

type hook struct {
	Action      string `json:"action"`
	PullRequest pr     `json:"pull_request"`
}

type pr struct {
	Url  string `json:"url"`
	Body string `json:"body"`
}

// GET /, handles index page.
func indexHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Yo-PR!")
}

// GET /pr, handles github webook post request.
func prHandler(rw http.ResponseWriter, req *http.Request) {
	var data hook
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&data); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If the type of PR action is "open", send a Yo to
	// all subscribers.
	if data.Action == "open" {
		log.Println(fmt.Sprintf("new pr opened: %s", data.PullRequest.Url))

		if err := yoClient.YoAll(); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	rw.WriteHeader(http.StatusOK)
}

func main() {
	yoToken := os.Getenv("YO_TOKEN")
	if yoToken == "" {
		log.Fatal("YO_TOKEN required.")
	}

	yoClient = yo.NewClient(yoToken)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/pr", prHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
