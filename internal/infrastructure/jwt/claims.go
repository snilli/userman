package jwt

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func (c Claims) IsAdmin() bool {
	return c.Role == "admin"
}
