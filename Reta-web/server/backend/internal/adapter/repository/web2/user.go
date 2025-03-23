package repositoryweb2

import (
	entitiesweb2 "backend/internal/domain/entities/web2"
	irepositoryweb2 "backend/internal/domain/irepository/web2"

	"gorm.io/gorm"
)

type UserRepository struct {
	gorm *gorm.DB
}

func NewUserRepository(gorm *gorm.DB) irepositoryweb2.IUserRepository {
	return &UserRepository{gorm: gorm}
}

func (u *UserRepository) CreateUser(user entitiesweb2.User) (entitiesweb2.User, error) {

}

func (u *UserRepository) DeleteUser(user entitiesweb2.User) (entitiesweb2.User, error) {

}

func (u *UserRepository) UpdateUser(user entitiesweb2.User) (entitiesweb2.User, error) {

}

func (u *UserRepository) CreateUserByManager(manager entitiesweb2.UserWithAuthority) (entitiesweb2.UserWithAuthority, error) {

}

func (u *UserRepository) DeleteUserByManager(manager entitiesweb2.UserWithAuthority) (entitiesweb2.UserWithAuthority, error) {

}

func (u *UserRepository) UpdateUserByManager(manager entitiesweb2.UserWithAuthority) (entitiesweb2.UserWithAuthority, error) {

}
