package startup

import (
	"flag"
	"math/rand"
	"time"
)

// DefaultPort is the default port to use by web service
const DefaultPort = 8080

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetPort returns port to use obtained from user or DefaultPort
func GetPort() (port int) {
	flag.IntVar(&port, "port", DefaultPort, "port to connect this server")
	flag.Parse()
	return port
}
