package dto

import (
	"testing"
)

func TestUpdateInDtoIsBlank(t *testing.T) {
	c := UpdateInputDto{}
	if !c.IsBlank() {
		t.Errorf("Expected IsBlank to be true, got false")
	}
	c.Name = "Jose da Silva"
	if c.IsBlank() {
		t.Errorf("Expected IsBlank to be false, got true")
	}
}

func TestUpdateInDtoValidate(t *testing.T) {
	c := UpdateInputDto{
		Name:     "Jose da Silva",
		Nickname: "jose_da_silva",
		Document: "20665660049",
		Phone:    "5513999999999",
		Email:    "test@test.com",
	}
	if err := c.Validate(); err != nil {
		t.Error("Expected Validate to be valid, got invalid")
	}
	c = UpdateInputDto{
		Name: "Jose da Silva",
	}
	if err := c.Validate(); err != nil {
		t.Error("Expected Validate to be valid, got invalid")
	}
	c = UpdateInputDto{
		Nickname: "jose_da_silva",
	}
	if err := c.Validate(); err != nil {
		t.Error("Expected Validate to be valid, got invalid")
	}
	c = UpdateInputDto{
		Document: "20665660049",
	}
	if err := c.Validate(); err != nil {
		t.Error("Expected Validate to be valid, got invalid")
	}
	c = UpdateInputDto{
		Phone: "5513999999999",
	}
	if err := c.Validate(); err != nil {
		t.Error("Expected Validate to be valid, got invalid")
	}
	c = UpdateInputDto{
		Email: "test@test.com",
	}
	if err := c.Validate(); err != nil {
		t.Error("Expected Validate to be valid, got invalid")
	}

	c = UpdateInputDto{}
	if err := c.Validate(); err.Error() != "at least one field must be filled" {
		t.Errorf("Invalid error message: %s", err.Error())
	}
	c = UpdateInputDto{
		Name:     "Jose",
		Nickname: "",
		Document: "20665660050",
		Phone:    "5513299999999",
		Email:    "test",
	}
	if err := c.Validate(); err.Error() != "name should have at least two parts | invalid document | invalid cell phone | invalid email" {
		t.Errorf("Invalid error message: %s", err.Error())
	}
}

func TestUpdateInDtoFormat(t *testing.T) {
	c := UpdateInputDto{
		Name:     "Jose da Silva",
		Nickname: "jose da   silva",
		Document: " 206.656.600-49",
		Phone:    " 013999999999 ",
		Email:    "test@test.com",
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
	c = UpdateInputDto{
		Name:     "",
		Nickname: "",
		Document: "",
		Phone:    "",
		Email:    "",
	}
	if err := c.Format(); err.Error() != "at least one field must be filled" {
		t.Errorf("Invalid result: %v", err)
	}
	c.Name = "Jose"
	if err := c.Format(); err.Error() != "name should have at least two parts" {
		t.Errorf("Invalid result: %v", err)
	}
	c.Name = "Jose da Silva"
	c.Document = "20665660050"
	if err := c.Format(); err.Error() != "invalid document" {
		t.Errorf("Invalid result: %v", err)
	}
	c.Document = "20665660049"
	c.Phone = "5513299999999"
	if err := c.Format(); err.Error() != "invalid cell phone" {
		t.Errorf("Invalid result: %v", err)
	}
	c.Phone = "5513999999999"
	c.Email = "teste"
	if err := c.Format(); err.Error() != "invalid email" {
		t.Errorf("Invalid result: %v", err)
	}
}




func TestUpdateInDtoGet(t *testing.T) {
	c := UpdateInputDto{
		Name:     "Jose da Silva",
		Nickname: "jose_da_silva",
		Document: "20665660049",
		Phone:    "5513999999999",
		Email:    "test@test.com",
	}
	name, nick, doc, phone, email := c.Get()
	if name != "Jose da Silva" {
		t.Errorf("Invalid result: %v", name)
	}
	if nick != "jose_da_silva" {
		t.Errorf("Invalid result: %v", nick)
	}
	if doc != "20665660049" {
		t.Errorf("Invalid result: %v", doc)
	}
	if phone != "5513999999999" {
		t.Errorf("Invalid result: %v", phone)
	}
	if email != "test@test.com" {
		t.Errorf("Invalid result: %v", email)
	}
}
