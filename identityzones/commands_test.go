package identityzones

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dave-malone/oauth"
	uaa "github.com/pivotalservices/go-uaac"
)

func getUaac(apiUrl string) (uaa.Client, error) {
	clientConfig := &oauth.ClientConfig{
		ApiAddress:   apiUrl,
		ClientID:     "fake-client",
		ClientSecret: "big-secret",
	}

	uaac, err := uaa.NewClient(clientConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize uaa client; %v", err)
	}
	return uaac, nil
}

func TestListIdentityZonesCommand(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `[{
        "id": "dummy-id",
        "subdomain": "dummy-subdomain",
        "name": "dummy-name",
        "version": 1,
        "description": "Dummy Description",
        "created": 946710000000,
        "last_modified": 946710000000
    }]`)
	}))
	defer ts.Close()

	uaac, err := getUaac(ts.URL)
	if err != nil {
		t.Errorf("Failed to get uaa client; %v", err)
		return
	}

	var zones []*IdentityZone
	command := NewListIdentityZonesCommand(uaac, &zones)

	if err := command.Execute(); err != nil {
		t.Errorf("Failed to get IdentityZones: %v", err)
	}

	if len(zones) == 0 {
		t.Error("[]*IdentityZone was empty")
	}
}
