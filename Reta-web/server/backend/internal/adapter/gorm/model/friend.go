package model

import (
	entities "backend/internal/domain/entities"
	"time"

	"gorm.io/gorm"
)

// 朋友
type Friend struct {
	UUID     string `gorm:"column:uuid;not null;unique;type:varchar(128)" json:"uuid"`
	Username string `gorm:"column:username;type:varchar(128)" json:"username"`
	Fullname string `gorm:"column:full_name;not null;type:varchar(100)" json:"full_name"`
	Nickname string `gorm:"column:nick_name;not null;type:varchar(100)" json:"nick_name"`
	Status   string `gorm:"column:status;not null;type:varchar(20)" json:"status"`
	Banned   bool   `gorm:"column:banned;not null;default:false;type:tinyint(1)" json:"banned"`

	CreatedAt *time.Time     `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:datetime" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

func (Friend) TableName() string {
	return "friend"
}

func (f Friend) ToDomain() entities.Friend {
	return entities.Friend{
		UUID:     f.UUID,
		Username: f.Username,
		Fullname: f.Fullname,
		Nickname: f.Nickname,
		Status:   f.Status,
		Banned:   f.Banned,
	}
}

func (f Friend) FromDomain(friend entities.Friend) Friend {
	return Friend{
		UUID:     friend.UUID,
		Username: friend.Username,
		Fullname: friend.Fullname,
		Nickname: friend.Nickname,
		Status:   friend.Status,
		Banned:   friend.Banned,
	}
}
