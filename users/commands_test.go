package users

import (
	"fmt"
	"io/ioutil"
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

func TestGetUsersCommandUaa20Model(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile("../testdata/2.0.0/GET.Users.json")

		if err != nil {
			panic("Failed to read ../testdata/2.0.0/GET.Users.json: " + err.Error())
		}

		w.Write(data)
	}))
	defer ts.Close()

	uaac, err := getUaac(ts.URL)
	if err != nil {
		t.Errorf("Failed to get uaa client; %v", err)
		return
	}

	var users Users
	command := NewGetUsersCommand(uaac, &users)

	if err := command.Execute(); err != nil {
		t.Errorf("Failed to get Users: %v", err)
	}

	if len(users.Users) == 0 {
		t.Error("[]Users was empty")
	}
}

func TestCreateUserCommand(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
		  "id" : "f14b29c0-89bd-4aef-bf50-6a3bef46d265",
		  "externalId" : "test-user",
		  "meta" : {
		    "version" : 0,
		    "created" : "2016-07-12T23:00:56.650Z",
		    "lastModified" : "2016-07-12T23:00:56.650Z"
		  },
		  "userName" : "lVHky5@test.org",
		  "name" : {
		    "familyName" : "family name",
		    "givenName" : "given name"
		  },
		  "emails" : [ {
		    "value" : "lVHky5@test.org",
		    "primary" : false
		  } ],
		  "groups" : [ {
		    "value" : "d42cfca3-7108-49b4-91d8-f76fad1328b3",
		    "display" : "scim.userids",
		    "type" : "DIRECT"
		  }, {
		    "value" : "b066910b-7890-4a76-befe-2dafd720afd0",
		    "display" : "user_attributes",
		    "type" : "DIRECT"
		  }, {
		    "value" : "a1447fd3-3f9b-4eb4-a268-e175d8d5ecb8",
		    "display" : "approvals.me",
		    "type" : "DIRECT"
		  }, {
		    "value" : "9c2b6aa0-c692-4f7d-ac77-bdccb15b5626",
		    "display" : "uaa.offline_token",
		    "type" : "DIRECT"
		  }, {
		    "value" : "3dbaa4e7-d7fa-4b5d-89d9-3fa38b6bdcde",
		    "display" : "cloud_controller.write",
		    "type" : "DIRECT"
		  }, {
		    "value" : "e1462b39-bd1b-40e2-a564-9b5e9320a014",
		    "display" : "password.write",
		    "type" : "DIRECT"
		  }, {
		    "value" : "c04b107a-6829-46db-a943-0b39f2bc2440",
		    "display" : "uaa.user",
		    "type" : "DIRECT"
		  }, {
		    "value" : "4837fd28-0d3c-4cab-a270-b8a3af25a939",
		    "display" : "cloud_controller.read",
		    "type" : "DIRECT"
		  }, {
		    "value" : "7f827a21-3d67-4c26-ab49-74709ef5e97e",
		    "display" : "roles",
		    "type" : "DIRECT"
		  }, {
		    "value" : "23b9ad28-87dd-4222-be37-200e46692cff",
		    "display" : "openid",
		    "type" : "DIRECT"
		  }, {
		    "value" : "c0f8e11f-ef98-42ca-8406-42ca7537040c",
		    "display" : "cloud_controller_service_permissions.read",
		    "type" : "DIRECT"
		  }, {
		    "value" : "881a903d-8b32-41b8-84a7-62ce6f2ca09f",
		    "display" : "profile",
		    "type" : "DIRECT"
		  }, {
		    "value" : "43a5a6e3-82c2-4458-8cd1-fb9b98caeeb5",
		    "display" : "oauth.approvals",
		    "type" : "DIRECT"
		  }, {
		    "value" : "6ad7ab9f-abc1-41c9-b668-e31bddc2a2b1",
		    "display" : "scim.me",
		    "type" : "DIRECT"
		  } ],
		  "approvals" : [ ],
		  "phoneNumbers" : [ {
		    "value" : "5555555555"
		  } ],
		  "active" : true,
		  "verified" : true,
		  "origin" : "uaa",
		  "zoneId" : "uaa",
		  "passwordLastModified" : "2016-07-12T23:00:56.000Z",
		  "schemas" : [ "urn:scim:schemas:core:1.0" ]
		}`)
	}))
	defer ts.Close()

	uaac, err := getUaac(ts.URL)
	if err != nil {
		t.Errorf("Failed to get uaa client; %v", err)
		return
	}

	user := &User{
		Username:   "fake-username",
		ExternalID: "fake-external-id",
	}

	command := NewCreateUserCommand(uaac, user)
	if err := command.Execute(); err != nil {
		t.Errorf("Failed execute CreateUserCommand: %v", err)
	}

	fmt.Println("guid: ", user.GUID)

	if len(user.GUID) == 0 {
		t.Errorf("Empty UserGuid after running CreateUserCommand %v", user)
	}
}
