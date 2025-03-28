package model

import (
	"backend/internal/domain/entities"
	"time"
)

type Comment struct {
	UUID    string    `gorm:"column:uuid;not null" json:"uuid"`
	Author  string    `gorm:"column:author;not null;type:varchar(128)" json:"author"`
	Message string    `gorm:"column:message;not null;type:varchar(256)" json:"message"`
	Time    time.Time `gorm:"column:time;type:datetime" json:"time"`
}

func (Comment) TableName() string {
	return "comment"
}

func (c Comment) ToDomain() entities.Comment {
	return entities.Comment{
		UUID:    c.UUID,
		Author:  c.Author,
		Message: c.Message,
		Time:    c.Time,
	}
}

func (c Comment) FromDomain(comment entities.Comment) Comment {
	return Comment{
		UUID:    comment.UUID,
		Author:  comment.Author,
		Message: comment.Message,
		Time:    comment.Time,
	}
}
