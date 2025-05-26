package middleware

import (
	"net/http"
	"userman/internal/infrastructure/jwt"
	"userman/internal/interface/http/common/middleware"

	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	jwtService jwt.JWTService
}

func NewAuthMiddleware(jwtService jwt.JWTService) middleware.AuthMiddleware {
	return &authMiddleware{
		jwtService: jwtService,
	}
}

func (m *authMiddleware) Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

var _ middleware.AuthMiddleware = &authMiddleware{}
