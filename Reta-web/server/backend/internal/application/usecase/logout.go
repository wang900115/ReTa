package usecase

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/irepository"
)

type LogoutUsecase struct {
	logoutRepo irepository.ILogoutRepository
}

func NewLogoutUsecase(logoutRepo irepository.ILogoutRepository) *LogoutUsecase {
	return &LogoutUsecase{logoutRepo: logoutRepo}
}

func (lu *LogoutUsecase) Logout(user entities.User) error {
	if err := lu.logoutRepo.Verify(user); err != nil {
		return lu.logoutRepo.Fail(user, err)
	}
	return lu.logoutRepo.Success(user)
}
