package repository

import (
	"backend/internal/adapter/gorm/model"
	"backend/internal/domain/entities"
	"backend/internal/domain/irepository"

	"gorm.io/gorm"
)

type PostRepository struct {
	gorm *gorm.DB
}

func NewPostRepository(gorm *gorm.DB) irepository.IPostRepository {
	return &PostRepository{gorm: gorm}
}

func (p *PostRepository) ListComment(post entities.Post) ([]entities.Comment, error) {
	var commentsModel []model.Comment
	if err := p.gorm.Table("post_comment").Joins("JOIN comment ON comment.UUID = post_comment.comment_uuid").Where("post_comment.post_uuid = ?", post.UUID).Find(&commentsModel).Error; err != nil {
		return []entities.Comment{}, err
	}
	var comments []entities.Comment
	for _, commentModel := range commentsModel {
		comments = append(comments, commentModel.ToDomain())
	}
	return comments, nil
}

type PostCommentRepository struct {
	gorm *gorm.DB
}

func NewPostCommentRepository(gorm *gorm.DB) irepository.IPostWithCommentRepository {
	return &PostCommentRepository{gorm: gorm}
}

func (pc *PostCommentRepository) CreateComment(postwithComment entities.PostWithComment, comment entities.Comment) (entities.PostWithComment, error) {
	var postModel model.Post
	if err := pc.gorm.First(&postModel, "uuid = ?", postwithComment.UUID).Error; err != nil {
		return entities.PostWithComment{}, err
	}

	commentModel := model.Comment{}.FromDomain(comment)
	if err := pc.gorm.Create(&commentModel).Error; err != nil {
		return entities.PostWithComment{}, err
	}

	tx := pc.gorm.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&postModel).Association("Comments").Append(&commentModel); err != nil {
		tx.Rollback()
		return entities.PostWithComment{}, err
	}

	var commentsModel []model.Comment
	if err := tx.Table("post_comment").Joins("JOIN comment ON comment.uuid = post_comment.comment_uuid").Where("post_comment.post_uuid", postwithComment.UUID).Find(&commentsModel).Error; err != nil {
		return entities.PostWithComment{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return entities.PostWithComment{}, err
	}

	postComment := model.PostWithComment{
		Post:     postModel,
		Comments: commentsModel,
	}

	return postComment.ToDomain(), nil

}

func (pc *PostCommentRepository) DeleteComment(postwithComment entities.PostWithComment, uuid string) (entities.PostWithComment, error) {
	var postModel model.Post
	if err := pc.gorm.First(&postModel, "uuid = ?", postwithComment.UUID).Error; err != nil {
		return entities.PostWithComment{}, err
	}

	var commentModel model.Comment
	if err := pc.gorm.First(&commentModel, "uuid = ?", uuid).Error; err != nil {
		return entities.PostWithComment{}, err
	}

	tx := pc.gorm.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&postModel).Association("Comments").Delete(&commentModel); err != nil {
		tx.Rollback()
		return entities.PostWithComment{}, err
	}

	var commentsModel []model.Comment
	if err := tx.Table("post_comment").Joins("JOIN comment ON comment.uuid = post_comment.comment_uuid").Where("post_comment.post_uuid", postwithComment.UUID).Find(&commentsModel).Error; err != nil {
		return entities.PostWithComment{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return entities.PostWithComment{}, err
	}

	postComment := model.PostWithComment{
		Post:     postModel,
		Comments: commentsModel,
	}

	return postComment.ToDomain(), nil
}

func (pc *PostCommentRepository) UpdateComment(postwithComment entities.PostWithComment, uuid string, msg string, url []string) (entities.PostWithComment, error) {
	var postModel model.Post
	if err := pc.gorm.First(&postModel, "uuid = ?", postwithComment.UUID).Error; err != nil {
		return entities.PostWithComment{}, err
	}

	var commentModel model.Comment
	if err := pc.gorm.First(&commentModel, "uuid = ?", uuid).Error; err != nil {
		return entities.PostWithComment{}, err
	}

	tx := pc.gorm.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	commentModel.Message = msg
	commentModel.MediaURL = url

	if err := tx.Save(&commentModel).Error; err != nil {
		tx.Rollback()
		return entities.PostWithComment{}, err
	}

	var commentsModel []model.Comment
	if err := tx.Table("post_comment").Joins("JOIN comment ON comment.uuid = post_comment.comment_uuid").Where("post_comment.post_uuid", postwithComment.UUID).Find(&commentsModel).Error; err != nil {
		tx.Rollback()
		return entities.PostWithComment{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return entities.PostWithComment{}, err
	}

	postComment := model.PostWithComment{
		Post:     postModel,
		Comments: commentsModel,
	}

	return postComment.ToDomain(), nil
}
