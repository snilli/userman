package mongo_repository

import (
	"context"
	"errors"
	"sort"
	"userman/internal/domain/user"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) user.UserRepository {
	return &UserRepository{collection: db.Collection("user")}
}

func (r *UserRepository) Create(ctx context.Context, user *user.User) (*user.User, error) {
	model := &UserModel{}
	model.FromDomain(user)

	res, err := r.collection.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}
	model.ID = res.InsertedID.(bson.ObjectID)

	return model.ToDomain(), nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*user.User, error) {
	model := &UserModel{}
	objID, err := model.GetObjectId(id)
	if err != nil {
		return nil, err
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return model.ToDomain(), nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	model := &UserModel{}
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return model.ToDomain(), nil
}
func (r *UserRepository) GetAllID(ctx context.Context, id string, limit int, direction string) ([]*user.User, error) {
	model := &UserModel{}

	filter := bson.M{}
	limit += 1

	if id != "" {
		objID, err := model.GetObjectId(id)
		if err != nil {
			return nil, err
		}

		if direction == "prev" {
			filter["_id"] = bson.M{"$lte": objID}
		} else {
			filter["_id"] = bson.M{"$gt": objID}
		}
	}
	var sortSet bson.D
	if direction == "prev" {
		limit += 1
		sortSet = bson.D{{Key: "_id", Value: -1}}
	} else {
		sortSet = bson.D{{Key: "_id", Value: 1}}
	}

	opts := options.Find().
		SetLimit(int64(limit)).
		SetSort(sortSet)

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	models := make([]*UserModel, 0, limit)
	if err = cursor.All(ctx, &models); err != nil {
		return nil, err
	}
	if direction == "prev" {
		sort.Slice(models, func(i, j int) bool { return models[i].ID.Hex() < models[j].ID.Hex() })
	}

	return ConvertModelsToEntities(models), nil
}

func (r *UserRepository) DeleteByID(ctx context.Context, id string) (*user.User, error) {
	model := &UserModel{}
	objID, err := model.GetObjectId(id)
	if err != nil {
		return nil, err
	}

	err = r.collection.FindOneAndDelete(ctx, bson.M{"_id": objID}).Decode(model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return model.ToDomain(), nil
}

func (r *UserRepository) Update(ctx context.Context, user *user.User) (*user.User, error) {
	model := &UserModel{}
	objID, err := model.GetObjectId(user.ID)
	if err != nil {
		return nil, err
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = r.collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M(user.GetChange())}, opts).Decode(model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return model.ToDomain(), nil
}

func (r *UserRepository) Count(ctx context.Context) (int64, error) {
	return r.collection.CountDocuments(ctx, bson.M{})
}
