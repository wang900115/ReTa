package model

import (
	"backend/internal/domain/entities"
	"time"
)

type Post struct {
	UUID         string   `json:"uuid"`
	Author       string   `json:"author"`
	Content      string   `json:"content"`
	MediaURL     []string `json:"media_urls"`
	Likes        int      `json:"likes"`
	CommentUUID  []string `json:"comment_uuid"`
	CommentCount int      `json:"comment_count"`

	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func (Post) TableName() string {
	return "post"
}

func (p Post) ToDomain() entities.Post {
	return entities.Post{
		UUID:         p.UUID,
		Author:       p.Author,
		Content:      p.Content,
		MediaURL:     p.MediaURL,
		Likes:        p.Likes,
		CommentUUID:  p.CommentUUID,
		CommentCount: p.CommentCount,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func (p Post) FromDomain(post entities.Post) Post {
	return Post{
		UUID:         post.UUID,
		Author:       post.Author,
		Content:      post.Content,
		MediaURL:     post.MediaURL,
		Likes:        post.Likes,
		CommentUUID:  post.CommentUUID,
		CommentCount: post.CommentCount,
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
	}
}

type PostWithComment struct {
	Post
	Comments []Comment `gorm:"many2many: post_comment; foreignKey:UUID;joinForeignKey:post_uuid;Reference:UUID;JoinReference:comment_uuid" json:"comments"`
}

func (PostWithComment) TableName() string {
	return "post_comment"
}

func (pc PostWithComment) ToDomain() entities.PostWithComment {
	comments := make([]entities.Comment, len(pc.Comments))
	for i, a := range pc.Comments {
		comments[i] = a.ToDomain()
	}
	return entities.PostWithComment{
		Post:     pc.Post.ToDomain(),
		Comments: comments,
	}
}

func (pc PostWithComment) FromDomain(postwithcomment entities.PostWithComment) PostWithComment {
	comments := make([]Comment, len(postwithcomment.Comments))
	for i, a := range postwithcomment.Comments {
		comments[i] = Comment{}.FromDomain(a)
	}
	return PostWithComment{
		Post:     Post{}.FromDomain(postwithcomment.Post),
		Comments: comments,
	}
}
