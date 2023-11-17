package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestTorontoTimeHandler tests the TorontoTimeHandler function
func TestTorontoWeatherHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/toronto-weather", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(weatherHandler)

	handler.ServeHTTP(rr, req)

	//Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//Check the response body is in the correct format
	var response map[string]string
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal("Failed to decode response body")
	}

	if _, ok := response["current_weather_toronto"]; !ok {
		t.Errorf("handler returned unexpected body: key 'current_weather_toronto' not found in response")
	}
}
