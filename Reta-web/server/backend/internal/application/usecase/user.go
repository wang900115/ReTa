package usercase

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/irepository"
)

type UserUsecase struct {
	userRepo irepository.IUserRepository
}

func NewUserUsecase(userRepo irepository.IUserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (uu *UserUsecase) CreateUser(wantCreateUser entities.User) (entities.User, error) {
	return uu.userRepo.CreateUser(wantCreateUser)
}

func (uu *UserUsecase) DeleteUser(wantDeleteUser entities.User) (entities.User, error) {
	return uu.userRepo.DeleteUser(wantDeleteUser)
}

func (uu *UserUsecase) UpdateUser(wantUpdateUser entities.User) (entities.User, error) {
	return uu.userRepo.UpdateUser(wantUpdateUser)
}

func (uu *UserUsecase) CreateComment(wantCreateCommentUser entities.User, beAddedPost entities.Post, addedComment entities.Comment) (entities.PostWithComment, error) {
	return uu.userRepo.CreateComment(wantCreateCommentUser, beAddedPost, addedComment)
}

func (uu *UserUsecase) DeleteComment(wantDeleteCommentUser entities.User, beDeletedPost entities.Post, UUID string) (entities.PostWithComment, error) {
	return uu.userRepo.DeleteComment(wantDeleteCommentUser, beDeletedPost, UUID)
}

func (uu *UserUsecase) UpdateComment(wantUpdateCommentUser entities.User, beUpdatedPost entities.Post, UUID string, MSG string) (entities.PostWithComment, error) {
	return uu.userRepo.UpdateComment(wantUpdateCommentUser, beUpdatedPost, UUID, MSG)
}

func (uu *UserUsecase) UpdatePassword(wantUpdatePasswordUser entities.User, newPassword string) error {
	return uu.userRepo.UpdatePassword(wantUpdatePasswordUser, newPassword)
}

func (uu *UserUsecase) ListAuthority(wantListAuthorityUser entities.User) ([]entities.Authority, error) {
	return uu.userRepo.ListAuthority(wantListAuthorityUser)
}

func (uu *UserUsecase) ListFriend(wantListFriendUser entities.User) ([]entities.Friend, error) {
	return uu.userRepo.ListFriend(wantListFriendUser)
}

func (uu *UserUsecase) ListChannel(wantListChannelUser entities.User) ([]entities.Channel, error) {
	return uu.userRepo.ListChannel(wantListChannelUser)
}

func (uu *UserUsecase) ListPost(wantListPostUser entities.User) ([]entities.Post, error) {
	return uu.userRepo.ListPost(wantListPostUser)
}

type UserAuthorityUsecase struct {
	UserUsecase
	userAuthRepo irepository.IUserWithAuthorityRepository
}

func NewUserAuthorityUsecase(userAuthRepo irepository.IUserWithAuthorityRepository) *UserAuthorityUsecase {
	return &UserAuthorityUsecase{userAuthRepo: userAuthRepo}
}

func (uau *UserAuthorityUsecase) CreateAuthorityByManager(wantAddAuthorityUser entities.User, addedAuthority entities.Authority) (entities.UserWithAuthority, error) {
	authorities, listerr := uau.ListAuthority(wantAddAuthorityUser)
	if listerr != nil {
		return entities.UserWithAuthority{}, listerr
	}
	userWithAuthority := entities.UserWithAuthority{
		User:        wantAddAuthorityUser,
		Authorities: authorities,
	}
	return uau.userAuthRepo.CreateAuthorityByManager(userWithAuthority, addedAuthority)
}

func (uau *UserAuthorityUsecase) DeleteAuthorityByManager(wantDeleteAuthorityUser entities.User, deletedAuthority entities.Authority) (entities.UserWithAuthority, error) {
	authorities, listerr := uau.ListAuthority(wantDeleteAuthorityUser)
	if listerr != nil {
		return entities.UserWithAuthority{}, listerr
	}
	userWithAuthority := entities.UserWithAuthority{
		User:        wantDeleteAuthorityUser,
		Authorities: authorities,
	}
	return uau.userAuthRepo.DeleteAuthorityByManager(userWithAuthority, deletedAuthority)
}

type UserFriendUsecase struct {
	UserUsecase
	userFriendRepo irepository.IUserWithFriendRepository
}

func NewUserFriendUsecase(userFriendRepo irepository.IUserWithFriendRepository) *UserFriendUsecase {
	return &UserFriendUsecase{userFriendRepo: userFriendRepo}
}

func (ufu *UserFriendUsecase) CreateFriendBySelf(wantAddFriendUser entities.User, addedFriend entities.Friend) (entities.UserWithFriend, error) {
	friends, listerr := ufu.ListFriend(wantAddFriendUser)
	if listerr != nil {
		return entities.UserWithFriend{}, listerr
	}
	userWithFriend := entities.UserWithFriend{
		User:    wantAddFriendUser,
		Friends: friends,
	}
	return ufu.userFriendRepo.CreateFriendBySelf(userWithFriend, addedFriend)
}

func (ufu *UserFriendUsecase) DeleteFriendBySelf(wantDeleteFriendUser entities.User, deletedFriend entities.Friend) (entities.UserWithFriend, error) {
	friends, listerr := ufu.ListFriend(wantDeleteFriendUser)
	if listerr != nil {
		return entities.UserWithFriend{}, listerr
	}
	userWithFriend := entities.UserWithFriend{
		User:    wantDeleteFriendUser,
		Friends: friends,
	}
	return ufu.userFriendRepo.DeleteFriendBySelf(userWithFriend, deletedFriend)
}

func (ufu *UserFriendUsecase) UpdateFriendInitiative(wantUpdateFriendUser entities.User, updatedFriend entities.Friend) (entities.UserWithFriend, error) {
	friends, listerr := ufu.ListFriend(wantUpdateFriendUser)
	if listerr != nil {
		return entities.UserWithFriend{}, listerr
	}
	userWithFriend := entities.UserWithFriend{
		User:    wantUpdateFriendUser,
		Friends: friends,
	}
	return ufu.userFriendRepo.UpdateFriendInitiative(userWithFriend, updatedFriend)
}

type UserChannelUsecase struct {
	UserUsecase
	userChannelRepo irepository.IUserWithChannelRepository
}

func NewUserChannelUsecase(userChannelRepo irepository.IUserWithChannelRepository) *UserChannelUsecase {
	return &UserChannelUsecase{userChannelRepo: userChannelRepo}
}

func (ucu *UserChannelUsecase) CreateChannel(wantAddChannelUser entities.User, addedChannel entities.Channel) (entities.UserWithChannel, error) {
	channels, listerr := ucu.ListChannel(wantAddChannelUser)
	if listerr != nil {
		return entities.UserWithChannel{}, listerr
	}
	userWithChannel := entities.UserWithChannel{
		User:     wantAddChannelUser,
		Channels: channels,
	}
	return ucu.userChannelRepo.CreateChannel(userWithChannel, addedChannel)
}

func (ucu *UserChannelUsecase) DeleteChannel(wantDeleteChannelUser entities.User, deletedChannel entities.Channel) (entities.UserWithChannel, error) {
	channels, listerr := ucu.ListChannel(wantDeleteChannelUser)
	if listerr != nil {
		return entities.UserWithChannel{}, listerr
	}
	userWithChannel := entities.UserWithChannel{
		User:     wantDeleteChannelUser,
		Channels: channels,
	}
	return ucu.userChannelRepo.DeleteChannel(userWithChannel, deletedChannel)
}

func (ucu *UserChannelUsecase) JoinPublicChannel(wantJoinChannelUser entities.User, joinChannelUUID string) (entities.UserWithChannel, error) {
	channels, listerr := ucu.ListChannel(wantJoinChannelUser)
	if listerr != nil {
		return entities.UserWithChannel{}, listerr
	}
	userWithChannel := entities.UserWithChannel{
		User:     wantJoinChannelUser,
		Channels: channels,
	}
	return ucu.userChannelRepo.JoinPublicChannel(userWithChannel, joinChannelUUID)
}

func (ucu *UserChannelUsecase) QuitChannel(wantAddChannelUser entities.User, quitChannelUUID string) (entities.UserWithChannel, error) {
	channels, listerr := ucu.ListChannel(wantAddChannelUser)
	if listerr != nil {
		return entities.UserWithChannel{}, listerr
	}
	userWithChannel := entities.UserWithChannel{
		User:     wantAddChannelUser,
		Channels: channels,
	}
	return ucu.userChannelRepo.QuitChannel(userWithChannel, quitChannelUUID)
}

type UserPostUsecase struct {
	UserUsecase
	userPostRepo irepository.IUserWithPostRepository
}

func NewUserPostUsecase(userPostRepo irepository.IUserWithPostRepository) *UserPostUsecase {
	return &UserPostUsecase{userPostRepo: userPostRepo}
}

func (upu *UserPostUsecase) CreatePost(wantAddPostUser entities.User, addedPost entities.Post) (entities.UserWithPost, error) {
	posts, listerr := upu.ListPost(wantAddPostUser)
	if listerr != nil {
		return entities.UserWithPost{}, listerr
	}
	userWithPost := entities.UserWithPost{
		User:  wantAddPostUser,
		Posts: posts,
	}
	return upu.userPostRepo.CreatePost(userWithPost, addedPost)
}

func (upu *UserPostUsecase) DeletePost(wantDeletePostUser entities.User, deletedPost entities.Post) (entities.UserWithPost, error) {
	posts, listerr := upu.ListPost(wantDeletePostUser)
	if listerr != nil {
		return entities.UserWithPost{}, listerr
	}
	userWithPost := entities.UserWithPost{
		User:  wantDeletePostUser,
		Posts: posts,
	}
	return upu.userPostRepo.DeletePost(userWithPost, deletedPost)
}

func (upu *UserPostUsecase) UpdatePost(wantUpdatePostUser entities.User, updatedPost entities.Post) (entities.UserWithPost, error) {
	posts, listerr := upu.ListPost(wantUpdatePostUser)
	if listerr != nil {
		return entities.UserWithPost{}, listerr
	}
	userWithPost := entities.UserWithPost{
		User:  wantUpdatePostUser,
		Posts: posts,
	}
	return upu.userPostRepo.CreatePost(userWithPost, updatedPost)
}
