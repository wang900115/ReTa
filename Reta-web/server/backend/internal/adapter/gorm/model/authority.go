package model

import (
	entities "backend/internal/domain/entities"
	"time"

	"gorm.io/gorm"
)

type Authority struct {
	UUID        string `gorm:"column:uuid;primary_key;type:varchar(36)"`
	ID          string `gorm:"column:id;not null;unique;type:varchar(128)" json:"id"`
	Name        string `gorm:"column:name;not null;unique;type:varchar(128)" json:"name"`
	Description string `gorm:"column:description;type:text" json:"description"`

	CreatedAt *time.Time     `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:datetime" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

func (Authority) TableName() string {
	return "authority"
}

func (a Authority) ToDomain() entities.Authority {
	return entities.Authority{
		UUID:        a.UUID,
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
	}
}

func (a Authority) FromDomain(authority entities.Authority) Authority {
	return Authority{
		UUID:        authority.UUID,
		ID:          authority.ID,
		Name:        authority.Name,
		Description: authority.Description,
	}
}
