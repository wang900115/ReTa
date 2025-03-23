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
