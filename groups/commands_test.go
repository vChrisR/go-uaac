package groups

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uaa "github.com/pivotalservices/go-uaac"
)

func getUaac(apiUrl string) (uaa.Client, error) {
	clientConfig := &uaa.ClientConfig{
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

func TestListGroupsCommand(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
      "resources" : [ {
        "id" : "0393bbd5-3512-414a-af09-bb2f59dede39",
        "meta" : {
          "version" : 1,
          "created" : "2016-07-12T23:00:54.467Z",
          "lastModified" : "2016-07-12T23:00:54.504Z"
        },
        "displayName" : "Cooler Group Name for Update",
        "zoneId" : "uaa",
        "description" : "the cool group",
        "members" : [ {
          "origin" : "uaa",
          "type" : "USER",
          "value" : "344bc255-8809-42ca-a0a6-8602213ef285"
        } ],
        "schemas" : [ "urn:scim:schemas:core:1.0" ]
      } ],
      "startIndex" : 1,
      "itemsPerPage" : 50,
      "totalResults" : 1,
      "schemas" : [ "urn:scim:schemas:core:1.0" ]
    }`)
	}))
	defer ts.Close()

	uaac, err := getUaac(ts.URL)
	if err != nil {
		t.Errorf("Failed to get uaa client; %v", err)
		return
	}

	var listGroupsResponse ListResponse
	command := NewListGroupsCommand(uaac, &listGroupsResponse)

	if err := command.Execute(); err != nil {
		t.Errorf("Failed to get Groups: %v", err)
	}

	if len(listGroupsResponse.Groups) == 0 {
		t.Error("ListResponse.Groups was empty")
	}
}
