package usecase

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/irepository"
)

type PostUsecase struct {
	postRepo irepository.IPostRepository
}

func NewPostUsecase(postRepo irepository.IPostRepository) *PostUsecase {
	return &PostUsecase{postRepo: postRepo}
}

func (pu *PostUsecase) ListComment(wantListCommentPost entities.Post) ([]entities.Comment, error) {
	return pu.postRepo.ListComment(wantListCommentPost)
}

type PostCommentUsecase struct {
	PostUsecase
	postCommentRepo irepository.IPostWithCommentRepository
}

func NewPostCommentUsecase(postCommentRepo irepository.IPostWithCommentRepository) *PostCommentUsecase {
	return &PostCommentUsecase{postCommentRepo: postCommentRepo}
}

func (pcu *PostCommentUsecase) CreateComment(wantAddedCommentPost entities.Post, addedComment entities.Comment) (entities.PostWithComment, error) {
	comments, listerr := pcu.ListComment(wantAddedCommentPost)
	if listerr != nil {
		return entities.PostWithComment{}, listerr
	}
	postWithComment := entities.PostWithComment{
		Post:     wantAddedCommentPost,
		Comments: comments,
	}
	return pcu.postCommentRepo.CreateComment(postWithComment, addedComment)
}

func (pcu *PostCommentUsecase) DeleteComment(wantDeletedCommentPost entities.Post, deletedCommentUUID string) (entities.PostWithComment, error) {
	comments, listerr := pcu.ListComment(wantDeletedCommentPost)
	if listerr != nil {
		return entities.PostWithComment{}, listerr
	}
	postWithComment := entities.PostWithComment{
		Post:     wantDeletedCommentPost,
		Comments: comments,
	}
	return pcu.postCommentRepo.DeleteComment(postWithComment, deletedCommentUUID)
}

func (pcu *PostCommentUsecase) UpdateComment(wantUpdateCommentPost entities.Post, updatedCommentUUID string, newMsg string, newURL []string) (entities.PostWithComment, error) {
	comments, listerr := pcu.ListComment(wantUpdateCommentPost)
	if listerr != nil {
		return entities.PostWithComment{}, listerr
	}
	postWithComment := entities.PostWithComment{
		Post:     wantUpdateCommentPost,
		Comments: comments,
	}
	return pcu.postCommentRepo.UpdateComment(postWithComment, updatedCommentUUID, newMsg, newURL)
}
