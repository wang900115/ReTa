package entitiesweb2

type User struct {
	UUID     string `json:"uuid"`
	IsEnable bool   `json:"is_enable"`
	Username string `json:"username"`
	Password string `json:"-"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`

	Status      string `json:"status"`
	Description string `json:"description"`
}

type UserWithAuthority struct {
	User
	Authorities []Authority `json:"authorities"`
}

type UserWithFriend struct {
	User
	Friends []Friend `json:"friends"`
	Total   int      `json:"total"`
}
