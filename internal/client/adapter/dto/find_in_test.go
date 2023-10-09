// Package dto is the package that defines the DTOs for the client adapter
package dto

import (
	"testing"
)

func TestFindInDtoValidate(t *testing.T) {
	t.Run("should return nil when the input is valid", func(t *testing.T) {
		input := FindInputDto{
			Page:     "1",
			PerPage:  "10",
			Name:     "John Doe",
			Nickname: "John",
			Document: "12345678901",
			Phone:    "11987654321",
			Email:    "test@test.com",
		}
		err := input.Validate()
		if err != nil {
			t.Errorf("expected nil, got %s", err.Error())
		}
	})
	t.Run("should return error when the page is not a number", func(t *testing.T) {
		input := FindInputDto{
			Page:     "a",
			PerPage:  "10",
			Name:     "John Doe",
			Nickname: "John",
			Document: "12345678901",
			Phone:    "11987654321",
			Email:    "test@test.com",
		}
		err := input.Validate()
		if err.Error() != "page must be a number" {
			t.Errorf("expected page must be a number, got %s", err.Error())
		}
	})
	t.Run("should return error when the perPage is not a number", func(t *testing.T) {
		input := FindInputDto{
			Page:     "1",
			PerPage:  "a",
			Name:     "John Doe",
			Nickname: "John",
			Document: "12345678901",
			Phone:    "11987654321",
			Email:    "test@test.com",
		}
		err := input.Validate()
		if err.Error() != "perPage must be a number" {
			t.Errorf("expected perPage must be a number, got %s", err.Error())
		}
	})
	t.Run("should return error when the document is not a number", func(t *testing.T) {
		input := FindInputDto{
			Page:     "1",
			PerPage:  "10",
			Name:     "John Doe",
			Nickname: "John",
			Document: "a",
			Phone:    "11987654321",
			Email:    "test@test.com",
		}
		err := input.Validate()
		if err.Error() != "document must be a number" {
			t.Errorf("expected document must be a number, got %s", err.Error())
		}
	})
	t.Run("should return error when the phone is not a number", func(t *testing.T) {
		input := FindInputDto{
			Page:     "1",
			PerPage:  "10",
			Name:     "John Doe",
			Nickname: "John",
			Document: "12345678901",
			Phone:    "a",
			Email:    "test@tes.com",
		}
		err := input.Validate()
		if err.Error() != "phone must be a number" {
			t.Errorf("expected phone must be a number, got %s", err.Error())
		}
	})
}

func TestFindInDtoGet(t *testing.T) {
	input := FindInputDto{
		Page:     "1",
		PerPage:  "10",
		Name:     "John Doe",
		Nickname: "John",
		Document: "12345678901",
		Phone:    "11987654321",
		Email:    "test@test.com",
	}
	page, perPage, name, nick, doc, email := input.Get()
	if page != "1" {
		t.Errorf("expected 1, got %s", page)
	}
	if perPage != "10" {
		t.Errorf("expected 10, got %s", perPage)
	}
	if name != "John Doe" {
		t.Errorf("expected John Doe, got %s", name)
	}
	if nick != "John" {
		t.Errorf("expected John, got %s", nick)
	}
	if doc != "12345678901" {
		t.Errorf("expected 12345678901, got %s", doc)
	}
	if email != "test@test.com" {
		t.Errorf("expected test@test.com, got %s", email)
	}
}
