package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const defaultPort = 8080

type weather struct {
	ID      int    `json:"id"`
	Feeling string `json:"feeling"`
}

func getPort() (port int) {
	flag.IntVar(&port, "port", defaultPort, "port to connect this server")
	flag.Parse()
	return port
}

// Weather handles the request on /weather endpoint
func weatherHandler(w http.ResponseWriter, req *http.Request) {
	weatherValue := weather{
		ID:      123,
		Feeling: "shdjkah182y3jasbnd",
	}

	prettyJSON, err := json.MarshalIndent(weatherValue, "", "    ")
	if err != nil {
		msg := fmt.Sprintf("error on json.MarshalIndent: %s.", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(prettyJSON))
}

func main() {
	port := getPort()

	http.HandleFunc("/weather", weatherHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("unexpected error while ListenAndServe: %s", err)
	}
}
