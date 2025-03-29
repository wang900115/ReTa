package repository

import (
	"backend/internal/adapter/redispool/model"
	entities "backend/internal/domain/entities"
	irepository "backend/internal/domain/irepository"
	"crypto/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

const (
	jwtsaltPrefix = "jwtsalt:"
	saltSize      = 16
)

type TokenRepository struct {
	redis      *redis.Client
	expiration time.Duration
	fresh      time.Duration
}

func NewTokenRepository(redis *redis.Client, expiration time.Duration, fresh time.Duration) irepository.ITokenRepository {
	return &TokenRepository{
		redis:      redis,
		expiration: expiration,
		fresh:      fresh,
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
	tokenClaimsModel := model.TokenClaims{}.FromDomain(TokenClaims)
	salt := t.generateRandomSalt(saltSize)
	return jwt.NewWithClaims(jwt.SigningMethodHS512, tokenClaimsModel).SignedString(append([]byte(tokenClaimsModel.UserUUID), salt...))
}

func (t *TokenRepository) Refresh(TokenClaims entities.TokenClaims) (string, error) {
	tokenClaimsModel := model.TokenClaims{}.FromDomain(TokenClaims)
	tokenClaimsModel.ExpiresAt = jwt.NewNumericDate(time.Now().Add(t.expiration))
	return t.Generate(tokenClaimsModel.ToDomain())

}

func (t *TokenRepository) Verify(token string) (entities.TokenClaims, error) {

}

func (t *TokenRepository) Delete(userUUID string) error {

}
