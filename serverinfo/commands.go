package serverinfo

import uaa "github.com/pivotalservices/go-uaac"

type getServerInfoCommand struct {
	//TODO - figure out how to embed this type
	uaac uaa.Client
	info *ServerInfo
}

func NewGetServerInfoCommand(uaac uaa.Client, info *ServerInfo) uaa.Command {
	return &getServerInfoCommand{
		uaac: uaac,
		info: info,
	}
}

func (c *getServerInfoCommand) Execute() error {
	req := c.uaac.NewRequest("GET", "/info")
	if err := c.uaac.ExecuteAndUnmarshall(req, &c.info); err != nil {
		return err
	}

	return nil
}
