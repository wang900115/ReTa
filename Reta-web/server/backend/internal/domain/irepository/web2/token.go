package irepositoryweb2

import entitiesweb2 "backend/internal/domain/entities/web2"

type ITokenRepository interface {
	Generate(entitiesweb2.TokenClaims) (string, error)
	Verify(string) (entitiesweb2.TokenClaims, error)
	Delete(string) error
}
