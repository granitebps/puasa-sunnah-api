package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClassifyRequestSource(t *testing.T) {
	tests := []struct {
		name      string
		apiClient string
		expected  string
	}{
		{
			name:      "should classify official application",
			apiClient: officialAPIClient,
			expected:  requestSourceOfficialApp,
		},
		{
			name:      "should classify missing header as public",
			apiClient: "",
			expected:  requestSourcePublic,
		},
		{
			name:      "should classify unknown client as public",
			apiClient: "unknown-client",
			expected:  requestSourcePublic,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, classifyRequestSource(tt.apiClient))
		})
	}
}
