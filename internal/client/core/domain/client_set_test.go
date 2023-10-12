package domain

import (
	"testing"
)

func TestClientSetLoad(t *testing.T) {
	t.Run("should load a client", func(t *testing.T) {
		repo := &RepoMock{}
		clientSet := NewClientSet(repo)
		err := clientSet.Load(1, 10, "John Doe", "John", "12345678901", "11987654321", "test@test.com")
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
}

func TestClientSetAppend(t *testing.T) {
	t.Run("should append a client", func(t *testing.T) {
		repo := &RepoMock{}
		clientSet := NewClientSet(repo)
		clientSet.Append("1", "John Doe", "John", 12345678901, 11987654321, "test@test.com")
		if clientSet.Count() != 1 {
			t.Errorf("Error: %d", clientSet.Count())
		}
		if clientSet.set[0].Name != "John Doe" {
			t.Errorf("Error: %s", clientSet.set[0].Name)
		}
		if clientSet.set[0].Nickname != "John" {
			t.Errorf("Error: %s", clientSet.set[0].Nickname)
		}
		if clientSet.set[0].Document != 12345678901 {
			t.Errorf("Error: %d", clientSet.set[0].Document)
		}
		if clientSet.set[0].Phone != 11987654321 {
			t.Errorf("Error: %d", clientSet.set[0].Phone)
		}
		if clientSet.set[0].Email != "test@test.com" {
			t.Errorf("Error: %s", clientSet.set[0].Email)
		}
	})
}

func TestClientSetSetOutput(t *testing.T) {
	t.Run("should set output", func(t *testing.T) {
		repo := &RepoMock{}
		clientSet := NewClientSet(repo)
		clientSet.Append("1", "John Doe", "John", 12345678901, 11987654321, "test@test.com")
		output := FindOutputDtoMock{}
		clientSet.SetOutput(&output)
		if output.Count() != 1 {
			t.Errorf("Error: %d", output.Count())
		}
		if output.Clients[0].Name != "John Doe" {
			t.Errorf("Error: %s", output.Clients[0].Name)
		}
		if output.Clients[0].Nickname != "John" {
			t.Errorf("Error: %s", output.Clients[0].Nickname)
		}
		if output.Clients[0].Document != 12345678901 {
			t.Errorf("Error: %d", output.Clients[0].Document)
		}
		if output.Clients[0].Phone != 11987654321 {
			t.Errorf("Error: %d", output.Clients[0].Phone)
		}
		if output.Clients[0].Email != "test@test.com" {
			t.Errorf("Error: %s", output.Clients[0].Email)
		}
	})
}
