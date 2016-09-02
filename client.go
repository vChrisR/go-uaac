package uaa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

type Client interface {
	Oauth2Client
	ExecuteRequest(r *Request) ([]byte, error)
	ExecuteAndUnmarshall(r *Request, target interface{}) error
}

type uaaClient struct {
	Oauth2Client
}

func NewClient(config *ClientConfig) (Client, error) {
	client, err := NewOauth2Client(config)
	if err != nil {
		return nil, err
	}
	return &uaaClient{Oauth2Client: client}, nil
}

func (u *uaaClient) ExecuteRequest(r *Request) ([]byte, error) {
	resp, err := u.DoRequest(r)
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

func (u *uaaClient) ExecuteAndUnmarshall(r *Request, target interface{}) error {
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
