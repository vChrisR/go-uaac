package serverinfo

import (
	"encoding/json"
	"testing"
)

func TestJSONUnmarshallServerInfo(t *testing.T) {
	responseBody := []byte(`{
      "app" : {
        "version" : "3.5.0-SNAPSHOT"
      },
      "links" : {
        "uaa" : "http://localhost:8080/uaa",
        "passwd" : "/forgot_password",
        "login" : "http://localhost:8080/uaa",
        "register" : "/create_account"
      },
      "zone_name" : "uaa",
      "entityID" : "cloudfoundry-saml-login",
      "commit_id" : "6681d65",
      "idpDefinitions" : {
        "SAML" : "http://localhost:8080/uaa/saml/discovery?returnIDParam=idp&entityID=cloudfoundry-saml-login&idp=SAML&isPassive=true"
      },
      "prompts" : {
        "username" : [ "text", "Email" ],
        "password" : [ "password", "Password" ]
      },
      "timestamp" : "2016-05-18T18:20:54+0000"
    }`)

	var info ServerInfo

	err := json.Unmarshal(responseBody, &info)
	if err != nil {
		t.Errorf("Failed to unmarshall json to ServerInfo: %v", err)
	}

	if len(info.Version()) == 0 {
		t.Error("Failed to unmarshall app.version to ServerInfo.Version")
	}
}
