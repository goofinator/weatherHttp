package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWeatherHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatalf("unexpected fail of NewRequest: %s", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(weatherHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("unexpected status code:\nwant: %v\ngot: %v",
			http.StatusOK, status)
	}

	testWeather(t, rr.Body.String())
}

func testWeather(t *testing.T, body string) {
	weatherValue := weather{}
	err := json.Unmarshal([]byte(body), &weatherValue)
	if err != nil {
		t.Errorf("unexpected error on json.Unmarshal:\nwant: %v\ngot: %v",
			nil, err)
	}
}
