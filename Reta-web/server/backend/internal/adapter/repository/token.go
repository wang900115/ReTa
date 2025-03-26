package repository

import (
	entities "backend/internal/domain/entities"
	irepository "backend/internal/domain/irepository"
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

func NewTokenRepository(redis *redis.Client) irepository.ITokenRepository {
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

func (t *TokenRepository) Generate(TokenClaims entities.TokenClaims) (string, error) {

}

func (t *TokenRepository) Verify(token string) (entities.TokenClaims, error) {

}

func (t *TokenRepository) Delete(userUUID string) error {

}
