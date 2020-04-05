package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const defaultPort = 8080

func getPort() (port int) {
	flag.IntVar(&port, "port", defaultPort, "port to connect this server")
	flag.Parse()
	return port
}

func weather(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "some weather report")
}

func main() {
	port := getPort()

	http.HandleFunc("/weather", weather)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("unexpected error while ListenAndServe: %s", err)
	}
}
