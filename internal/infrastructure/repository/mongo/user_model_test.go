package mongo_repository

import (
	"testing"
	"time"
	"userman/internal/domain/common"
	"userman/internal/domain/user"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestUserModel_FromDomain(t *testing.T) {
	now := time.Now()
	domainUser := &user.User{
		BaseEntity: &common.BaseEntity{
			ID:        "507f1f77bcf86cd799439011",
			CreatedAt: now,
			UpdatedAt: now,
		},
		Name:     "test",
		Email:    "test@test.com",
		Password: "password",
		Role:     "user",
	}

	model := &UserModel{}
	model.FromDomain(domainUser)

	assert.Equal(t, "test", model.Name)
	assert.Equal(t, "test@test.com", model.Email)
	assert.Equal(t, "password", model.Password)
	assert.Equal(t, "user", model.Role)
	assert.NotEqual(t, bson.NilObjectID, model.ID)
}

func TestUserModel_ToDomain(t *testing.T) {
	objectID := bson.NewObjectID()
	now := time.Now()

	model := &UserModel{
		ID:        objectID,
		Name:      "test",
		Email:     "test@test.com",
		Password:  "password",
		Role:      "user",
		CreatedAt: bson.NewDateTimeFromTime(now),
		UpdatedAt: bson.NewDateTimeFromTime(now),
	}

	domainUser := model.ToDomain()

	assert.Equal(t, objectID.Hex(), domainUser.ID)
	assert.Equal(t, "test", domainUser.Name)
	assert.Equal(t, "test@test.com", domainUser.Email)
	assert.Equal(t, "password", domainUser.Password)
	assert.Equal(t, "user", domainUser.Role)
}

func TestUserModel_FromDomain_EmptyID(t *testing.T) {
	domainUser := &user.User{
		BaseEntity: &common.BaseEntity{ID: ""},
		Name:       "test",
	}

	model := &UserModel{}
	model.FromDomain(domainUser)

	assert.Equal(t, "test", model.Name)
	assert.Equal(t, bson.NilObjectID, model.ID)
}
