package times

import (
	"github.com/abdullahPrasetio/base-go/configs"
	"github.com/abdullahPrasetio/base-go/utils/log"
	"github.com/abdullahPrasetio/base-go/utils/times"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init() {
	configs.LoadConfig("./../../..")

	log.LoadLogger()
}

func TestGetTimeNow(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected string
	}{
		{
			name:     "DefaultFormat",
			format:   "",
			expected: "2006-01-02 15:04:05",
		},
		{
			name:     "CustomFormat",
			format:   "2006-01-02",
			expected: "2006-01-02",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := times.GetTimeNow(tt.format)

			// Assertion for ValueTime
			assert.WithinDuration(t, time.Now(), result.ValueTime, 1*time.Second, "ValueTime should be within 1 second of current time")

			// Assertion for ValueString
			expectedString := result.ValueTime.Format(tt.expected)
			assert.Equal(t, expectedString, result.ValueString, "ValueString should match formatted ValueTime")
		})
	}
}
