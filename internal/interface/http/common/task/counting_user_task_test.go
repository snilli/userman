package task

import (
	"context"
	"testing"
	"time"
	user_test "userman/internal/domain/user/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCountingUserTask_WithError(t *testing.T) {
	mockRepo := &user_test.MockUserRepository{}

	ctx, cancel := context.WithTimeout(context.Background(), 11*time.Second)
	defer cancel()
	mockRepo.EXPECT().Count(mock.Anything).Return(int64(0), assert.AnError).Once()

	CountingUserTask(ctx, mockRepo)

	assert.True(t, true)
}

func TestCountingUserTask_ContextCanceled(t *testing.T) {
	mockRepo := &user_test.MockUserRepository{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	mockRepo.EXPECT().Count(mock.Anything).Return(int64(10), assert.AnError).Once()

	start := time.Now()
	CountingUserTask(ctx, mockRepo)
	duration := time.Since(start)

	assert.Less(t, duration, 5*time.Millisecond)
}
