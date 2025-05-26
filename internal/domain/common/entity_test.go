package common

import (
	"reflect"
	"testing"
	"time"
)

func TestBaseEntity(t *testing.T) {
	e := BaseEntity{
		ID:        "a",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("Get ID", func(t *testing.T) {
		if e.GetID() != "a" {
			t.Errorf("Expected id %s, got %s", "a", e.GetID())
		}
	})

	t.Run("Set and Get change", func(t *testing.T) {
		e.SetChange("a", 1)
		change := e.GetChange()
		if !reflect.DeepEqual(change, map[string]any{"a": 1}) {
			t.Error("Expected change should be same value")
		}
	})
}
