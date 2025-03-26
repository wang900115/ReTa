package irepository

import entities "backend/internal/domain/entities"

type IUserRepository interface {
	CreateUser(entities.User) (entities.User, error)
	DeleteUser(entities.User) (entities.User, error)
	UpdateUser(entities.User) (entities.User, error)
	ListAuthority(entities.User) ([]entities.Authority, error)
	ListFriend(entities.User) ([]entities.Friend, error)
	UpdatePassword(entities.User, string) error
}

type IUserWithAuthorityRepository interface {
	CreateAuthorityByManager(entities.UserWithAuthority, entities.Authority) (entities.UserWithAuthority, error)
	DeleteAuthorityByManager(entities.UserWithAuthority, entities.Authority) (entities.UserWithAuthority, error)
}

type IUserWithFriendRepository interface {
	CreateFriendBySelf(entities.UserWithFriend, entities.Friend) (entities.UserWithFriend, error)
	DeleteFriendBySelf(entities.UserWithFriend, entities.Friend) (entities.UserWithFriend, error)
	UpdateFriendInitiative(entities.UserWithFriend, entities.Friend) (entities.UserWithFriend, error)
}
