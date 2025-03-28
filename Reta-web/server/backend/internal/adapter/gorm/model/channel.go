package model

import (
	"backend/internal/domain/entities"
	"time"

	"gorm.io/gorm"
)

type Channel struct {
	UUID        string            `gorm:"column:uuid;primary_key;type:varchar(36)"`
	Host        string            `gorm:"column:host;not null;type:varchar(128)" json:"host"`
	Name        string            `gorm:"column:name;not null;unique;type:varchar(128)" json:"name"`
	Type        string            `gorm:"column:type;not null;type:varchar(36)" json:"type"`
	Public      bool              `gorm:"column:public;not null;default:true;type:tinyint(1)" json:"public"`
	Members     []string          `gorm:"column:members;type:JSON" json:"members"`
	MemberCount int               `gorm:"column:member_count;not null;default:1 type:int" json:"member_count"`
	Permissions map[string]string `gorm:"column:permissions;type:JSON" json:"permissions"`
	Description string            `gorm:"column:description; type:text" json:"description"`

	CreatedAt *time.Time     `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:datetime" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

func (Channel) TableName() string {
	return "channel"
}

func (c Channel) ToDomain() entities.Channel {
	return entities.Channel{
		UUID:        c.UUID,
		Host:        c.Host,
		Name:        c.Name,
		Type:        c.Type,
		Public:      c.Public,
		Members:     c.Members,
		MemberCount: c.MemberCount,
		Permissions: c.Permissions,
		Description: c.Description,
	}
}

func (c Channel) FromDomain(channel entities.Channel) Channel {
	return Channel{
		UUID:        channel.UUID,
		Host:        channel.Host,
		Name:        channel.Name,
		Type:        channel.Type,
		Public:      channel.Public,
		Members:     channel.Members,
		MemberCount: channel.MemberCount,
		Permissions: channel.Permissions,
		Description: channel.Description,
	}
}
