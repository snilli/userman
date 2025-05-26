package user

import (
	"net/http/httptest"
	"testing"
	user_test "userman/internal/interface/http/user/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group("/user")
	handler := user_test.NewMockUserHandler(t)
	handler.EXPECT().GetAllUser(mock.Anything).Once()

	UserRouter(group, handler)
	req := httptest.NewRequest("GET", "/user/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
