package groups

import uaa "github.com/pivotalservices/go-uaac"

type listGroupsCommand struct {
	uaac         uaa.Client
	listResponse *ListResponse
}

func NewListGroupsCommand(uaac uaa.Client, listGroupsResponse *ListResponse) uaa.Command {
	return &listGroupsCommand{
		uaac:         uaac,
		listResponse: listGroupsResponse,
	}
}

func (c *listGroupsCommand) Execute() error {
	req := c.uaac.NewRequest("GET", "/Groups")
	if err := c.uaac.ExecuteAndUnmarshall(req, &c.listResponse); err != nil {
		return err
	}

	return nil
}
