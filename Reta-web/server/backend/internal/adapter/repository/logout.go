package repository

import (
	"backend/internal/adapter/gorm/model"
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

func (l *LogoutRepository) Verify(user entities.User) error {
	var userModel model.User
	if err := l.gorm.First(&userModel, "uuid = ?", user.UUID).Error; err != nil {
		return entities.User{}, err
	}
	userModel.Password = ""
	return userModel.ToDomain(), nil
}

// !成功
func (l *LogoutRepository) Success(user entities.User) error {
	return nil
}

// !失敗
func (l *LogoutRepository) Fail(user entities.User, err error) error {
	return nil
}
