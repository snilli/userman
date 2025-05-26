package middleware

import (
	"github.com/gin-gonic/gin"
)

type LoggingMiddleware interface {
	Log() gin.HandlerFunc
}
