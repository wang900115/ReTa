package repositoryweb2

import (
	entitiesweb2 "backend/internal/domain/entities/web2"
	irepositoryweb2 "backend/internal/domain/irepository/web2"
	"crypto/rand"

	"github.com/redis/go-redis/v9"
)

const (
	jwtsaltPrefix = "jwtsalt:"
	saltSize      = 16
)

type TokenRepository struct {
	redis *redis.Client
}

func NewTokenRepository(redis *redis.Client) irepositoryweb2.ITokenRepository {
	return &TokenRepository{
		redis: redis,
	}
}

func (TokenRepository) prefixName(userUUID string) string {
	return jwtsaltPrefix + userUUID
}

func (TokenRepository) generateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return salt
}

func (t *TokenRepository) Generate(TokenClaims entitiesweb2.TokenClaims) (string, error) {

}

func (t *TokenRepository) Verify(token string) (entitiesweb2.TokenClaims, error) {

}

func (t *TokenRepository) Delete(userUUID string) error {

}
