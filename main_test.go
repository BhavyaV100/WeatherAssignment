package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Declare apiURL as a global variable accessible in the test file
var apiURLTest string

func TestGetWeatherData(t *testing.T) {
	// Mock OpenWeatherMap API response
	mockResponse := `{"name":"Toronto","main":{"temp":280.75,"humidity":84},"weather":[{"description":"Mist"}]}`

	// Create a mock server to simulate the OpenWeatherMap API
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	// Override the API URL in the main code with the mock server URL
	apiURLTest = mockServer.URL

	// Create a request to the /weather endpoint
	req := httptest.NewRequest("GET", "/weather", nil)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the getWeatherData function with the mock request and recorder
	getWeatherData(rr, req)

	// Check if the status code is HTTP 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Example assertions for the response body
	expectedResponseBody := `{"current_location":"Toronto","weather":{"description":"mist","temperature":280.75,"humidity":84}}`
	if rr.Body.String() != expectedResponseBody {
		t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedResponseBody)
	}
}
