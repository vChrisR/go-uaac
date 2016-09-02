package uaa

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
      "access_token" : "4a53a3331b2445cfaca43c9af00439e8",
      "token_type" : "bearer",
      "expires_in" : 43199,
      "scope" : "clients.read emails.write scim.userids password.write idps.write notifications.write login scim.write critical_notifications.write",
      "jti" : "4a53a3331b2445cfaca43c9af00439e8"
    }`)
	}))
	defer ts.Close()

	if _, err := getUaac(ts.URL); err != nil {
		t.Errorf("Failed to initialize uaa client; %v", err)
		return
	}
}

func getUaac(apiUrl string) (Client, error) {
	clientConfig := &ClientConfig{
		ApiAddress:   apiUrl,
		ClientID:     "fake-client",
		ClientSecret: "big-secret",
	}

	oauthClient, err := NewClient(clientConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize oauth client; %v", err)
	}

	uaac := &uaaClient{Oauth2Client: oauthClient}
	return uaac, nil
}
