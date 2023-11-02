package dto

import (
	"testing"
)

func TestNewGetClientByNicknameInputDto(t *testing.T) {
	t.Run("should not return nil", func(t *testing.T) {
		dto := NewGetClientByNicknameInputDto()
		if dto == nil {
			t.Error("NewGetClientByNicknameInputDto() should not return nil")
		}
	})
}

func TestGetClientByNicknameInputDto_GetId(t *testing.T) {
	t.Run("should return an id", func(t *testing.T) {
		dto := NewGetClientByNicknameInputDto()
		dto.Id = "123"
		if dto.GetId() != "123" {
			t.Error("GetClientByNicknameInputDto.GetId() should return 123")
		}
	})
}

func TestGetClientByNicknameInputDto_GetName(t *testing.T) {
	t.Run("should return a name", func(t *testing.T) {
		dto := NewGetClientByNicknameInputDto()
		dto.Name = "John Doe"
		if dto.GetName() != "John Doe" {
			t.Error("GetClientByNicknameInputDto.GetName() should return John Doe")
		}
	})
}

func TestGetClientByNicknameInputDto_GetNickname(t *testing.T) {
	t.Run("should return a nickname", func(t *testing.T) {
		dto := NewGetClientByNicknameInputDto()
		dto.Nickname = "John"
		if dto.GetNickname() != "John" {
			t.Error("GetClientByNicknameInputDto.GetNickname() should return John")
		}
	})
}

func TestGetClientByNicknameInputDto_GetDocument(t *testing.T) {
	t.Run("should return a document", func(t *testing.T) {
		dto := NewGetClientByNicknameInputDto()
		dto.Document = "12345678901"
		if idoc, _ := dto.GetDocument(); idoc != 12345678901 {
			t.Error("GetClientByNicknameInputDto.GetDocument() should return 12345678901")
		}
	})
	t.Run("should return an invalid document error", func(t *testing.T) {
		dto := NewGetClientByNicknameInputDto()
		dto.Document = "1234567890a"
		if _, err := dto.GetDocument(); err == nil {
			t.Error("GetClientByNicknameInputDto.GetDocument() should return an error")
		}
	})
}

func TestGetClientByNicknameInputDto_GetPhone(t *testing.T) {
	t.Run("should return a phone", func(t *testing.T) {
		dto := NewGetClientByNicknameInputDto()
		dto.Phone = "12345678901"
		if iphone, _ := dto.GetPhone(); iphone != 12345678901 {
			t.Error("GetClientByNicknameInputDto.GetPhone() should return 12345678901")
		}
	})
	t.Run("should return an invalid phone error", func(t *testing.T) {
		dto := NewGetClientByNicknameInputDto()
		dto.Phone = "1234567890a"
		if _, err := dto.GetPhone(); err == nil {
			t.Error("GetClientByNicknameInputDto.GetPhone() should return an error")
		}
	})
}

func TestGetClientByNicknameInputDto_GetEmail(t *testing.T) {
	t.Run("should return an email", func(t *testing.T) {
		dto := NewGetClientByNicknameInputDto()
		dto.Email = "test@test.com"
		if dto.GetEmail() != "test@test.com" {
			t.Error("GetClientByNicknameInputDto.GetEmail() should return test@test.com ")
		}
	})
}
