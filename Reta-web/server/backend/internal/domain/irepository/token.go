package irepository

import entities "backend/internal/domain/entities"

type ITokenRepository interface {
	// 產生 token
	Generate(entities.TokenClaims) (string, error)
	// 刷新 token
	Refresh(entities.TokenClaims) (string, error)
	// 認證 token
	Verify(string) (entities.TokenClaims, error)
	// 刪除 token
	Delete(string) error
}
