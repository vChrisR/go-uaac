package identityzones

type IdentityZone struct {
	ID           string
	Subdomain    string
	Name         string
	Version      int
	Description  string
	Created      int
	LastModified int `json:"last_modified"`
}
