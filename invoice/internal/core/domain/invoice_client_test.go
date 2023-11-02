package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestInvoiceClientLoad(t *testing.T) {
	t.Run("should load invoice client", func(t *testing.T) {
		repo := new(RepoMock)
		client := NewInvoiceClient(repo)
		client.Load("nickname", "clientId", "name", "email", 123456789, 123456789)
		_, err := uuid.Parse(client.GetId())
		if err != nil {
			t.Errorf("id is empty")
		}
		if client.GetNickname() != "nickname" {
			t.Errorf("nickname is not equal")
		}
		if client.GetClientId() != "clientId" {
			t.Errorf("clientId is not equal")
		}
		if client.GetName() != "name" {
			t.Errorf("name is not equal")
		}
		if client.GetEmail() != "email" {
			t.Errorf("email is not equal")
		}
		if client.GetDocument() != 123456789 {
			t.Errorf("document is not equal")
		}
		if client.GetPhone() != 123456789 {
			t.Errorf("phone is not equal")
		}
		if client.GetPhone() != 123456789 {
			t.Errorf("mobile phone is not equal")
		}
	})
}

func TestInvoiceClientSave(t *testing.T) {
	t.Run("should save invoice client", func(t *testing.T) {
		repo := new(RepoMock)
		client := NewInvoiceClient(repo)
		client.Load("nickname", "clientId", "name", "email", 123456789, 123456789)
		if err := client.Save(); err != nil {
			t.Errorf("error on save")
		}
	})
	t.Run("should return error on save", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "saveInvoiceClientError"
		client := NewInvoiceClient(repo)
		client.Load("nickname", "clientId", "name", "email", 123456789, 123456789)
		if err := client.Save(); err == nil {
			t.Errorf("should not be nil")
		}
	})

}

func TestInvoiceClientUpdate(t *testing.T) {
	t.Run("should update invoice client", func(t *testing.T) {
		repo := new(RepoMock)
		client := NewInvoiceClient(repo)
		client.Load("nickname", "clientId", "name", "email", 123456789, 123456789)
		if err := client.Update(); err != nil {
			t.Errorf("error on update")
		}
	})
	t.Run("should return error on update", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "updateInvoiceClientError"
		client := NewInvoiceClient(repo)
		client.Load("nickname", "clientId", "name", "email", 123456789, 123456789)
		if err := client.Update(); err == nil {
			t.Errorf("should not be nil")
		}
	})

}
