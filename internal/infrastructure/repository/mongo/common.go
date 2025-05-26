package mongo_repository

import (
	"userman/internal/domain/common"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Model[T common.Entity] interface {
	FromDomain(T)
	ToDomain() T
	GetObjectId(string) (bson.ObjectID, error)
}

type BaseModel struct{}

func (b *BaseModel) GetObjectId(id string) (bson.ObjectID, error) {
	return bson.ObjectIDFromHex(id)
}

func ConvertModelsToEntities[T common.Entity, M Model[T]](models []M) []T {
	entities := make([]T, 0, len(models))
	for _, model := range models {
		entities = append(entities, model.ToDomain())
	}

	return entities
}
