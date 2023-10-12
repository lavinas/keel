package service

import (
	"testing"
)

func TestFindExecute(t *testing.T) {
	// ok
	t.Run("should find clients", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientSetMock{}
		input := &FindInputDtoMock{}
		input.Status = "ok"
		output := &FindOutputDtoMock{}
		service := NewFind(config, log, client, input, output)
		err := service.Execute()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	// invalid input
	t.Run("should return error when input is invalid", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientSetMock{}
		client.Status = "ok"
		input := &FindInputDtoMock{}
		input.Status = "invalid"
		output := &FindOutputDtoMock{}
		service := NewFind(config, log, client, input, output)
		err := service.Execute()
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if err != nil && err.Error() != "bad request: invalid input" {
			t.Errorf("Error: %s", err.Error())
		}
	})
	// internal error
	t.Run("should return error when internal error", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "ok"
		log := &LogMock{}
		client := &ClientSetMock{}
		client.Status = "internal"
		input := &FindInputDtoMock{}
		input.Status = "internal"
		output := &FindOutputDtoMock{}
		service := NewFind(config, log, client, input, output)
		err := service.Execute()
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if err != nil && err.Error() != "internal error" {
			t.Errorf("Error: %s", err.Error())
		}
	})
}

func TestFindGetAll(t *testing.T) {
	t.Run("should get perPage config error", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "invalid"
		log := &LogMock{}
		client := &ClientSetMock{}
		input := &FindInputDtoMock{}
		input.Status = "blank"
		output := &FindOutputDtoMock{}
		service := NewFind(config, log, client, input, output)
		page, perPage, _, _, _, _, _ := service.getAll()
		if page != 1 {
			t.Errorf("invalid page. Expected: 1, got: %d", page)
		}
		if perPage != 10 {
			t.Errorf("invalid perPage. Expected: 10, got: %d", perPage)
		}
	})
	t.Run("should get invalid perPage number", func(t *testing.T) {
		config := &ConfigMock{}
		config.Status = "invalid"
		log := &LogMock{}
		client := &ClientSetMock{}
		input := &FindInputDtoMock{}
		input.Status = "invalid"
		output := &FindOutputDtoMock{}
		service := NewFind(config, log, client, input, output)
		_, perPage, _, _, _, _, _ := service.getAll()
		if perPage != 10 {
			t.Errorf("invalid perPage. Expected: 10, got: %d", perPage)
		}
	})
}
