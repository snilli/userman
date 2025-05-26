package common

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"userman/internal/domain/common"
)

type CursorData struct {
	ID        string
	Direction string
	Limit     int
}

type CursorPaginated[T any] struct {
	Data       []T
	NextCursor *string
	PrevCursor *string
	TotalCount int
}

type CursorService[T common.Entity] interface {
	Decode(cursor string) (*CursorData, error)
	Encode(data T, direction string, limit int) (string, error)
	BuildPaginated(data []T, cursor *CursorData) (*CursorPaginated[T], error)
}

type cursorService[T common.Entity] struct{}

func NewCursorService[T common.Entity]() CursorService[T] {
	return &cursorService[T]{}
}

func (cm *cursorService[T]) Decode(cursor string) (cursorData *CursorData, err error) {
	cursorData = &CursorData{}
	if cursor == "" {
		return cursorData, nil
	}

	decoded, err := base64.URLEncoding.DecodeString(cursor)
	if err != nil {
		return nil, fmt.Errorf("invalid cursor format: %w", err)
	}

	if err := json.Unmarshal(decoded, cursorData); err != nil {
		return nil, fmt.Errorf("invalid cursor data: %w", err)
	}

	return cursorData, nil
}

func (cm *cursorService[T]) Encode(data T, direction string, limit int) (string, error) {
	entity := any(data).(common.Entity)
	cursorData := CursorData{
		ID:        entity.GetID(),
		Direction: direction,
		Limit:     limit,
	}

	jsonData, err := json.Marshal(cursorData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal cursor data: %w", err)
	}

	encoded := base64.URLEncoding.EncodeToString(jsonData)
	return encoded, nil
}

func (cm *cursorService[T]) BuildPaginated(data []T, cursor *CursorData) (*CursorPaginated[T], error) {
	if len(data) == 0 {
		return &CursorPaginated[T]{
			Data: data,
		}, nil
	}

	hasNext := false
	hasPrev := false

	if cursor.Direction == "next" {
		hasNext = len(data) > cursor.Limit
		if hasNext {
			data = data[:cursor.Limit]
		}
		hasPrev = cursor.ID != "" && cursor.ID != data[0].GetID()
	} else {
		startTrim := 0
		hasPrev = cursor.Limit+2 == len(data)
		if hasPrev {
			startTrim = 1
		}
		hasNext = true
		data = data[startTrim : len(data)-1]
	}

	var nextCursor, prevCursor *string

	if hasNext {
		lastUser := data[len(data)-1]
		cursor, err := cm.Encode(lastUser, "next", cursor.Limit)
		if err != nil {
			return nil, fmt.Errorf("failed to encode next cursor: %w", err)
		}
		nextCursor = &cursor
	}

	if hasPrev {
		firstUser := data[0]
		cursor, err := cm.Encode(firstUser, "prev", cursor.Limit)
		if err != nil {
			return nil, fmt.Errorf("failed to encode prev cursor: %w", err)
		}
		prevCursor = &cursor
	}

	return &CursorPaginated[T]{
		Data:       data,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		TotalCount: len(data),
	}, nil
}
