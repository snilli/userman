package middleware

import (
	"log"
	"time"
	"userman/internal/interface/http/common/middleware"

	"github.com/gin-gonic/gin"
)

type loggingMiddleware struct{}

func NewLoggingMiddleware() middleware.LoggingMiddleware {
	return &loggingMiddleware{}
}

func (lm *loggingMiddleware) Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		duration := time.Since(start)
		log.Printf("[%s] %s - %v", method, path, duration)
	}
}

var _ middleware.LoggingMiddleware = &loggingMiddleware{}
