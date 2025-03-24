package irepositoryweb2

import entitiesweb2 "backend/internal/domain/entities/web2"

type IUserRepository interface {
	CreateUser(entitiesweb2.User) (entitiesweb2.User, error)
	DeleteUser(entitiesweb2.User) (entitiesweb2.User, error)
	UpdateUser(entitiesweb2.User) (entitiesweb2.User, error)

	CreateUserByManager(entitiesweb2.UserWithAuthority) (entitiesweb2.UserWithAuthority, error)
	DeleteUserByManager(entitiesweb2.UserWithAuthority) (entitiesweb2.UserWithAuthority, error)
	UpdateUserByManager(entitiesweb2.UserWithAuthority) (entitiesweb2.UserWithAuthority, error)
}

type IUserWithAuthorityRepository interface {
	CreateAuthorityByManager(entitiesweb2.UserWithAuthority) (entitiesweb2.UserWithAuthority, error)
	DeleteAuthorityByManager(entitiesweb2.UserWithAuthority) (entitiesweb2.UserWithAuthority, error)
	UpdateAuthorityByManager(entitiesweb2.UserWithAuthority) (entitiesweb2.UserWithAuthority, error)
}

type IUserWithFriendRepository interface {
	CreateFriendByManager(entitiesweb2.UserWithFriend) (entitiesweb2.UserWithFriend, error)
	DeleteFriendByManager(entitiesweb2.UserWithFriend) (entitiesweb2.UserWithFriend, error)
	UpdateFriendByManager(entitiesweb2.UserWithFriend) (entitiesweb2.UserWithFriend, error)
}
