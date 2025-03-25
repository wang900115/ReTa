package entitiesweb2

type TokenClaims struct {
	UserUUID string `json:"user_uuid"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname"`
}
