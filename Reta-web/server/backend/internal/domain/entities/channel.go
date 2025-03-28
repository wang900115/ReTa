package entities

type Channel struct {
	UUID        string            `json:"uuid"`
	Host        string            `json:"Host"`
	Name        string            `json:"name"`
	Type        string            `json:"type"`
	Public      bool              `json:"public"`
	Members     []string          `json:"members"`
	MemberCount int               `json:"member_count"`
	Permissions map[string]string `json:"permissions"`
	Description string            `json:"description"`
}
