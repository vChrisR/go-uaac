package identityzones

import (
	"fmt"

	uaa "github.com/pivotalservices/go-uaac"
)

type listIdentityZonesCommand struct {
	uaac  uaa.Client
	zones *[]*IdentityZone
}

type getIdentityZoneByIDCommand struct {
	uaac   uaa.Client
	zoneID string
	zone   *IdentityZone
}

type deleteIdentityZoneCommand struct {
	uaac uaa.Client
	zone *IdentityZone
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

func NewDeleteIdentityZoneCommand(uaac uaa.Client, zone *IdentityZone) uaa.Command {
	return &deleteIdentityZoneCommand{
		uaac: uaac,
		zone: zone,
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

func (c *deleteIdentityZoneCommand) Execute() error {
	req := c.uaac.NewRequest("DELETE", "/identity-zones/"+c.zone.ID)
	resp, err := c.uaac.DoRequest(req)
	if err != nil {
		return fmt.Errorf("Failed to Delete identity zone with id %s. HTTP Response Code: %d; error: %v", c.zone.ID, resp.StatusCode, err)
	}

	return nil
}
