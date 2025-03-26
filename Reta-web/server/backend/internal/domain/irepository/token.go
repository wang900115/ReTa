package irepository

import entities "backend/internal/domain/entities"

type ITokenRepository interface {
	Generate(entities.TokenClaims) (string, error)
	Verify(string) (entities.TokenClaims, error)
	Delete(string) error
}
