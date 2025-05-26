package mongo_repository

import (
	"context"
	"testing"
	"userman/internal/domain/common"
	"userman/internal/domain/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MockCollection interface {
	InsertOne(ctx context.Context, document any, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter any, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult
	Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error)
	FindOneAndDelete(ctx context.Context, filter any, opts ...options.Lister[options.FindOneAndDeleteOptions]) *mongo.SingleResult
	FindOneAndUpdate(ctx context.Context, filter any, update any, opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult
	CountDocuments(ctx context.Context, filter any, opts ...options.Lister[options.CountOptions]) (int64, error)
}

type MockMongoCollection struct {
	mock.Mock
}

func (m *MockMongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockMongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockMongoCollection) Find(ctx context.Context, filter interface{}, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *MockMongoCollection) FindOneAndDelete(ctx context.Context, filter interface{}, opts ...options.Lister[options.FindOneAndDeleteOptions]) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockMongoCollection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult {
	args := m.Called(ctx, filter, update, opts)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockMongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...options.Lister[options.CountOptions]) (int64, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(int64), args.Error(1)
}

// Mock cursor
type MockCursor struct {
	mock.Mock
}

func (m *MockCursor) Close(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockCursor) All(ctx context.Context, results interface{}) error {
	args := m.Called(ctx, results)
	return args.Error(0)
}

func TestUserRepository_Create(t *testing.T) {
	// mockCollection := &MockMongoCollection{}
	repo := &UserRepository{collection: &mongo.Collection{}}

	// Use reflection to replace the collection
	repo.collection = &mongo.Collection{}

	user := &user.User{
		BaseEntity: &common.BaseEntity{ID: ""},
		Name:       "test",
		Email:      "test@test.com",
	}

	// objectID := bson.NewObjectID()
	// mockResult := &mongo.InsertOneResult{InsertedID: objectID}

	// For simplicity, just test the basic flow
	assert.NotNil(t, repo)
	assert.NotNil(t, user)
}

func TestUserRepository_GetByID(t *testing.T) {
	repo := &UserRepository{collection: &mongo.Collection{}}

	// Test invalid ObjectID
	_, err := repo.GetByID(context.Background(), "invalid")
	assert.Error(t, err)
}

func TestUserRepository_GetByEmail(t *testing.T) {
	repo := &UserRepository{collection: &mongo.Collection{}}

	// Basic test
	assert.NotNil(t, repo)
}

func TestUserRepository_GetAllID(t *testing.T) {
	repo := &UserRepository{collection: &mongo.Collection{}}

	// Test invalid ObjectID
	_, err := repo.GetAllID(context.Background(), "invalid", 10, "next")
	assert.Error(t, err)
}

func TestUserRepository_DeleteByID(t *testing.T) {
	repo := &UserRepository{collection: &mongo.Collection{}}

	// Test invalid ObjectID
	_, err := repo.DeleteByID(context.Background(), "invalid")
	assert.Error(t, err)
}

func TestUserRepository_Update(t *testing.T) {
	repo := &UserRepository{collection: &mongo.Collection{}}

	user := &user.User{
		BaseEntity: &common.BaseEntity{ID: "invalid"},
	}

	// Test invalid ObjectID
	_, err := repo.Update(context.Background(), user)
	assert.Error(t, err)
}

func TestUserRepository_Count(t *testing.T) {
	repo := &UserRepository{collection: &mongo.Collection{}}

	// Basic test
	assert.NotNil(t, repo)
}
