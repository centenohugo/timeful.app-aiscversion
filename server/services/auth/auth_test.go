package auth

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

// Google answers a rejected refresh_token grant with HTTP 200 and an `error` field rather than a
// non-2xx status, so the only way to tell success from failure is to inspect the payload.
func TestRefreshSucceeded(t *testing.T) {
	tests := []struct {
		name string
		res  AccessTokenResponse
		want bool
	}{
		{
			name: "valid refresh",
			res:  AccessTokenResponse{AccessToken: "ya29.token", ExpiresIn: 3599},
			want: true,
		},
		{
			name: "expired refresh token (testing-mode 7 day expiry)",
			res:  AccessTokenResponse{Error: bson.M{"error": "invalid_grant"}},
			want: false,
		},
		{
			name: "revoked consent",
			res:  AccessTokenResponse{AccessToken: "", Error: bson.M{"error": "invalid_grant", "error_description": "Token has been expired or revoked."}},
			want: false,
		},
		{
			name: "empty response with no error field",
			res:  AccessTokenResponse{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RefreshSucceeded(tt.res); got != tt.want {
				t.Errorf("RefreshSucceeded(%+v) = %v, want %v", tt.res, got, tt.want)
			}
		})
	}
}
