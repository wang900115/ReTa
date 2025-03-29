package usecase

import (
	"backend/internal/domain/irepository"
)

type LoginUsecase struct {
	loginRepo irepository.ILoginRepository
}

func NewLoginUsecase(loginRepo irepository.ILoginRepository) *LoginUsecase {
	return &LoginUsecase{loginRepo: loginRepo}
}

func (lu *LoginUsecase) Login(username string, password string) error {
	user, err := lu.loginRepo.Verify(username, password)
	if err != nil {
		return lu.loginRepo.Fail(user, err)
	}
	return lu.loginRepo.Success(user)
}
