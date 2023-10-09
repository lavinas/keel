package dto

import (
	"testing"
)

func TestValidate(t *testing.T) {
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
