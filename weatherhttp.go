package main

import (
	"flag"
	"fmt"
)

const defaultPort = 8080

func getPort() (port int) {
	flag.IntVar(&port, "port", defaultPort, "port to connect this server")
	flag.Parse()
	return port
}

func main() {
	port := getPort()
	fmt.Println(port)
}
