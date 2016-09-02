package serverinfo

type ServerInfo struct {
	App            map[string]string
	Links          map[string]string
	ZoneName       string `json:"zone_name"`
	EntityID       string
	CommitID       string `json:"commit_id"`
	IDPDefinitions map[string]string
	Prompts        map[string][]string
	Timestamp      string
}

func (s ServerInfo) Version() string {
	return s.App["version"]
}
