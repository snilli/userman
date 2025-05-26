package user

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	DeleteByID(ctx context.Context, id string) (*User, error)
	GetAllID(ctx context.Context, id string, limit int, direction string) ([]*User, error)
	Count(ctx context.Context) (int64, error)
}
