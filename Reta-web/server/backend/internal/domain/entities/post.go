package entities

import (
	"time"
)

type Post struct {
	UUID         string    `json:"uuid"`
	Author       string    `json:"author"`
	Likes        int       `json:"likes"`
	CommentUUID  []string  `json:"comment_uuid"`
	CommentCount int       `json:"comment_count"`
	CreatedAt    time.Time `json:"create_time"`
	UpdatedAt    time.Time `json:"update_time"`
}

type PostWithComment struct {
	Post
	Comments []Comment `json:"comments"`
}
