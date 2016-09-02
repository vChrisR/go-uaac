package uaa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/dave-malone/oauth"
)

type Client interface {
	NewRequest(method, path string) *oauth.Request
	ExecuteRequest(r *oauth.Request) ([]byte, error)
	ExecuteAndUnmarshall(r *oauth.Request, target interface{}) error

	// GetServerInfo() (ServerInfo, error)
	// ListOauthClients() (OauthClients, error)
	// ListIdentityZones() ([]IdentityZone, error)
	// ListUsers() (Users, error)
	// CreateUser(user *User) (*UserGuid, error)
}

type uaaClient struct {
	oauthClient *oauth.Client
}

func NewClient(config *oauth.ClientConfig) (Client, error) {
	client, err := oauth.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &uaaClient{client}, nil
}

func (u *uaaClient) NewRequest(method, path string) *oauth.Request {
	return u.oauthClient.NewRequest(method, path)
}

func (u *uaaClient) ExecuteRequest(r *oauth.Request) ([]byte, error) {
	resp, err := u.oauthClient.DoRequest(r)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute request %v: %v\n", r, err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %v\n", err)
	}

	return body, nil
}

func (u *uaaClient) ExecuteAndUnmarshall(r *oauth.Request, target interface{}) error {
	body, err := u.ExecuteRequest(r)
	if err != nil {
		return fmt.Errorf("Failed to execute request %v: %v", r, err)
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		return fmt.Errorf("Unable to unmarshall response body to type %s; error: %v; response body: %s", reflect.TypeOf(target), err.Error(), string(body))
	}

	return nil
}

/*
func (c *uaaClient) GetServerInfo() (ServerInfo, error) {
	var info ServerInfo
	req := c.oauthClient.NewRequest("GET", "/info")
	req.AddHeader("Accept", "application/json")

	if err := c.executeAndUnmarshall(req, &info); err != nil {
		return info, err
	}

	return info, nil
}

func (c *uaaClient) ListOauthClients() (OauthClients, error) {
	var clients OauthClients
	req := c.oauthClient.NewRequest("GET", "/oauth/clients")

	if err := c.executeAndUnmarshall(req, &clients); err != nil {
		return clients, err
	}

	return clients, nil
}

func (c *uaaClient) ListIdentityZones() ([]IdentityZone, error) {
	var zones []IdentityZone
	req := c.oauthClient.NewRequest("GET", "/identity-zones")
	if err := c.executeAndUnmarshall(req, &zones); err != nil {
		return zones, err
	}

	return zones, nil
}

func (c *uaaClient) ListUsers() (Users, error) {
	var users Users

	req := c.oauthClient.NewRequest("GET", "/Users")
	req.AddHeader("Accept", "application/json")

	if err := c.executeAndUnmarshall(req, &users); err != nil {
		return users, err
	}

	return users, nil
}

func (u *uaaClient) CreateUser(user *User) (*UserGuid, error) {
	req := u.oauthClient.NewRequest("POST", "/Users")
	req.SetPayload(user)
	req.AddHeader("Accept", "application/json")
	req.AddHeader("Content-Type", "application/json")

	resp, err := u.oauthClient.DoRequest(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute request %v: %v\n", req, err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %v\n", err)
	}

	var createUserResponse map[string]interface{}
	err = json.Unmarshal(body, &createUserResponse)
	if err != nil {
		return nil, fmt.Errorf("Failed to process %v response: %v", req, err)
	}

	if resp.StatusCode == 409 {
		return getUserGuid(createUserResponse["user_id"])
	}

	return getUserGuid(createUserResponse["id"])
}
*/
