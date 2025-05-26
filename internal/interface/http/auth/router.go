package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup, authHandler AuthHandler) {
	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)
}
