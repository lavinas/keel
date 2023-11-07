package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestInvoiceClientLoad(t *testing.T) {
	t.Run("should load invoice client", func(t *testing.T) {
		repo := new(RepoMock)
		client := NewInvoiceClient(repo)
		client.Load("","nickname", "clientId", "name", "email", 123456789, 123456789, time.Time{})
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
	t.Run("should load id and created_at on load invoice client", func(t *testing.T) {
		repo := new(RepoMock)
		client := NewInvoiceClient(repo)
		client.Load(uuid.NewString(),"nickname", "clientId", "name", "email", 123456789, 123456789, time.Now())
		_, err := uuid.Parse(client.GetId())
		if err != nil {
			t.Errorf("id is empty")
		}
		if client.GetCreatedAt().IsZero() {
			t.Errorf("created_at is empty")
		}
	})
}

func TestInvoiceClientSave(t *testing.T) {
	t.Run("should save invoice client", func(t *testing.T) {
		repo := new(RepoMock)
		client := NewInvoiceClient(repo)
		client.Load("","nickname", "clientId", "name", "email", 123456789, 123456789, time.Time{})
		if err := client.Save(); err != nil {
			t.Errorf("error on save")
		}
	})
	t.Run("should return error on save", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "saveInvoiceClientError"
		client := NewInvoiceClient(repo)
		client.Load("","nickname", "clientId", "name", "email", 123456789, 123456789, time.Time{})
		if err := client.Save(); err == nil {
			t.Errorf("should not be nil")
		}
	})

}

func TestInvoiceClientUpdate(t *testing.T) {
	t.Run("should update invoice client", func(t *testing.T) {
		repo := new(RepoMock)
		client := NewInvoiceClient(repo)
		client.Load("","nickname", "clientId", "name", "email", 123456789, 123456789, time.Time{})
		if err := client.Update(); err != nil {
			t.Errorf("error on update")
		}
	})
	t.Run("should return error on update", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "updateInvoiceClientError"
		client := NewInvoiceClient(repo)
		client.Load("","nickname", "clientId", "name", "email", 123456789, 123456789, time.Time{})
		if err := client.Update(); err == nil {
			t.Errorf("should not be nil")
		}
	})

}

func TestInvoiceClientLoadGetClientNicknameDto(t *testing.T) {
	t.Run("should load get client nickname dto", func(t *testing.T) {
		repo := new(RepoMock)
		client := NewInvoiceClient(repo)
		input := new(GetClientByNicknameInputDtoMock)
		if err := client.LoadGetClientNicknameDto(input); err != nil {
			t.Errorf("error on load get client nickname dto")
		}
	})
	t.Run("should return document error on load get client nickname dto", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "loadGetClientNicknameDtoError"
		client := NewInvoiceClient(repo)
		input := new(GetClientByNicknameInputDtoMock)
		input.Status = "documentError"
		if err := client.LoadGetClientNicknameDto(input); err == nil {
			t.Errorf("should not be nil")
		}
	})
	t.Run("should return phone error on load get client nickname dto", func(t *testing.T) {
		repo := new(RepoMock)
		repo.Status = "loadGetClientNicknameDtoError"
		client := NewInvoiceClient(repo)
		input := new(GetClientByNicknameInputDtoMock)
		input.Status = "phoneError"
		if err := client.LoadGetClientNicknameDto(input); err == nil {
			t.Errorf("should not be nil")
		}
	})
}

func TestGetLastInvoiceClient(t *testing.T) {
	t.Run("should get last invoice client", func(t *testing.T) {
		repo := new(RepoMock)
		client := NewInvoiceClient(repo)
		client.GetLastInvoiceClientId("nickname", time.Now())
	})
}
