// Package dto is the package that defines the DTOs for the client adapter
package dto

import (
	"testing"
)

func TestInsertInDtoValidate(t *testing.T) {
	c := InsertInputDto{
		Name:     "Jose da Silva",
		Nickname: "jose_da_silva",
		Document: "20665660049",
		Phone:    "5513999999999",
		Email:    "test@test.com",
	}
	if err := c.Validate(); err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	c = InsertInputDto{
		Name:     "Jose",
		Nickname: "",
		Document: "20665660050",
		Phone:    "5513299999999",
		Email:    "teste",
	}
	if err := c.Validate(); err.Error() != "name should have at least two parts | nickname is blank | invalid document | invalid cell phone | invalid email" {
		t.Errorf("Invalid result: %v", err)
	}
}

func TestInsertInDtoFormat(t *testing.T) {
	c := InsertInputDto{
		Name:     "Jose da Silva",
		Nickname: "jose da   silva",
		Document: " 206.656.600-49",
		Phone:    " 013999999999 ",
		Email:    " test@test.com ",
	}
	if err := c.Format(); err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if c.Name != "Jose Da Silva" {
		t.Errorf("Invalid result: %v", c.Name)
	}
	if c.Nickname != "jose_da_silva" {
		t.Errorf("Invalid result: %v", c.Nickname)
	}
	if c.Document != "20665660049" {
		t.Errorf("Invalid result: %v", c.Document)
	}
	if c.Phone != "5513999999999" {
		t.Errorf("Invalid result: %v", c.Phone)
	}
	if c.Email != "test@test.com" {
		t.Errorf("Invalid result: %v", c.Email)
	}
	c = InsertInputDto{
		Name:     "",
		Nickname: "",
		Document: "",
		Phone:    "",
		Email:    "",
	}
	if err := c.Format(); err.Error() != "name is blank" {
		t.Errorf("Invalid result: %v", err)
	}
	c.Name = "Jose"
	if err := c.Format(); err.Error() != "name should have at least two parts" {
		t.Errorf("Invalid result: %v", err)
	}
	c.Name = "Jose da Silva"
	if err := c.Format(); err.Error() != "nickname is blank" {
		t.Errorf("Invalid result: %v", err)
	}
	c.Nickname = "jose da silva"
	if err := c.Format(); err.Error() != "document is blank" {
		t.Errorf("Invalid result: %v", err)
	}
	c.Document = "20665660050"
	if err := c.Format(); err.Error() != "invalid document" {
		t.Errorf("Invalid result: %v", err)
	}
	c.Document = "20665660049"
	if err := c.Format(); err.Error() != "phone is blank" {
		t.Errorf("Invalid result: %v", err)
	}
	c.Phone = "013299999999"
	if err := c.Format(); err.Error() != "invalid cell phone" {
		t.Errorf("Invalid result: %v", err)
	}
	c.Phone = "013999999999"
	if err := c.Format(); err.Error() != "email is blank" {
		t.Errorf("Invalid result: %v", err)
	}
	c.Email = "teste"
	if err := c.Format(); err.Error() != "invalid email" {
		t.Errorf("Invalid result: %v", err)
	}
}

func TestInsertInDtoGet(t *testing.T) {
	c := InsertInputDto{
		Name:     "Jose da Silva",
		Nickname: "jose_da_silva",
		Document: "20665660049",
		Phone:    "5513999999999",
		Email:    "test@test.com",
	}
	name, nickname, document, phone, email := c.Get()
	if name != "Jose da Silva" {
		t.Errorf("Invalid result: %v", name)
	}
	if nickname != "jose_da_silva" {
		t.Errorf("Invalid result: %v", nickname)
	}
	if document != "20665660049" {
		t.Errorf("Invalid result: %v", document)
	}
	if phone != "5513999999999" {
		t.Errorf("Invalid result: %v", phone)
	}
	if email != "test@test.com" {
		t.Errorf("Invalid result: %v", email)
	}
}

func TestInsertInDtoIsBlank(t *testing.T) {
	c := InsertInputDto{}
	if !c.IsBlank() {
		t.Error("Expected IsBlank to be true, got false")
	}
	c = InsertInputDto{
		Name: "Jose da Silva",
	}
	if c.IsBlank() {
		t.Error("Expected IsBlank to be false, got true")
	}
}
