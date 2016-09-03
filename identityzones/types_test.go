package identityzones

import (
	"encoding/json"
	"testing"
)

func TestJSONUnmarshallIdentityZone(t *testing.T) {
	responseBody := []byte(`{
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
	  "version" : 1,
		"description": "The Twiglet Zone",
	  "created" : 1468364452298,
	  "last_modified" : 1468364452298
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
