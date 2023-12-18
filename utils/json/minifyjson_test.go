package json

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMinifyJson(t *testing.T) {
	// Set up test cases
	tests := []struct {
		name     string
		input    interface{}
		expected map[string]interface{}
	}{
		{
			name: "SimpleCase",
			input: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			expected: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "LongStringValue",
			input: map[string]interface{}{
				"key1": "value1",
				"key2": "very long string that exceeds 10000 characters " + strings.Repeat("a", 10000),
			},
			expected: map[string]interface{}{
				"key1": "value1",
			},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinifyJson(tt.input)

			// Check if the result matches the expected output
			assert.Equal(t, tt.expected, result, "MinifyJson result does not match the expected output")
		})
	}
}
