package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/goofinator/weatherHttp/internal/controllers"
)

// Run starts the web service
func Run(port int) {
	http.HandleFunc("/weather", controllers.WeatherHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("unexpected error while ListenAndServe: %s", err)
	}
}
