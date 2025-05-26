package dto

import (
	"testing"
	"userman/internal/infrastructure/jwt"

	"github.com/stretchr/testify/assert"
)

func TestMapTokenToDto(t *testing.T) {
	tokenDetails := &jwt.TokenDetails{
		AccessToken: "token",
		ExpiresIn:   10,
	}

	result := MapTokenToDto(tokenDetails)

	assert.NotNil(t, result)
	assert.Equal(t, tokenDetails.AccessToken, result.AccessToken)
	assert.Equal(t, tokenDetails.ExpiresIn, result.ExpiresIn)
}
