package dto

import (
	"testing"
)

func TestUpdateOutFill(t *testing.T) {
	c := NewUpdateOutputDto()
	c.Fill("1", "Jose da Silva", "jose_da_silva", "20665660049", "5513999999999", "test@test.com")
	if c.Id != "1" {
		t.Errorf("Expected Id to be 1, got %s", c.Id)
	}
	if c.Name != "Jose da Silva" {
		t.Errorf("Expected Name to be Jose da Silva, got %s", c.Name)
	}
	if c.Nickname != "jose_da_silva" {
		t.Errorf("Expected Nickname to be jose_da_silva, got %s", c.Nickname)
	}
	if c.Document != "20665660049" {
		t.Errorf("Expected Document to be 20665660049, got %s", c.Document)
	}
	if c.Phone != "5513999999999" {
		t.Errorf("Invalid result: %v", c.Phone)
	}
	if c.Email != "test@test.com" {
		t.Errorf("Expected Document to be 20665660049, got %s", c.Document)
	}
}
