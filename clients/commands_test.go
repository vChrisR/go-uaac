package clients

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

func TestListOauthClientsCommand(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
      "resources" : [ {
        "scope" : [ "clients.read", "clients.write" ],
        "client_id" : "vleI0g",
        "resource_ids" : [ "none" ],
        "authorized_grant_types" : [ "client_credentials" ],
        "redirect_uri" : [ "http*://ant.path.wildcard/**/passback/*", "http://test1.com" ],
        "autoapprove" : [ ],
        "authorities" : [ "clients.read", "clients.write" ],
        "lastModified" : 1463595914005
      } ],
      "startIndex" : 1,
      "itemsPerPage" : 1,
      "totalResults" : 1,
      "schemas" : [ "http://cloudfoundry.org/schema/scim/oauth-clients-1.0" ]
    }`)
	}))
	defer ts.Close()

	uaac, err := getUaac(ts.URL)
	if err != nil {
		t.Errorf("Failed to get uaa client; %v", err)
		return
	}

	var clients OauthClients
	command := NewListClientsCommand(uaac, &clients)

	if err := command.Execute(); err != nil {
		t.Errorf("Failed to get OauthClients: %v", err)
	}

	if len(clients.Clients) == 0 {
		t.Error("[]OauthClient was empty")
	}
}
