package user

import (
	"context"
	"fmt"
	common "userman/internal/application/common/service"
	"userman/internal/domain/user"
)

type GetAllUserInput struct {
	NextCursor string
	PrevCursor string
	Limit      int
}
type UserService interface {
	CreateUser(ctx context.Context, name string, email string, password string) (*user.User, error)
	GetUserByID(ctx context.Context, id string) (*user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	GetAllUser(ctx context.Context, input *GetAllUserInput) (*common.CursorPaginated[*user.User], error)
	UpdateUserByID(ctx context.Context, id string, name string, email string) (*user.User, error)
	DeleteUserByID(ctx context.Context, id string) (*user.User, error)
}

type userService struct {
	userRepo      user.UserRepository
	cursorService common.CursorService[*user.User]
}

func NewUserService(userRepo user.UserRepository) UserService {
	return &userService{userRepo: userRepo, cursorService: common.NewCursorService[*user.User]()}
}

func (u *userService) CreateUser(ctx context.Context, name string, email string, password string) (*user.User, error) {
	res, err := u.userRepo.Create(ctx, user.NewUser(name, email, password))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userService) GetAllUser(ctx context.Context, input *GetAllUserInput) (*common.CursorPaginated[*user.User], error) {
	var cursorStr string
	cursor := &common.CursorData{
		Direction: "next",
		ID:        "",
		Limit:     input.Limit,
	}
	if input.NextCursor != "" {
		cursorStr = input.NextCursor
	} else if input.PrevCursor != "" {
		cursorStr = input.PrevCursor
	}

	if cursorStr != "" {
		res, err := u.cursorService.Decode(cursorStr)
		if err != nil {
			return nil, fmt.Errorf("failed to decode cursor: %w", err)
		}
		cursor = res
	}

	users, err := u.userRepo.GetAllID(ctx, cursor.ID, cursor.Limit, cursor.Direction)
	if err != nil {
		return nil, err
	}

	res, err := u.cursorService.BuildPaginated(users, cursor)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userService) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) UpdateUserByID(ctx context.Context, id string, name string, email string) (*user.User, error) {
	entity, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	entity.Update(&user.UpdateUser{Name: name, Email: email})
	if len(entity.GetChange()) == 0 {
		return entity, nil
	}

	updatedUser, err := u.userRepo.Update(ctx, entity)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *userService) DeleteUserByID(ctx context.Context, id string) (*user.User, error) {
	user, err := u.userRepo.DeleteByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
