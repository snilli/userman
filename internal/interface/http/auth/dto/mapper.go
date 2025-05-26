package dto

import (
	"userman/internal/infrastructure/jwt"
)

func MapTokenToDto(token *jwt.TokenDetails) *AccessTokenResponse {
	return &AccessTokenResponse{
		AccessToken: token.AccessToken,
		ExpiresIn:   token.ExpiresIn,
	}
}
