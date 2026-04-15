package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		header      string
		expectedKey string
		expectErr   bool
	}{
		{
			name:      "no authorization header",
			header:    "",
			expectErr: true,
		},
		{
			name:      "malformed header - wrong prefix",
			header:    "Bearer sometoken",
			expectErr: true,
		},
		{
			name:        "valid ApiKey header",
			header:      "ApiKey mysecretkey123",
			expectedKey: "mysecretkey123",
			expectErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.header != "" {
				headers.Set("Authorization", tt.header)
			}

			key, err := GetAPIKey(headers)

			if tt.expectErr && err == nil {
				t.Errorf("expected an error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("did not expect error but got: %v", err)
			}
			if key != tt.expectedKey {
				t.Errorf("expected key %q but got %q", tt.expectedKey, key)
			}
		})
	}
}
