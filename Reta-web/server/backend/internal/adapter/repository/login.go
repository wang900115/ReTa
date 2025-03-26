package repository

import (
	entities "backend/internal/domain/entities"
	irepository "backend/internal/domain/irepository"

	"gorm.io/gorm"
)

type LoginRepository struct {
	gorm *gorm.DB
}

func NewLoginRepository(gorm *gorm.DB) irepository.ILoginRepository {
	return &LoginRepository{
		gorm: gorm,
	}
}

func (l *LoginRepository) Verify(username, password string) (entities.User, error) {

}

func (l *LoginRepository) Success(user entities.User) error {

}

func (l *LoginRepository) Fail(user entities.User, err error) error {

}
