package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJWTService_GenerateToken(t *testing.T) {
	service := NewJWTService("secret", 3600)

	token, err := service.GenerateToken("user123", "admin")

	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.NotEmpty(t, token.AccessToken)
	assert.Greater(t, token.ExpiresIn, time.Now().Unix())
}

func TestJWTService_ValidateToken(t *testing.T) {
	service := NewJWTService("secret", 3600)

	// Generate token first
	token, err := service.GenerateToken("user123", "admin")
	assert.NoError(t, err)

	// Validate token
	claims, err := service.ValidateToken(token.AccessToken)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, "user123", claims.UserID)
	assert.Equal(t, "admin", claims.Role)
}

func TestJWTService_ValidateToken_Invalid(t *testing.T) {
	service := NewJWTService("secret", 3600)

	// Test invalid token
	claims, err := service.ValidateToken("invalid.token.string")
	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestJWTService_ValidateToken_WrongSecret(t *testing.T) {
	service1 := NewJWTService("secret1", 3600)
	service2 := NewJWTService("secret2", 3600)

	// Generate with service1
	token, err := service1.GenerateToken("user123", "admin")
	assert.NoError(t, err)

	// Validate with service2 (wrong secret)
	claims, err := service2.ValidateToken(token.AccessToken)
	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestNewJWTService(t *testing.T) {
	service := NewJWTService("test", 300)
	assert.NotNil(t, service)
}
