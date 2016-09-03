package groups

import (
	"encoding/json"
	"testing"
)

func TestJSONUnmarshallGroups(t *testing.T) {
	responseBody := []byte(`{
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

	var groupsResponse ListResponse

	err := json.Unmarshal(responseBody, &groupsResponse)
	if err != nil {
		t.Errorf("Failed to unmarshall json to ListResponse: %v", err)
	}

	if len(groupsResponse.Groups) == 0 {
		t.Errorf("ListResponse.Groups was empty after unmarshalling from json")
	}
	if groupsResponse.ItemsPerPage == 0 {
		t.Errorf("ListGroups.ItemsPerPage field empty after unmarshalling from json")
	}
	if groupsResponse.StartIndex == 0 {
		t.Errorf("ListGroups.StartIndex field empty after unmarshalling from json")
	}
	if groupsResponse.TotalResults == 0 {
		t.Errorf("ListGroups.TotalResults field empty after unmarshalling from json")
	}

	for i, group := range groupsResponse.Groups {
		if len(group.Description) == 0 {
			t.Errorf("groupsResponse.Groups[%d].Description was empty", i)
		}
		if len(group.DisplayName) == 0 {
			t.Errorf("groupsResponse.Groups[%d].DisplayName was empty", i)
		}
		if len(group.ID) == 0 {
			t.Errorf("groupsResponse.Groups[%d].ID was empty", i)
		}
		if len(group.ZoneID) == 0 {
			t.Errorf("groupsResponse.Groups[%d].ZoneID was empty", i)
		}
		if group.Meta == nil {
			t.Errorf("group Metadata was nil")
		}
		if len(group.Meta.Created) == 0 {
			t.Errorf("group.Meta.Created was empty")
		}
		if len(group.Meta.LastModified) == 0 {
			t.Errorf("group.Meta.LastModified was empty")
		}
		if group.Meta.Version == 0 {
			t.Errorf("group.Meta.Version was empty")
		}
		if len(group.Members) == 0 {
			t.Errorf("Members was empty")
		}

		for _, member := range group.Members {
			if len(member.Origin) == 0 {
				t.Errorf("Member.Origin was empty")
			}
			if len(member.Type) == 0 {
				t.Errorf("Member.Type was empty")
			}
			if len(member.Value) == 0 {
				t.Errorf("Member.Value was empty")
			}
		}
	}
}
