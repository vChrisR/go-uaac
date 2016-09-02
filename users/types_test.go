package users

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestJSONUnmarshallUsersWithUaa20Model(t *testing.T) {
	responseBody, err := ioutil.ReadFile("./testdata/2.0.0/GET.Users.json")

	if err != nil {
		panic("Failed to read ./testdata/2.0.0/GET.Users.json: " + err.Error())
	}

	var users Users

	err = json.Unmarshal(responseBody, &users)
	if err != nil {
		t.Errorf("Failed to unmarshall json to Users: %v", err)
	}

	if len(users.Users) == 0 {
		t.Error("Failed to unmarshall resources json field onto Users.Users")
	}

	user := users.Users[0]

	if len(user.GUID) == 0 {
		t.Error("Failed to unmarshall resources[0].id field onto Users.Users[0].GUID")
	}
	if len(user.Username) == 0 {
		t.Error("Failed to unmarshall resources[0].username field onto Users.Users[0].Username")
	}
}

func readTestdata(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
