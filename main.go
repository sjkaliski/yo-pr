package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/sjkaliski/go-yo"
)

var (
	yoToken  = flag.String("yo_token", "", "Yo API Token.")
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
	if data.PullRequest.Action == "open" {
		log.Println(fmt.Sprintf("new pr opened: %s", data.PullRequest.Url))

		if err := yoClient.YoAll(); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	rw.WriteHeader(http.StatusOK)
}

func main() {
	flag.Parse()
	if *yoToken == "" {
		log.Fatal("yo_token required.")
	}

	yoClient = yo.NewClient(*yoToken)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/pr", prHandler)
	http.ListenAndServe(":8080", nil)
}
