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

func (uu *UserUsecase) UpdatePassword(wantUpdatePasswordUser entities.User, newPassword string) error {
	return uu.userRepo.UpdatePassword(wantUpdatePasswordUser, newPassword)
}

func (uu *UserUsecase) ListAuthority(wantListAuthorityUser entities.User) ([]entities.Authority, error) {
	return uu.userRepo.ListAuthority(wantListAuthorityUser)
}

func (uu *UserUsecase) ListFriend(wantListFriendUser entities.User) ([]entities.Friend, error) {
	return uu.userRepo.ListFiend(wantListFriendUser)
}

type UserAuthorityUsercase struct {
	UserUsecase
	userAuthRepo irepository.IUserWithAuthorityRepository
}

func NewUserAuthorityUsecase(userAuthRepo irepository.IUserWithAuthorityRepository) *UserAuthorityUsercase {
	return &UserAuthorityUsercase{userAuthRepo: userAuthRepo}
}

func (uau *UserAuthorityUsercase) CreateAuthorityByManager(wantAddAuthorityUser entities.User, addedAuthority entities.Authority) (entities.UserWithAuthority, error) {
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

func (uau *UserAuthorityUsercase) DeleteAuthorityByManager(wantDeleteAuthorityUser entities.User, deletedAuthority entities.Authority) (entities.UserWithAuthority, error) {
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

type UserFriendUsercase struct {
	UserUsecase
	userFriendRepo irepository.IUserWithFriendRepository
}

func NewUserFriendUsecase(userFriendRepo irepository.IUserWithFriendRepository) *UserFriendUsercase {
	return &UserFriendUsercase{userFriendRepo: userFriendRepo}
}

func (ufu *UserFriendUsercase) CreateFriendBySelf(wantAddFriendUser entities.User, addedFriend entities.Friend) (entities.UserWithFriend, error) {
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

func (ufu *UserFriendUsercase) DeleteFriendBySelf(wantDeleteFriendUser entities.User, deletedFriend entities.Friend) (entities.UserWithFriend, error) {
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

func (ufu *UserFriendUsercase) UpdateFriendInitiative(wantUpdateFriendUser entities.User, updatedFriend entities.Friend) (entities.UserWithFriend, error) {
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
