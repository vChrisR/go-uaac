package serverinfo

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

func TestGetServerInfoCommand(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
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
	}))
	defer ts.Close()

	uaac, err := getUaac(ts.URL)
	if err != nil {
		t.Errorf("Failed to get uaa client; %v", err)
		return
	}

	var serverInfo ServerInfo
	command := NewGetServerInfoCommand(uaac, &serverInfo)

	if err := command.Execute(); err != nil {
		t.Errorf("Failed to get ServerInfo: %v", err)
	}

	if len(serverInfo.Version()) == 0 {
		t.Error("ServerInfo.Version() was blank")
	}
}
