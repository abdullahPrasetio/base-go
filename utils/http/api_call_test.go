package http

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPerformAPICall(t *testing.T) {
	// Mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with a mock JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer server.Close()

	// Test case 1: Successful GET request
	apiCall := APICall{
		URL:    server.URL,
		Method: "GET",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "",
	}

	response, err := PerformAPICall(apiCall)
	assert.NoError(t, err, "Unexpected error in successful API call")
	assert.Equal(t, `{"message": "success"}`, string(response), "Unexpected response in successful API call")

	// Test case 2: Error in request creation (invalid URL)
	apiCall.URL = ":invalid-url"
	_, err = PerformAPICall(apiCall)
	assert.Error(t, err, "parse \":invalid-url\": missing protocol scheme\n")

	//// Test case 3: Error in request creation (invalid method)
	//apiCall.URL = server.URL
	//apiCall.Method = "INVALID_METHOD"
	//_, err = PerformAPICall(apiCall)
	//
	//fmt.Println(err)
	//assert.Error(t, err, "Expected error in request creation, but got none")
	//
	//// Test case 4: Error in request creation (invalid body)
	//apiCall.Method = "POST"
	//apiCall.Body = "invalid-json-body"
	//_, err = PerformAPICall(apiCall)
	//assert.Error(t, err, "Expected error in request creation, but got none")
	// Test case 5: Error in request execution
	apiCall.Method = "GET"
	apiCall.URL = "http://nonexistent"
	_, err = PerformAPICall(apiCall)
	assert.Error(t, err, "Expected error in request execution, but got none")

}
func TestPerformAPICall_ErrorReadingBody(t *testing.T) {
	// Mock HTTP server that always returns a response causing an error when reading the body
	errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate an error by setting inconsistent Content-Length header
		w.Header().Set("Content-Length", "10")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer errorServer.Close()

	// Test case: Error in reading response body (server returns a response causing an error)
	apiCall := APICall{
		URL:    errorServer.URL,
		Method: "GET",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "",
	}

	response, err := PerformAPICall(apiCall)
	assert.Error(t, err, "Expected error in reading response body, but got none")
	assert.Nil(t, response, "Expected response to be nil when there's an error in reading body")
}

func TestPerformAPICall_NoErrorReadingBody(t *testing.T) {
	// Mock HTTP server that returns a normal response
	normalServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer normalServer.Close()

	// Test case: No error in reading response body
	apiCall := APICall{
		URL:    normalServer.URL,
		Method: "GET",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "",
	}

	response, err := PerformAPICall(apiCall)
	assert.NoError(t, err, "Unexpected error in reading response body")
	assert.NotNil(t, response, "Expected response to be non-nil when there's no error in reading body")
}
