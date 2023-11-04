package service

import (
	"strings"
	"testing"
)

func TestGetExecute(t *testing.T) {
	t.Run("should get by id", func(t *testing.T) {
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbyid"
		output := &InsertOutputDtoMock{}
		param := "1"
		service := NewGet(log, client)
		err := service.Execute(param, "id", output)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	t.Run("should get by nickname", func(t *testing.T) {
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbynick"
		output := &InsertOutputDtoMock{}
		param := "nickname"
		service := NewGet(log, client)
		err := service.Execute(param, "nickname", output)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	t.Run("should get by email", func(t *testing.T) {
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbyemail"
		output := &InsertOutputDtoMock{}
		param := "email"
		service := NewGet(log, client)
		err := service.Execute(param, "email", output)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	t.Run("should get by document", func(t *testing.T) {
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbydoc"
		output := &InsertOutputDtoMock{}
		param := "12345678"
		service := NewGet(log, client)
		err := service.Execute(param, "document", output)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	t.Run("should get by phone", func(t *testing.T) {
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbyphone"
		output := &InsertOutputDtoMock{}
		param := "12345678"
		service := NewGet(log, client)
		err := service.Execute(param, "phone", output)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	// not found
	t.Run("should return error when not found no numeric", func(t *testing.T) {
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "ok"
		output := &InsertOutputDtoMock{}
		param := "aa"
		service := NewGet(log, client)
		err := service.Execute(param, "document", output)
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if err != nil && !strings.Contains(err.Error(), "bad request") {
			t.Errorf("expected bad request error, got %s", err.Error())
		}
	})
	t.Run("should return error when not found numeric", func(t *testing.T) {
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "ok"
		output := &InsertOutputDtoMock{}
		param := "12345678"
		service := NewGet(log, client)
		err := service.Execute(param, "document", output)
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if err != nil && !strings.Contains(err.Error(), "no content")  {
			t.Errorf("expected no content error, got %s", err.Error())
		}
	})
	// bad request
	t.Run("should return error when blank param", func(t *testing.T) {
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "ok"
		output := &InsertOutputDtoMock{}
		param := ""
		service := NewGet(log, client)
		err := service.Execute(param, "id", output)
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if err != nil && err.Error() != "bad request: blank param" {
			t.Errorf("expected bad request error, got %s", err.Error())
		}
	})
	// invalid param type
	t.Run("should return error when invalid param type", func(t *testing.T) {
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "ok"
		output := &InsertOutputDtoMock{}
		param := "aa"
		service := NewGet(log, client)
		err := service.Execute(param, "type", output)
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if err != nil && strings.Contains(err.Error(), "bad resquest") {
			t.Errorf("expected bad request error, got %s", err.Error())
		}
	})
}