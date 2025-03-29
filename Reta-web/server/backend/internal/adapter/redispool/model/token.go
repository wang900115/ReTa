package model

import (
	entities "backend/internal/domain/entities"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserUUID string `json:"user_uuid"`
	Username string `json:"username"`
	Fullname string `json:"full_name"`
	Nickname string `json:"nick_name"`

	jwt.RegisteredClaims `json:"registered_claims"`
}

func (tc TokenClaims) ToDomain() entities.TokenClaims {
	return entities.TokenClaims{
		UserUUID: tc.UserUUID,
		Username: tc.Username,
		Fullname: tc.Fullname,
		Nickname: tc.Nickname,
	}
}

func (tc TokenClaims) FromDomain(tokenClaims entities.TokenClaims) TokenClaims {
	return TokenClaims{
		UserUUID: tokenClaims.UserUUID,
		Username: tokenClaims.Username,
		Fullname: tokenClaims.Fullname,
		Nickname: tokenClaims.Nickname,
	}
}
