package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	defaultPort   = 8080
	maxFeelingLen = 80
)

type weather struct {
	ID      int    `json:"id"`
	Feeling string `json:"feeling"`
}

func getPort() (port int) {
	flag.IntVar(&port, "port", defaultPort, "port to connect this server")
	flag.Parse()
	return port
}

func weatherHandler(w http.ResponseWriter, req *http.Request) {
	weatherValue := newRandomWeather()

	prettyJSON, err := json.MarshalIndent(*weatherValue, "", "    ")
	if err != nil {
		msg := fmt.Sprintf("error on json.MarshalIndent: %s.", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(prettyJSON))
}

func newRandomWeather() *weather {
	// +1 to produce at least 1 symbol
	n := rand.Intn(maxFeelingLen-1) + 1

	weatherValue := &weather{
		ID:      rand.Int(),
		Feeling: randString(n),
	}

	return weatherValue
}

func randString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	port := getPort()
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/weather", weatherHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("unexpected error while ListenAndServe: %s", err)
	}
}
