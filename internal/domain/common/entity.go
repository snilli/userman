package common

import "time"

type Entity interface {
	GetID() string
	GetChange() map[string]any
	SetChange(key string, value any)
}

type BaseEntity struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	change    map[string]any
}

func (e *BaseEntity) GetID() string {
	return e.ID
}

func (e *BaseEntity) SetChange(key string, value any) {
	if e.change == nil {
		e.change = make(map[string]any)
	}
	e.change[key] = value
}

func (e *BaseEntity) GetChange() map[string]any {
	return e.change
}
