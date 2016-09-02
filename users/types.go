package users

import "fmt"

type UserGuid string

func getUserGuid(i interface{}) (*UserGuid, error) {
	s, ok := i.(string)
	if ok != true {
		return nil, fmt.Errorf("%v is not a valid user guid", i)
	}

	guid := UserGuid(s)
	return &guid, nil
}

type Users struct {
	Users        []User `json:"resources"`
	StartIndex   int
	ItemsPerPage int
	TotalResults int
	Schemas      []string
}

type User struct {
	GUID       string      `json:"id,omitempty"`
	ExternalID string      `json:"externalId"`
	Username   string      `json:"userName"`
	Emails     []UserEmail `json:"emails"`
	Origin     string      `json:"origin"`
}

type UserEmail struct {
	Value   string `json:"value,omitempty"`
	Primary bool   `json:"primary"`
}
