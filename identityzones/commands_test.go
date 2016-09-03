package identityzones

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

func TestGetIdentityZoneByIDCommand(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
		  "id" : "twiglet-get",
		  "subdomain" : "twiglet-get",
		  "config" : {
		    "tokenPolicy" : {
		      "accessTokenValidity" : -1,
		      "refreshTokenValidity" : -1,
		      "jwtRevocable" : false,
		      "activeKeyId" : null,
		      "keys" : { }
		    },
		    "samlConfig" : {
		      "assertionSigned" : true,
		      "requestSigned" : true,
		      "wantAssertionSigned" : false,
		      "wantAuthnRequestSigned" : false,
		      "assertionTimeToLiveSeconds" : 600,
		      "certificate" : null,
		      "privateKey" : null,
		      "privateKeyPassword" : null
		    },
		    "links" : {
		      "logout" : {
		        "redirectUrl" : "/login",
		        "redirectParameterName" : "redirect",
		        "disableRedirectParameter" : true,
		        "whitelist" : null
		      },
		      "selfService" : {
		        "selfServiceLinksEnabled" : true,
		        "signup" : "/create_account",
		        "passwd" : "/forgot_password"
		      }
		    },
		    "prompts" : [ {
		      "name" : "username",
		      "type" : "text",
		      "text" : "Email"
		    }, {
		      "name" : "password",
		      "type" : "password",
		      "text" : "Password"
		    }, {
		      "name" : "passcode",
		      "type" : "password",
		      "text" : "One Time Code (Get on at /passcode)"
		    } ],
		    "idpDiscoveryEnabled" : false
		  },
		  "name" : "The Twiglet Zone",
		  "version" : 0,
		  "created" : 1468364452298,
		  "last_modified" : 1468364452298
		}`)
	}))
	defer ts.Close()

	uaac, err := getUaac(ts.URL)
	if err != nil {
		t.Errorf("Failed to get uaa client; %v", err)
		return
	}

	var zone IdentityZone
	command := NewGetIdentityZoneByIDCommand(uaac, "abc123", &zone)

	if err := command.Execute(); err != nil {
		t.Errorf("Failed to get IdentityZones: %v", err)
	}

	if &zone == nil {
		t.Error("IdentityZone not found")
	}
}

func TestDeleteIdentityZoneCommand(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{}`)
	}))
	defer ts.Close()

	uaac, err := getUaac(ts.URL)
	if err != nil {
		t.Errorf("Failed to get uaa client; %v", err)
		return
	}

	zone := &IdentityZone{ID: "abc123"}
	command := NewDeleteIdentityZoneCommand(uaac, zone)

	if err := command.Execute(); err != nil {
		t.Errorf("Failed to delete IdentityZone: %v", err)
	}
}
