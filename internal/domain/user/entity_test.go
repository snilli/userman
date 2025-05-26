package user

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	name := "John Doe"
	email := "john@example.com"
	password := "password123"

	user := NewUser(name, email, password)

	if user == nil {
		t.Fatal("Expected user to be created, got nil")
	}

	if user.Name != name {
		t.Errorf("Expected name %s, got %s", name, user.Name)
	}

	if user.Email != email {
		t.Errorf("Expected email %s, got %s", email, user.Email)
	}

	if user.Role != "user" {
		t.Errorf("Expected default role 'user', got %s", user.Role)
	}

	t.Run("Change Role should be work", func(t *testing.T) {
		result := user.IsAdmin()
		if result {
			t.Error("Expected IsAdmin() to return false")
		}
		user.SetRole("admin")
		result = user.IsAdmin()

		if !result {
			t.Error("Expected IsAdmin() to return true")
		}
	})

	t.Run("Check Password should be work", func(t *testing.T) {
		if err := user.CheckPassword(password); err != nil {
			t.Error("Expected password should be correct")
		}
	})

	t.Run("Set Role should be work", func(t *testing.T) {
		user.SetRole("admin")
		if user.Role != "admin" {
			t.Errorf("Expected role 'admin', got %s", user.Role)
		}
	})

	t.Run("Update should be work", func(t *testing.T) {
		newName := "Jane Doe"
		newEmail := "jane@example.com"

		user.Update(&UpdateUser{Name: newName, Email: newEmail})
		if user.Name != newName {
			t.Errorf("Expected name %s, got %s", newName, user.Role)
		}

		if user.Email != newEmail {
			t.Errorf("Expected role %s, got %s", newEmail, user.Role)
		}
	})
}
