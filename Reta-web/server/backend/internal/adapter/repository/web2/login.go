package repositorweb2

import (
	entitiesweb2 "backend/internal/domain/entities/web2"
	irepositoryweb2 "backend/internal/domain/irepository/web2"

	"gorm.io/gorm"
)

type LoginRepository struct {
	gorm *gorm.DB
}

func NewLoginRepository(gorm *gorm.DB) irepositoryweb2.ILoginRepository {
	return &LoginRepository{
		gorm: gorm,
	}
}

func (l *LoginRepository) Verify(username, password string) (entitiesweb2.User, error) {

}

func (l *LoginRepository) Success(user entitiesweb2.User) error {

}

func (l *LoginRepository) Fail(user entitiesweb2.User, err error) error {

}
