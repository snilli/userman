package dto

import (
	"errors"
	"strings"
)

type CreateUserBodyRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (body *CreateUserBodyRequest) Validate() error {
	if !strings.ContainsAny(body.Password, "abcdefghijklmnopqrstuvwxyz") ||
		!strings.ContainsAny(body.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") ||
		!strings.ContainsAny(body.Password, "0123456789") {
		return errors.New("Password must have lowercase, uppercase and number")
	}

	return nil
}

type GetUserQueryRequest struct {
	ID string `uri:"id" binding:"required,mongodb"`
}

type GetAllUserQueryRequest struct {
	NextToken string `form:"next_token" binding:"omitempty"`
	PrevToken string `form:"prev_token" binding:"omitempty"`
	Limit     int    `form:"limit" binding:"omitempty,min=1,max=100"`
}

func (query *GetAllUserQueryRequest) Validate() error {
	if query.Limit > 0 && (query.NextToken != "" || query.PrevToken != "") {
		return errors.New("Cannot use both limit and token together")
	}

	if query.NextToken != "" && query.PrevToken != "" {
		return errors.New("Cannot use both tokens together")
	}

	if query.Limit == 0 {
		query.Limit = 3
	}

	return nil
}

type UpdateUserUriRequest struct {
	ID string `uri:"id" binding:"required,mongodb"`
}

type UpdateUserBodyRequest struct {
	Name  string `json:"name" binding:"required,min=2,max=100"`
	Email string `json:"email" binding:"required,email"`
}

type DeleteUserUriRequest struct {
	ID string `uri:"id" binding:"required,mongodb"`
}
