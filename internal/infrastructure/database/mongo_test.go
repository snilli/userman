package database

import "testing"

func TestNewMongoDatabase_InvalidURI(t *testing.T) {
	_, err := NewMongoDatabase("invalid-uri", "testdb")
	if err == nil {
		t.Error("Expected error for invalid URI, got nil")
	}
}

func TestNewMongoDatabase_ValidParams(t *testing.T) {
	db, err := NewMongoDatabase("mongodb://localhost:27017", "testdb")
	if err != nil {
		t.Skipf("MongoDB not available: %v", err)
	}

	if db == nil {
		t.Error("Expected database instance, got nil")
	}
}
