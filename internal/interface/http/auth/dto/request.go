package dto

import (
	"errors"
	"strings"
)

type RegisterBodyRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (body *RegisterBodyRequest) Validate() error {
	if !strings.ContainsAny(body.Password, "abcdefghijklmnopqrstuvwxyz") ||
		!strings.ContainsAny(body.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") ||
		!strings.ContainsAny(body.Password, "0123456789") {
		return errors.New("Password must have lowercase, uppercase and number")
	}

	return nil
}

type LoginUriRequest struct {
	ID string `uri:"id" binding:"required,mongodb"`
}

type LoginBodyRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (body *LoginBodyRequest) Validate() error {
	if !strings.ContainsAny(body.Password, "abcdefghijklmnopqrstuvwxyz") ||
		!strings.ContainsAny(body.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") ||
		!strings.ContainsAny(body.Password, "0123456789") {
		return errors.New("Password must have lowercase, uppercase and number")
	}

	return nil
}
