package entitiesweb2

type Friend struct {
	FriendUUID     string `json:"friend_uuid"`
	FriendUsername string `json:"friend_username"`
	FriendFullname string `json:"friend_fullname"`
	FriendNickname string `json:"friend_nickname"`
	FriendStatus   string `json:"friend_status"`
	FriendBanned   bool   `json:"friend_banned"`
}
