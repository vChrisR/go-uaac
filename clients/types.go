package clients

type OauthClients struct {
	Clients      []OauthClient `json:"resources"`
	StartIndex   int
	ItemsPerPage int
	TotalResults int
	Schemas      []string
}

type OauthClient struct {
	ID   string `json:"client_id"`
	Name string
	//AutoApprove          bool - TODO - this field can be a bool or an array??
	Action                 string
	Scope                  []string
	ResourceIDs            []string `json:"resource_ids"`
	Authorities            []string
	AuthorizedGrantTypes   []string `json:"authorized_grant_types"`
	LastModified           int
	RedirectURI            []string `json:"redirect_uri"`
	SignupRedirectURL      string   `json:"signup_redirect_url"`
	ChangeEmailRedirectURL string   `json:"change_email_redirect_url"`
}
