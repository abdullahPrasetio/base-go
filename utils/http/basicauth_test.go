package http

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBasicAuthString(t *testing.T) {
	// Set up test cases
	tests := []struct {
		name     string
		username string
		password string
		expected string
	}{
		{
			name:     "BasicAuthCase",
			username: "user",
			password: "pass",
			expected: "Basic dXNlcjpwYXNz",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the BasicAuthString function
			result := BasicAuthString(tt.username, tt.password)

			// Decode the result for comparison
			decodedResult, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(result, "Basic "))
			assert.NoError(t, err, "Error decoding result")

			decodedParts := strings.SplitN(string(decodedResult), ":", 2)
			assert.Equal(t, 2, len(decodedParts), "Decoded result does not contain username and password")
			// Construct the expected result
			expected := tt.username + ":" + tt.password
			expectedBase64 := "Basic " + base64.StdEncoding.EncodeToString([]byte(expected))

			// Check if the result matches the expected output
			assert.Equal(t, expectedBase64, result, "BasicAuthString result does not match the expected output")

			// Check if the decoded result matches the expected value
			assert.Equal(t, tt.username, decodedParts[0], "Decoded username does not match the expected value")
			assert.Equal(t, tt.password, decodedParts[1], "Decoded password does not match the expected value")
		})
	}
}
