package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	uaa "github.com/pivotalservices/go-uaac"
)

type getUsersCommand struct {
	//TODO - figure out how to embed this type
	uaac  uaa.Client
	users *Users
}

type createUserCommand struct {
	//TODO - figure out how to embed this type
	uaac uaa.Client
	user *User
	guid *UserGuid
}

func NewGetUsersCommand(uaac uaa.Client, users *Users) uaa.Command {
	return &getUsersCommand{
		uaac:  uaac,
		users: users,
	}
}

func NewCreateUserCommand(uaac uaa.Client, user *User) uaa.Command {
	return &createUserCommand{
		uaac: uaac,
		user: user,
	}
}

func (c *getUsersCommand) Execute() error {
	req := c.uaac.NewRequest("GET", "/Users")
	if err := c.uaac.ExecuteAndUnmarshall(req, &c.users); err != nil {
		return fmt.Errorf("Failed to GET Users: %v", err)
	}

	return nil
}

func (c *createUserCommand) Execute() error {
	req := c.uaac.NewRequest("POST", "/Users")
	req.SetPayload(c.user)
	req.AddHeader("Accept", "application/json")
	req.AddHeader("Content-Type", "application/json")

	resp, err := c.uaac.DoRequest(req)
	if err != nil {
		return fmt.Errorf("Failed to execute request %v: %v\n", req, err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to read response body: %v\n", err)
	}

	var createUserResponse map[string]interface{}
	if err = json.Unmarshal(body, &createUserResponse); err != nil {
		return fmt.Errorf("Failed to process %v response: %v", req, err)
	}

	var id *UserGuid
	if resp.StatusCode == 409 {
		id, err = getUserGuid(createUserResponse["user_id"])
	} else {
		id, err = getUserGuid(createUserResponse["id"])
	}
	
	if err != nil {
	   return fmt.Errorf("Failed to find user id in createuser response: %v", err)
	}

	c.user.GUID = *id

	return err
}
