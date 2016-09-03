package identityzones

import uaa "github.com/pivotalservices/go-uaac"

type listIdentityZonesCommand struct {
	//TODO - figure out how to embed this type
	uaac  uaa.Client
	zones *[]*IdentityZone
}

type getIdentityZoneByIDCommand struct {
	uaac   uaa.Client
	zoneID string
	zone   *IdentityZone
}

func NewListIdentityZonesCommand(uaac uaa.Client, zones *[]*IdentityZone) uaa.Command {
	return &listIdentityZonesCommand{
		uaac:  uaac,
		zones: zones,
	}
}

func NewGetIdentityZoneByIDCommand(uaac uaa.Client, zoneID string, zone *IdentityZone) uaa.Command {
	return &getIdentityZoneByIDCommand{
		uaac:   uaac,
		zoneID: zoneID,
		zone:   zone,
	}
}

func (c *listIdentityZonesCommand) Execute() error {
	req := c.uaac.NewRequest("GET", "/identity-zones")
	if err := c.uaac.ExecuteAndUnmarshall(req, &c.zones); err != nil {
		return err
	}

	return nil
}

func (c *getIdentityZoneByIDCommand) Execute() error {
	req := c.uaac.NewRequest("GET", "/identity-zones/"+c.zoneID)
	if err := c.uaac.ExecuteAndUnmarshall(req, &c.zone); err != nil {
		return err
	}

	return nil
}
