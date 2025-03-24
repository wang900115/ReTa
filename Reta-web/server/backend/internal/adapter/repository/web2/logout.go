package repositoryweb2

import (
	entitiesweb2 "backend/internal/domain/entities/web2"
	irepositoryweb2 "backend/internal/domain/irepository/web2"

	"gorm.io/gorm"
)

type LogoutRepository struct {
	gorm *gorm.DB
}

func NewLogoutRepository(gorm *gorm.DB) irepositoryweb2.ILogoutRepository {
	return &LogoutRepository{
		gorm: gorm,
	}
}

func (l *LogoutRepository) Verify(username, password string) (entitiesweb2.User, error) {

}

func (l *LogoutRepository) Success(user entitiesweb2.User) error {

}

func (l *LogoutRepository) Fail(user entitiesweb2.User, err error) error {

}
