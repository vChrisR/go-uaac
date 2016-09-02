package identityzones

import uaa "github.com/pivotalservices/go-uaac"

type listIdentityZonesCommand struct {
	//TODO - figure out how to embed this type
	uaac  uaa.Client
	zones *[]*IdentityZone
}

func NewListIdentityZonesCommand(uaac uaa.Client, zones *[]*IdentityZone) uaa.Command {
	return &listIdentityZonesCommand{
		uaac:  uaac,
		zones: zones,
	}
}

func (c *listIdentityZonesCommand) Execute() error {
	req := c.uaac.NewRequest("GET", "/identity-zones")
	if err := c.uaac.ExecuteAndUnmarshall(req, &c.zones); err != nil {
		return err
	}

	return nil
}
