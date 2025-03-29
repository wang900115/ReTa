package model

import (
	"backend/internal/domain/entities"
	"time"
)

type Comment struct {
	UUID     string    `gorm:"column:uuid;not null" json:"uuid"`
	PostID   string    `gorm:"column:post_id;not null" json:"post_id"`
	Author   string    `gorm:"column:author;not null;type:varchar(128)" json:"author"`
	Message  string    `gorm:"column:message;not null;type:varchar(256)" json:"message"`
	MediaURL []string  `gorm:"column:media_url;type;JSON" json:"media_urls"`
	Time     time.Time `gorm:"column:time;type:datetime" json:"time"`
}

func (Comment) TableName() string {
	return "comment"
}

func (c Comment) ToDomain() entities.Comment {
	return entities.Comment{
		UUID:     c.UUID,
		PostID:   c.PostID,
		Author:   c.Author,
		Message:  c.Message,
		MediaURL: c.MediaURL,
		Time:     c.Time,
	}
}

func (c Comment) FromDomain(comment entities.Comment) Comment {
	return Comment{
		UUID:     comment.UUID,
		PostID:   comment.PostID,
		Author:   comment.Author,
		Message:  comment.Message,
		MediaURL: comment.MediaURL,
		Time:     comment.Time,
	}
}
