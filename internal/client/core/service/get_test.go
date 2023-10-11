package service

import (
	"testing"
)

func TestGetExecute(t *testing.T) {
	t.Run("should get by id", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbyid"
		output := &InsertOutputDtoMock{}
		param := "1"
		service := NewGet(log, client, param, output)
		err := service.Execute()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	t.Run("should get by nickname", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbynick"
		output := &InsertOutputDtoMock{}
		param := "nickname"
		service := NewGet(log, client, param, output)
		err := service.Execute()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	t.Run("should get by email", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbyemail"
		output := &InsertOutputDtoMock{}
		param := "email"
		service := NewGet(log, client, param, output)
		err := service.Execute()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	t.Run("should get by document", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbydoc"
		output := &InsertOutputDtoMock{}
		param := "12345678"
		service := NewGet(log, client, param, output)
		err := service.Execute()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	t.Run("should get by phone", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbyphone"
		output := &InsertOutputDtoMock{}
		param := "12345678"
		service := NewGet(log, client, param, output)
		err := service.Execute()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	// not found
	t.Run("should return error when not found no numeric", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "ok"
		output := &InsertOutputDtoMock{}
		param := "aa"
		service := NewGet(log, client, param, output)
		err := service.Execute()
		if err == nil {
			t.Errorf("Error: %s", err.Error())
		}
		if err.Error() != "not found: "+param {
			t.Errorf("Error: %s", err.Error())
		}
	})
	t.Run("should return error when not found numeric", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "ok"
		output := &InsertOutputDtoMock{}
		param := "12345678"
		service := NewGet(log, client, param, output)
		err := service.Execute()
		if err == nil {
			t.Errorf("Error: %s", err.Error())
		}
		if err.Error() != "not found: "+param {
			t.Errorf("Error: %s", err.Error())
		}
	})
	// bad request
	t.Run("should return error when blank param", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "ok"
		output := &InsertOutputDtoMock{}
		param := ""
		service := NewGet(log, client, param, output)
		err := service.Execute()
		if err == nil {
			t.Errorf("Error: %s", err.Error())
		}
		if err.Error() != "bad request: blank param" {
			t.Errorf("Error: %s", err.Error())
		}
	})
	// get findbyid error
	t.Run("should return error when findbyid error", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbyiderror"
		output := &InsertOutputDtoMock{}
		param := "1"
		service := NewGet(log, client, param, output)
		err := service.Execute()
		if err == nil {
			t.Errorf("Error: %s", err.Error())
		}
		if err.Error() != "findbyid error" {
			t.Errorf("Error: %s", err.Error())
		}
	})
	// get findbynick error
	t.Run("should return error when findbynick error", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbynickerror"
		output := &InsertOutputDtoMock{}
		param := "nickname"
		service := NewGet(log, client, param, output)
		err := service.Execute()
		if err == nil {
			t.Errorf("Error: %s", err.Error())
		}
		if err.Error() != "findbynick error" {
			t.Errorf("Error: %s", err.Error())
		}
	})
	// get findbyemail error
	t.Run("should return error when findbyemail error", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbyemailerror"
		output := &InsertOutputDtoMock{}
		param := "email"
		service := NewGet(log, client, param, output)
		err := service.Execute()
		if err == nil {
			t.Errorf("Error: %s", err.Error())
		}
		if err.Error() != "findbyemail error" {
			t.Errorf("Error: %s", err.Error())
		}
	})
	// get findbydoc error
	t.Run("should return error when findbydoc error", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbydocerror"
		output := &InsertOutputDtoMock{}
		param := "12345678"
		service := NewGet(log, client, param, output)
		err := service.Execute()
		if err == nil {
			t.Errorf("Error: %s", err.Error())
		}
		if err.Error() != "findbydoc error" {
			t.Errorf("Error: %s", err.Error())
		}
	})
	// get findbyphone error
	t.Run("should return error when findbyphone error", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientMock{}
		client.Status = "findbyphoneerror"
		output := &InsertOutputDtoMock{}
		param := "12345678"
		service := NewGet(log, client, param, output)
		err := service.Execute()
		if err == nil {
			t.Errorf("Error: %s", err.Error())
		}
		if err.Error() != "findbyphone error" {
			t.Errorf("Error: %s", err.Error())
		}
	})
}
