package masking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMasking(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		start    int
		end      int
		expected string
	}{
		{
			name:     "MaskMiddle",
			text:     "1234567890",
			start:    2,
			end:      7,
			expected: "12******90",
		},
		{
			name:     "MaskAll",
			text:     "abcdefgh",
			start:    0,
			end:      7,
			expected: "********",
		},
		{
			name:     "MaskNone",
			text:     "password",
			start:    0,
			end:      1,
			expected: "password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Masking(tt.text, tt.start, tt.end)
			assert.Equal(t, tt.expected, result, "Masking result should match expected value")
		})
	}
}
