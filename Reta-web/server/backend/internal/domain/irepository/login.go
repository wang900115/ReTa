package irepository

import entities "backend/internal/domain/entities"

type ILoginRepository interface {
	// 驗證
	Verify(username string, password string) (entities.User, error)
	// 驗證成功
	Success(entities.User) error
	// 驗證失敗
	Fail(entities.User, error) error
}
