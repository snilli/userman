package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtService struct {
	secretKey []byte
	expires   time.Duration
}

type TokenDetails struct {
	AccessToken string
	ExpiresIn   int64
}

type JWTService interface {
	GenerateToken(userID string, role string) (*TokenDetails, error)
	ValidateToken(tokenString string) (*Claims, error)
}

func NewJWTService(secretKey string, expires int) JWTService {
	return &jwtService{
		secretKey: []byte(secretKey),
		expires:   time.Duration(expires) * time.Second,
	}
}

func (s *jwtService) GenerateToken(userID string, role string) (*TokenDetails, error) {
	expirationTime := time.Now().Add(s.expires)

	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Subject:   userID,
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return nil, err
	}

	return &TokenDetails{
		AccessToken: tokenString,
		ExpiresIn:   expirationTime.Unix(),
	}, nil
}

func (s *jwtService) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
