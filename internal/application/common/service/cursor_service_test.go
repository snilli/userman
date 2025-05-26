package common

import (
	"testing"
	"userman/internal/domain/common"

	"github.com/stretchr/testify/assert"
)

func TestCursorService_Decode(t *testing.T) {

	service := NewCursorService[*common.BaseEntity]()
	// Empty cursor
	result, err := service.Decode("")
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Valid cursor
	validCursor := "eyJJRCI6InRlc3QiLCJEaXJlY3Rpb24iOiJuZXh0IiwiTGltaXQiOjEwfQ=="
	result, err = service.Decode(validCursor)
	assert.NoError(t, err)
	assert.Equal(t, "test", result.ID)
}

func TestCursorService_Encode(t *testing.T) {
	entity := common.BaseEntity{ID: "test123"}
	service := NewCursorService[*common.BaseEntity]()

	cursor, err := service.Encode(&entity, "next", 10)
	assert.NoError(t, err)
	assert.NotEmpty(t, cursor)
}

func TestCursorService_BuildPaginated(t *testing.T) {
	service := NewCursorService[*common.BaseEntity]()

	// Empty data
	cursorData := &CursorData{Direction: "next", Limit: 10}
	result, err := service.BuildPaginated([]*common.BaseEntity{}, cursorData)
	assert.NoError(t, err)
	assert.Empty(t, result.Data)

	// With data
	data := []*common.BaseEntity{{ID: "1"}, {ID: "2"}}
	result, err = service.BuildPaginated(data, cursorData)
	assert.NoError(t, err)
	assert.Len(t, result.Data, 2)
}
