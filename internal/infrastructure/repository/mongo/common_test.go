package mongo_repository

import (
	"testing"
	"userman/internal/domain/common"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockEntity struct {
	*common.BaseEntity
}

type MockModel struct {
	BaseModel
	ID     bson.ObjectID
	entity MockEntity
}

func (m *MockModel) FromDomain(entity MockEntity) {
	m.entity = entity
}

func (m *MockModel) ToDomain() MockEntity {
	return MockEntity{
		BaseEntity: &common.BaseEntity{
			ID: m.ID.Hex(),
		},
	}
}

func TestBaseModel_GetObjectId(t *testing.T) {
	baseModel := &BaseModel{}

	// Valid ObjectID
	objectID, err := baseModel.GetObjectId("507f1f77bcf86cd799439011")
	assert.NoError(t, err)
	assert.NotEqual(t, bson.NilObjectID, objectID)

	// Invalid ObjectID
	_, err = baseModel.GetObjectId("invalid")
	assert.Error(t, err)
}

func TestConvertModelsToEntities(t *testing.T) {
	models := []*MockModel{}
	models = append(models, &MockModel{ID: bson.NewObjectID()})
	models = append(models, &MockModel{ID: bson.NewObjectID()})

	entities := ConvertModelsToEntities(models)

	assert.Len(t, entities, 2)
	assert.NotNil(t, entities[0].GetID())
	assert.NotNil(t, entities[1].GetID())
}

func TestConvertModelsToEntities_Empty(t *testing.T) {
	var models []*MockModel

	entities := ConvertModelsToEntities(models)

	assert.Empty(t, entities)
}
