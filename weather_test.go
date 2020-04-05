package main

import (
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
	handler := http.HandlerFunc(weather)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("unexpected status code:\ngot: %v,\nwant: %v",
			status, http.StatusOK)
	}

	expected := "some weather report"
	if rr.Body.String() != expected {
		t.Errorf("unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
