package identityzones

import (
	"encoding/json"
	"testing"
)

func TestJSONUnmarshallIdentityZone(t *testing.T) {
	responseBody := []byte(`{
        "id": "dummy-id",
        "subdomain": "dummy-subdomain",
        "name": "dummy-name",
        "version": 1,
        "description": "Dummy Description",
        "created": 946710000000,
        "last_modified": 946710000000
    }`)

	var zone IdentityZone

	err := json.Unmarshal(responseBody, &zone)
	if err != nil {
		t.Errorf("Failed to unmarshall json to Zone: %v", err)
	}

	if len(zone.ID) == 0 {
		t.Errorf("Zone.ID field empty after unmarshalling from json")
	}
	if len(zone.Subdomain) == 0 {
		t.Errorf("Zone.Subdomain field empty after unmarshalling from json")
	}
	if len(zone.Name) == 0 {
		t.Errorf("Zone.Name field empty after unmarshalling from json")
	}
	if len(zone.Description) == 0 {
		t.Errorf("Zone.Description field empty after unmarshalling from json")
	}
	if zone.Version == 0 {
		t.Errorf("Zone.Version field empty after unmarshalling from json")
	}
	if zone.Created == 0 {
		t.Errorf("Zone.Created field empty after unmarshalling from json")
	}
	if zone.LastModified == 0 {
		t.Errorf("Zone.LastModified field empty after unmarshalling from json")
	}
}
