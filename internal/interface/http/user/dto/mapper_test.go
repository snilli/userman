package dto

import (
	"testing"
	common "userman/internal/application/common/service"
	"userman/internal/domain/user"

	"github.com/stretchr/testify/assert"
)

func TestMapDomainToDto(t *testing.T) {
	domainUser := user.NewUser("name1", "a@a.com", "Abc123qwe")

	result := MapDomainToDto(domainUser)

	assert.Equal(t, domainUser.ID, result.ID)
	assert.Equal(t, domainUser.Name, result.Name)
	assert.Equal(t, domainUser.Email, result.Email)
	assert.Equal(t, domainUser.CreatedAt, result.CreatedAt)
	assert.Equal(t, domainUser.UpdatedAt, result.UpdatedAt)
}

func TestMapCursorDomainToDto(t *testing.T) {
	next := "next"
	prev := "prev"
	users := []*user.User{
		user.NewUser("name1", "a@a.com", "Abc123qwe"),
		user.NewUser("name2", "b@b.com", "Abc123qwe"),
	}

	cursor := &common.CursorPaginated[*user.User]{
		Data:       users,
		NextCursor: &next,
		PrevCursor: &prev,
		TotalCount: 10,
	}

	result := MapCursorDomainToDto(cursor)

	assert.Len(t, result.Data, 2)
	assert.Equal(t, &next, result.NextCursor)
	assert.Equal(t, &prev, result.PrevCursor)
	assert.Equal(t, 10, result.TotalCount)
	assert.Equal(t, "name1", result.Data[0].Name)
	assert.Equal(t, "name2", result.Data[1].Name)
}

func TestMapCursorDomainToDto_EmptyData(t *testing.T) {
	next := "next"
	cursor := &common.CursorPaginated[*user.User]{
		Data:       []*user.User{},
		NextCursor: &next,
		TotalCount: 0,
	}

	result := MapCursorDomainToDto(cursor)

	assert.Len(t, result.Data, 0)
	assert.Equal(t, &next, result.NextCursor)
	assert.Equal(t, 0, result.TotalCount)
}
