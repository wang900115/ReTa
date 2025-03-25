package irepositoryweb2

import entitiesweb2 "backend/internal/domain/entities/web2"

type IUserRepository interface {
	CreateUser(entitiesweb2.User) (entitiesweb2.User, error)
	DeleteUser(entitiesweb2.User) (entitiesweb2.User, error)
	UpdateUser(entitiesweb2.User) (entitiesweb2.User, error)
}

type IUserWithAuthorityRepository interface {
	CreateAuthorityByManager(entitiesweb2.UserWithAuthority, entitiesweb2.Authority) (entitiesweb2.UserWithAuthority, error)
	DeleteAuthorityByManager(entitiesweb2.UserWithAuthority, entitiesweb2.Authority) (entitiesweb2.UserWithAuthority, error)
}

type IUserWithFriendRepository interface {
	CreateFriendBySelf(entitiesweb2.UserWithFriend, entitiesweb2.Friend) (entitiesweb2.UserWithFriend, error)
	DeleteFriendBySelf(entitiesweb2.UserWithFriend, entitiesweb2.Friend) (entitiesweb2.UserWithFriend, error)
	UpdateFriendBySelf(entitiesweb2.UserWithFriend, entitiesweb2.Friend) (entitiesweb2.UserWithFriend, error)
}
