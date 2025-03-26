package irepository

import entities "backend/internal/domain/entities"

type ILoginRepository interface {
	Verify(username string, password string) (entities.User, error)
	Success(entities.User) error
	Fail(entities.User, error) error
}
