package auth

import (
	"net/http/httptest"
	"testing"
	auth_test "userman/internal/interface/http/auth/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group("/auth")
	handler := auth_test.NewMockAuthHandler(t)
	handler.EXPECT().Login(mock.Anything).Once()

	AuthRouter(group, handler)
	req := httptest.NewRequest("POST", "/auth/login", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
