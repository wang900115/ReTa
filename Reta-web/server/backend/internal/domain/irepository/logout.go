package irepository

import entities "backend/internal/domain/entities"

type ILogoutRepository interface {
	Verify(user entities.User, uuid string) (entities.User, error)
	Success(entities.User) error
	Fail(entities.User, error) error
}
