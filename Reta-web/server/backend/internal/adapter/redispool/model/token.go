package model

import (
	entitiesweb2 "backend/internal/domain/entities/web2"
)

type TokenClaims struct {
	UserUUID string `json:"user_uuid"`
	Username string `json:"username"`
	Fullname string `json:"full_name"`
	Nickname string `json:"nick_name"`
}

func (tc TokenClaims) ToDomain() entitiesweb2.TokenClaims {
	return entitiesweb2.TokenClaims{
		UserUUID: tc.UserUUID,
		Username: tc.Username,
		Fullname: tc.Fullname,
		Nickname: tc.Nickname,
	}
}

func (tc TokenClaims) FromDomain(tokenClaims entitiesweb2.TokenClaims) TokenClaims {
	return TokenClaims{
		UserUUID: tokenClaims.UserUUID,
		Username: tokenClaims.Username,
		Fullname: tokenClaims.Fullname,
		Nickname: tokenClaims.Nickname,
	}
}
