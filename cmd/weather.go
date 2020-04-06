package main

import (
	"github.com/goofinator/weatherHttp/internal/init/startup"
	"github.com/goofinator/weatherHttp/internal/web"
)

func main() {
	port := startup.GetPort()

	web.Run(port)
}
