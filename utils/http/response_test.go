package http

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAPIResponseSuccess(t *testing.T) {
	// Set up test cases
	tests := []struct {
		name          string
		responseDesc  string
		responseData  interface{}
		expectedCode  string
		expectedDesc  string
		expectedData  interface{}
		expectedError interface{}
	}{
		{
			name:          "SuccessCase",
			responseDesc:  "Success",
			responseData:  map[string]string{"key": "value"},
			expectedCode:  "00",
			expectedDesc:  "Success",
			expectedData:  map[string]string{"key": "value"},
			expectedError: nil,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the APIResponseSuccess function
			result := APIResponseSuccess(tt.responseDesc, tt.responseData)

			// Convert the result to JSON for comparison
			resultJSON, _ := json.Marshal(result)

			// Convert the expected result to JSON for comparison
			expectedJSON, _ := json.Marshal(Response{
				ResponseCode:   tt.expectedCode,
				ResponseDesc:   tt.expectedDesc,
				ResponseData:   tt.expectedData,
				ResponseErrors: tt.expectedError,
			})

			// Check if the result matches the expected output
			assert.JSONEq(t, string(expectedJSON), string(resultJSON), "APIResponseSuccess result does not match the expected output")
		})
	}
}

func TestAPIResponseError(t *testing.T) {
	// Set up test cases
	tests := []struct {
		name          string
		responseDesc  string
		responseCode  string
		errors        interface{}
		expectedCode  string
		expectedDesc  string
		expectedData  interface{}
		expectedError interface{}
	}{
		{
			name:          "ErrorCase",
			responseDesc:  "Error",
			responseCode:  "99",
			errors:        map[string]string{"errorKey": "errorMessage"},
			expectedCode:  "99",
			expectedDesc:  "Error",
			expectedData:  nil,
			expectedError: map[string]string{"errorKey": "errorMessage"},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the APIResponseError function
			result := APIResponseError(tt.responseDesc, tt.responseCode, tt.errors)

			// Convert the result to JSON for comparison
			resultJSON, _ := json.Marshal(result)

			// Convert the expected result to JSON for comparison
			expectedJSON, _ := json.Marshal(Response{
				ResponseCode:   tt.expectedCode,
				ResponseDesc:   tt.expectedDesc,
				ResponseData:   tt.expectedData,
				ResponseErrors: tt.expectedError,
			})

			// Check if the result matches the expected output
			assert.JSONEq(t, string(expectedJSON), string(resultJSON), "APIResponseError result does not match the expected output")
		})
	}
}
