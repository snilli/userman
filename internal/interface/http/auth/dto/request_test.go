package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterBodyRequest_Validate(t *testing.T) {
	req := &RegisterBodyRequest{
		Password: "Password123",
	}

	err := req.Validate()
	assert.NoError(t, err)
}

func TestRegisterBodyRequest_Invalid(t *testing.T) {
	req := &RegisterBodyRequest{
		Password: "asdasds",
	}

	err := req.Validate()
	assert.Error(t, err)
}

func TestLoginBodyRequest_Validate(t *testing.T) {
	req := &LoginBodyRequest{
		Password: "ValidPass1",
	}

	err := req.Validate()
	assert.NoError(t, err)
}

func TestLoginBodyRequest_Invalid(t *testing.T) {
	req := &LoginBodyRequest{
		Password: "asdsad",
	}

	err := req.Validate()
	assert.Error(t, err)
}
