package groups

//PagedResponse Paginated api response values
type PagedResponse struct {
	TotalResults int `json:"totalresults"`
	ItemsPerPage int `json:"itemsperpage"`
	StartIndex   int `json:"startindex"`
}

//ListResponse mapping to the UAA GET /Groups call
type ListResponse struct {
	*PagedResponse
	Groups []*Group `json:"resources"`
}

//Metadata for each Group
type Metadata struct {
	Version int
	//example: 2016-07-12T23:00:54.467Z
	Created string
	//example: 2016-07-12T23:00:54.504Z
	LastModified string
}

//Member represents a member of a Group
type Member struct {
	//The alias of the identity provider that authenticated this user. "uaa" is an internal UAA user.
	Origin string
	//Either "USER" or "GROUP"
	Type string
	//Globally unique identifier of the member, either a user ID or another group ID
	Value string
}

//Group represents a group in UAA
type Group struct {
	//The identifier specified upon creation of the group, unique within the identity zone
	ID string
	//Human readable description of the group, displayed e.g. when approving scopes
	DisplayName string
	//Identifier for the identity zone to which the group belongs
	ZoneID      string
	Description string
	Meta        *Metadata
	Members     []*Member
}
