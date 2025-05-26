package user

import (
	"context"
	"testing"
	"userman/internal/domain/user"
	user_test "userman/internal/domain/user/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := &user_test.MockUserRepository{}
	service := NewUserService(mockRepo)

	mockUser := &user.User{}
	mockRepo.On("Create", mock.Anything, mock.Anything).Return(mockUser, nil)

	result, err := service.CreateUser(context.Background(), "test", "test@test.com", "password")
	assert.NoError(t, err)
	assert.Equal(t, mockUser, result)
}

func TestUserService_GetUserByID(t *testing.T) {
	mockRepo := &user_test.MockUserRepository{}
	service := NewUserService(mockRepo)

	mockUser := &user.User{}
	mockRepo.On("GetByID", mock.Anything, "123").Return(mockUser, nil)

	result, err := service.GetUserByID(context.Background(), "123")
	assert.NoError(t, err)
	assert.Equal(t, mockUser, result)
}

func TestUserService_GetUserByEmail(t *testing.T) {
	mockRepo := &user_test.MockUserRepository{}
	service := NewUserService(mockRepo)

	mockUser := &user.User{}
	mockRepo.On("GetByEmail", mock.Anything, "test@test.com").Return(mockUser, nil)

	result, err := service.GetUserByEmail(context.Background(), "test@test.com")
	assert.NoError(t, err)
	assert.Equal(t, mockUser, result)
}

func TestUserService_GetAllUser(t *testing.T) {
	mockRepo := &user_test.MockUserRepository{}
	service := NewUserService(mockRepo)

	mockUsers := []*user.User{{}, {}}
	mockRepo.On("GetAllID", mock.Anything, "", 10, "next").Return(mockUsers, nil)

	input := &GetAllUserInput{Limit: 10}
	result, err := service.GetAllUser(context.Background(), input)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestUserService_UpdateUserByID(t *testing.T) {
	mockRepo := &user_test.MockUserRepository{}
	service := NewUserService(mockRepo)

	mockUser := user.NewUser("newname", "XXXXXXXXXXXX", "password")
	mockRepo.On("GetByID", mock.Anything, "123").Return(mockUser, nil)
	mockRepo.On("Update", mock.Anything, mock.Anything).Return(mockUser, nil)

	result, err := service.UpdateUserByID(context.Background(), "123", "newname", "new@test.com")
	assert.NoError(t, err)
	assert.Equal(t, mockUser, result)
}

func TestUserService_DeleteUserByID(t *testing.T) {
	mockRepo := &user_test.MockUserRepository{}
	service := NewUserService(mockRepo)

	mockUser := &user.User{}
	mockRepo.On("DeleteByID", mock.Anything, "123").Return(mockUser, nil)

	result, err := service.DeleteUserByID(context.Background(), "123")
	assert.NoError(t, err)
	assert.Equal(t, mockUser, result)
}
