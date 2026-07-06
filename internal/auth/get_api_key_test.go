package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		wantKey string
		wantErr bool
	}{
		{
			name:    "returns api key when authorization header is valid",
			headers: http.Header{"Authorization": []string{"ApiKey secret-key"}},
			wantKey: "secret-key",
		},
		{
			name:    "returns error when authorization header is missing",
			headers: http.Header{},
			wantErr: true,
		},
		{
			name:    "returns error when authorization header is malformed",
			headers: http.Header{"Authorization": []string{"Bearer secret-key"}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.headers)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected an error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("GetAPIKey() returned error %v", err)
			}

			if gotKey != tt.wantKey {
				t.Fatalf("GetAPIKey() = %q, want %q", gotKey, tt.wantKey)
			}
		})
	}
}
