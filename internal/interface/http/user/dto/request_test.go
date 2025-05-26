package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserBodyRequest_Validate(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"Valid password", "Password123", false},
		{"Missing lowercase", "PASSWORD123", true},
		{"Missing uppercase", "password123", true},
		{"Missing number", "PasswordABC", true},
		{"Only lowercase", "password", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &CreateUserBodyRequest{Password: tt.password}
			err := req.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetAllUserQueryRequest_Validate(t *testing.T) {
	tests := []struct {
		name      string
		req       GetAllUserQueryRequest
		wantErr   bool
		wantLimit int
	}{
		{
			name:      "Default limit",
			req:       GetAllUserQueryRequest{},
			wantErr:   false,
			wantLimit: 3,
		},
		{
			name:    "Both tokens error",
			req:     GetAllUserQueryRequest{NextToken: "next", PrevToken: "prev"},
			wantErr: true,
		},
		{
			name:    "Limit with token error",
			req:     GetAllUserQueryRequest{Limit: 10, NextToken: "next"},
			wantErr: true,
		},
		{
			name:      "Valid next token",
			req:       GetAllUserQueryRequest{NextToken: "next"},
			wantErr:   false,
			wantLimit: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantLimit, tt.req.Limit)
			}
		})
	}
}
