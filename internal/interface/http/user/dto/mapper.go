package dto

import (
	common "userman/internal/application/common/service"
	"userman/internal/domain/user"
)

func MapDomainToDto(user *user.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func MapCursorDomainToDto(cursor *common.CursorPaginated[*user.User]) *UserCursorResponse {
	data := &UserCursorResponse{
		Data:       make([]UserResponse, 0, len(cursor.Data)),
		NextCursor: cursor.NextCursor,
		PrevCursor: cursor.PrevCursor,
		TotalCount: cursor.TotalCount,
	}

	for _, user := range cursor.Data {
		data.Data = append(data.Data, *MapDomainToDto(user))
	}

	return data
}
