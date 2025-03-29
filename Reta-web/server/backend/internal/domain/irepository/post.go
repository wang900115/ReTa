package irepository

import entities "backend/internal/domain/entities"

type IPostRepository interface {
	// 列出貼文留言
	ListComment(entities.Post) ([]entities.Comment, error)
}

type IPostWithCommentRepository interface {
	// 新增該貼文的留言
	CreateComment(entities.PostWithComment, entities.Comment) (entities.PostWithComment, error)
	// 刪除該貼文的留言
	DeleteComment(entities.PostWithComment, string) (entities.PostWithComment, error)
	// 更新該貼文的留言
	UpdateComment(entities.PostWithComment, string, string, []string) (entities.PostWithComment, error)
}
