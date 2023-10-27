package domain

import (
	"testing"
	
	"github.com/google/uuid"
)

func TestInvoiceClientLoad(t *testing.T) {
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
}