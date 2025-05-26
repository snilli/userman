package task

import (
	"context"
	"testing"
	"userman/internal/domain/user"
	user_test "userman/internal/domain/user/mock"

	"github.com/stretchr/testify/mock"
)

func TestSeedingAdminTask_AdminNotExists(t *testing.T) {
	ctx := context.Background()
	newUser := user.NewUser("admin", "admin@admin.com", "Adm1nUserman")
	newUser.SetRole("admin")
	mockRepo := user_test.NewMockUserRepository(t)

	mockRepo.EXPECT().GetByEmail(ctx, "admin@admin.com").Return(nil, nil).Once()
	mockRepo.EXPECT().Create(ctx, mock.AnythingOfType("*user.User")).Return(&user.User{}, nil).Once()

	SeedingAdminTask(ctx, mockRepo)

	mockRepo.AssertNumberOfCalls(t, "GetByEmail", 1)
	mockRepo.AssertCalled(t, "GetByEmail", ctx, "admin@admin.com")
	mockRepo.AssertNumberOfCalls(t, "Create", 1)

}

func TestSeedingAdminTask_AdminExists(t *testing.T) {
	ctx := context.Background()
	mockRepo := user_test.NewMockUserRepository(t)

	newUser := user.NewUser("admin", "admin@admin.com", "Adm1nUserman")
	newUser.SetRole("admin")
	mockRepo.On("GetByEmail", ctx, "admin@admin.com").Return(newUser, nil)

	SeedingAdminTask(ctx, mockRepo)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertNumberOfCalls(t, "GetByEmail", 1)
	mockRepo.AssertCalled(t, "GetByEmail", ctx, "admin@admin.com")
	mockRepo.AssertNotCalled(t, "Create")
}
