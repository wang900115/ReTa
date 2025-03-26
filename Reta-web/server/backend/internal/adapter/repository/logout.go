package repository

import (
	entities "backend/internal/domain/entities"
	irepository "backend/internal/domain/irepository"

	"gorm.io/gorm"
)

type LogoutRepository struct {
	gorm *gorm.DB
}

func NewLogoutRepository(gorm *gorm.DB) irepository.ILogoutRepository {
	return &LogoutRepository{
		gorm: gorm,
	}
}

func (l *LogoutRepository) Verify(user entities.User, password string) (entities.User, error) {

}

func (l *LogoutRepository) Success(user entities.User) error {

}

func (l *LogoutRepository) Fail(user entities.User, err error) error {

}
