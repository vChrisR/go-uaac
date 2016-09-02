package uaa

import (
	"encoding/json"
	"testing"
)

func TestJSONUnmarshallAccessToken(t *testing.T) {
	responseBody := []byte(`{
      "access_token":"dummy-access-token",
      "token_type":"dummy-token-type",
      "expires_in":3600,
      "scope":"dummy-scope",
      "jti":"dummy-jti"
    }`)

	var token AccessToken

	err := json.Unmarshal(responseBody, &token)
	if err != nil {
		t.Errorf("Failed to unmarshall json to AccessToken: %v", err)
	}

	if len(token.Token) == 0 {
		t.Errorf("AccessToken.Token field empty after unmarshalling from json")
	}

	if len(token.Type) == 0 {
		t.Errorf("AccessToken.Type field empty after unmarshalling from json")
	}

	if token.ExpiresIn == 0 {
		t.Errorf("AccessToken.ExpiresIn field empty after unmarshalling from json")
	}

	if len(token.Scope) == 0 {
		t.Errorf("AccessToken.Scope field empty after unmarshalling from json")
	}

	if len(token.JTI) == 0 {
		t.Errorf("AccessToken.JTI field empty after unmarshalling from json")
	}
}
