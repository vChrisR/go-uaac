package clients

import uaa "github.com/pivotalservices/go-uaac"

type clientsCommand struct {
	uaac uaa.Client
}

type listClientsCommand struct {
	//TODO - figure out how to embed this type
	uaac    uaa.Client
	clients *OauthClients
}

func NewListClientsCommand(uaac uaa.Client, clients *OauthClients) uaa.Command {
	return &listClientsCommand{
		uaac:    uaac,
		clients: clients,
	}
}

func (c *listClientsCommand) Execute() error {
	req := c.uaac.NewRequest("GET", "/oauth/clients")

	if err := c.uaac.ExecuteAndUnmarshall(req, &c.clients); err != nil {
		return err
	}

	return nil
}
