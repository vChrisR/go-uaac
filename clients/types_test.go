package clients

import (
	"encoding/json"
	"testing"
)

func TestJSONUnmarshallOauthClient(t *testing.T) {
	responseBody := []byte(`{
      "client_id" : "foo",
      "name" : "Foo Client Name",
      "scope" : ["uaa.none"],
      "resource_ids" : ["none"],
      "authorities" : ["cloud_controller.read","cloud_controller.write","scim.read"],
      "authorized_grant_types" : ["client_credentials"],
      "lastModified" : 1426260091149
    }`)

	var client OauthClient

	err := json.Unmarshal(responseBody, &client)
	if err != nil {
		t.Errorf("Failed to unmarshall json to OauthClient: %v", err)
	}

	if len(client.ID) == 0 {
		t.Errorf("client.ID field empty after unmarshalling from json")
	}
	if len(client.Name) == 0 {
		t.Errorf("client.Name field empty after unmarshalling from json")
	}

	if len(client.Scope) == 0 {
		t.Errorf("client.Scope field empty after unmarshalling from json")
	}
	if len(client.ResourceIDs) == 0 {
		t.Errorf("client.ResourceIDs field empty after unmarshalling from json")
	}
	if len(client.Authorities) == 0 {
		t.Errorf("client.Authorities field empty after unmarshalling from json")
	}
	if len(client.AuthorizedGrantTypes) == 0 {
		t.Errorf("client.AuthorizedGrantTypes field empty after unmarshalling from json")
	}
	if client.LastModified == 0 {
		t.Errorf("client.LastModified field empty after unmarshalling from json")
	}
}

func TestJSONUnmarshallOauthClients(t *testing.T) {
	responseBody := []byte(`{
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

	var clients OauthClients

	err := json.Unmarshal(responseBody, &clients)
	if err != nil {
		t.Errorf("Failed to unmarshall json to OauthClients: %v", err)
	}
}
