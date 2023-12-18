package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// APICall struct represents the parameters needed for an API call
type APICall struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    string
}

// PerformAPICall function makes a dynamic API call based on the provided parameters
func PerformAPICall(apiCall APICall) ([]byte, error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest(apiCall.Method, apiCall.URL, bytes.NewBuffer([]byte(apiCall.Body)))
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range apiCall.Headers {
		req.Header.Set(key, value)
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
