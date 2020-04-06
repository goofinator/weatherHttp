package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/goofinator/weatherHttp/internal/utils"
)

// MaxFeelingLen is the maximum legth of randomly generated string
const MaxFeelingLen = 80

// Weather struct holds weather data
type Weather struct {
	ID      int    `json:"id"`
	Feeling string `json:"feeling"`
}

// WeatherHandler is a handler for /weather endpoint
func WeatherHandler(w http.ResponseWriter, req *http.Request) {
	weatherValue := newRandomWeather()

	prettyJSON, err := json.MarshalIndent(*weatherValue, "", "    ")
	if err != nil {
		msg := fmt.Sprintf("error on json.MarshalIndent: %s.", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(prettyJSON))
}

func newRandomWeather() *Weather {
	// +1 to produce at least 1 symbol
	n := rand.Intn(MaxFeelingLen-1) + 1

	weatherValue := &Weather{
		ID:      rand.Int(),
		Feeling: utils.RandString(n),
	}

	return weatherValue
}
