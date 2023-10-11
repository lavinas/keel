package domain

import (
	"testing"
)

func TestLoad(t *testing.T) {
	t.Run("should load a client", func(t *testing.T) {
		repo := &RepoMock{}
		client := NewClient(repo)
		client.Load("1", "John Doe", "John", 12345678901, 11987654321, "test@test.com")
		if client.ID != "1" {
			t.Errorf("expected 1, got %s", client.ID)
		}
		if client.Name != "John Doe" {
			t.Errorf("expected John Doe, got %s", client.Name)
		}
		if client.Nickname != "John" {
			t.Errorf("expected John, got %s", client.Nickname)
		}
		if client.Document != 12345678901 {
			t.Errorf("expected 12345678901, got %d", client.Document)
		}
		if client.Phone != 11987654321 {
			t.Errorf("expected 11987654321, got %d", client.Phone)
		}
		if client.Email != "test@test.com" {
			t.Errorf("expected test@test.com, got %s", client.Email)
		}
	})
}
