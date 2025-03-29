package repository

import (
	"backend/internal/adapter/gorm/model"
	entities "backend/internal/domain/entities"
	irepository "backend/internal/domain/irepository"

	"golang.org/x/crypto/bcrypt"
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
	var userModel model.User
	if err := l.gorm.First(&userModel, "username = ?", username).Error; err != nil {
		return entities.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(password)); err != nil {
		return entities.User{}, err
	}

	userModel.Password = ""
	return userModel.ToDomain(), nil
}

// !成功
func (l *LoginRepository) Success(user entities.User) error {
	return nil
}

// !失敗
func (l *LoginRepository) Fail(user entities.User, err error) error {
	return nil
}
