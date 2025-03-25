package entitiesweb2

type Friend struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname"`
	Status   string `json:"status"`
	Banned   bool   `json:"banned"`
}
