package dto

import (
	"testing"
)

func TestClearName(t *testing.T) {
	r, err := formatName("   Jose   da Silva  ")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "Jose Da Silva" {
		t.Errorf("Invalid result: %v", r)
	}
	r, err = formatName("   joSe   da Silva  ")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "Jose Da Silva" {
		t.Errorf("Invalid result: %v", r)
	}
	r, err = formatName("    ")
	if err.Error() != "name is blank" {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "" {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestClearNickname(t *testing.T) {
	r, err := formatNickname("   Jose   da Silva  222")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "jose_da_silva_222" {
		t.Errorf("Invalid result: %v", r)
	}
	r, err = formatNickname("    ")
	if err.Error() != "nickname is blank" {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "" {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestClearDocument(t *testing.T) {
	r, err := formatDocument("206.  656.600-49")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "20665660049" {
		t.Errorf("Invalid result: %v", r)
	}

	r, err = formatDocument("044. 179328-24")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "04417932824" {
		t.Errorf("Invalid result: %v", r)
	}


	if r, err := formatDocument(""); err.Error() != "document is blank" {
		t.Errorf("Invalid result: %v", r)
	}
	if r, err := formatDocument("dasdasdsa--asdasd"); err.Error() != "invalid document" {
		t.Errorf("Invalid result: %v", r)
	}
	if _, err := formatDocument("206.656.600-50"); err.Error() != "invalid document" {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestClearPhone(t *testing.T) {
	r, err := formatPhone("(013) 9-9999-9999")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "5513999999999" {
		t.Errorf("Invalid result: %v", r)
	}
	if _, err := formatPhone(""); err.Error() != "phone is blank" {
		t.Errorf("Invalid result: %v", err)
	}
}

func TestValidate(t *testing.T) {
	c := CreateInputDto{
		Name:     "Jose da Silva",
		Nickname: "jose_da_silva",
		Document: "20665660049",
		Phone:    "5513999999999",
		Email:    "test@test.com",
	}
	if err := c.Validate(); err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	c = CreateInputDto{
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
