package user

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.RouterGroup, userHandler UserHandler) {
	api.POST("/", userHandler.CreateUser)
	api.GET("/:id", userHandler.GetUser)
	api.GET("/", userHandler.GetAllUser)
	api.DELETE("/:id", userHandler.DeleteUser)
	api.PATCH("/:id", userHandler.UpdateUser)
}
