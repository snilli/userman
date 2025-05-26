package mongo_repository

import (
	"userman/internal/domain/common"
	"userman/internal/domain/user"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserModel struct {
	BaseModel
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Name      string        `bson:"name"`
	Email     string        `bson:"email"`
	Password  string        `bson:"password"`
	Role      string        `bson:"role"`
	CreatedAt bson.DateTime `bson:"created_at"`
	UpdatedAt bson.DateTime `bson:"updated_at"`
}

func (m *UserModel) FromDomain(user *user.User) {
	if user.ID != "" {
		if oid, err := bson.ObjectIDFromHex(user.ID); err == nil {
			m.ID = oid
		}
	}

	m.Name = user.Name
	m.Email = user.Email
	m.Password = user.Password
	m.Role = user.Role
	m.CreatedAt = bson.NewDateTimeFromTime(user.CreatedAt)
	m.UpdatedAt = bson.NewDateTimeFromTime(user.UpdatedAt)
}

func (m *UserModel) ToDomain() *user.User {
	return &user.User{
		BaseEntity: &common.BaseEntity{
			ID:        m.ID.Hex(),
			CreatedAt: m.CreatedAt.Time(),
			UpdatedAt: m.UpdatedAt.Time(),
		},
		Name:     m.Name,
		Email:    m.Email,
		Password: m.Password,
		Role:     m.Role,
	}
}

var _ Model[*user.User] = &UserModel{}
