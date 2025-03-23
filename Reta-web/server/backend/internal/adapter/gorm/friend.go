package model

import (
	entitiesweb2 "backend/internal/domain/entities/web2"
	"time"

	"gorm.io/gorm"
)

type Friend struct {
	FriendUUID     string `gorm:"column:friend_uuid;not null;unique;type:varchar(128)" json:"friend_uuid"`
	FriendUsername string `gorm:"column:friend_username;type:varchar(128)" json:"friend_username"`
	FriendFullname string `gorm:"column:friend_full_name;not null;type:varchar(100)" json:"friend_full_name"`
	FriendNickname string `gorm:"column:friend_nick_name;not null;type:varchar(100)" json:"friend_nick_name"`
	FriendStatus   string `gorm:"column:friend_status;not null;type:varchar(20)" json:"friend_status"`

	CreatedAt *time.Time     `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:datetime" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

func (Friend) TableName() string {
	return "friend"
}

func (f Friend) ToDomain() entitiesweb2.Friend {
	return entitiesweb2.Friend{
		FriendUUID:     f.FriendUUID,
		FriendUsername: f.FriendUsername,
		FriendFullname: f.FriendFullname,
		FriendNickname: f.FriendNickname,
		FriendStatus:   f.FriendStatus,
	}
}

func (f Friend) FromDomain(friend entitiesweb2.Friend) Friend {
	return Friend{
		FriendUUID:     friend.FriendUUID,
		FriendUsername: friend.FriendUsername,
		FriendFullname: friend.FriendFullname,
		FriendNickname: friend.FriendNickname,
		FriendStatus:   friend.FriendStatus,
	}
}
