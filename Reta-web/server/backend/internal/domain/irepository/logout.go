package irepository

import entities "backend/internal/domain/entities"

type ILogoutRepository interface {
	// 驗證
	Verify(user entities.User) error
	// 驗證成功
	Success(entities.User) error
	// 驗證失敗
	Fail(entities.User, error) error
}
