package service

import (
	"os"
	"testing"
)

func TestFindExecute(t *testing.T) {
	// ok
	t.Run("should find clients", func(t *testing.T) {
		pp := os.Getenv(per_page)
		os.Setenv(per_page, "10")
		defer func() {
			os.Setenv(per_page, pp)
		}()
		log := &LogMock{}
		client := &ClientSetMock{}
		input := &FindInputDtoMock{}
		input.Status = "ok"
		output := &FindOutputDtoMock{}
		service := NewFind(log, client)
		err := service.Execute(input, output)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	// invalid input
	t.Run("should return error when input is invalid", func(t *testing.T) {
		pp := os.Getenv(per_page)
		os.Setenv(per_page, "10")
		defer func() {
			os.Setenv(per_page, pp)
		}()
		log := &LogMock{}
		client := &ClientSetMock{}
		client.Status = "ok"
		input := &FindInputDtoMock{}
		input.Status = "invalid"
		output := &FindOutputDtoMock{}
		service := NewFind(log, client)
		err := service.Execute(input, output)
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if err != nil && err.Error() != "bad request: invalid input" {
			t.Errorf("Error: %s", err.Error())
		}
	})
	// internal error
	t.Run("should return error when internal error", func(t *testing.T) {
		pp := os.Getenv(per_page)
		os.Setenv(per_page, "10")
		defer func() {
			os.Setenv(per_page, pp)
		}()
		log := &LogMock{}
		client := &ClientSetMock{}
		client.Status = "internal"
		input := &FindInputDtoMock{}
		input.Status = "internal"
		output := &FindOutputDtoMock{}
		service := NewFind(log, client)
		err := service.Execute(input, output)
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
		pp := os.Getenv(per_page)
		os.Setenv(per_page, "")
		defer func() {
			os.Setenv(per_page, pp)
		}()
		log := &LogMock{}
		client := &ClientSetMock{}
		input := &FindInputDtoMock{}
		input.Status = "blank"
		service := NewFind(log, client)
		page, perPage, _, _, _, _, _ := service.getAll(input)
		if page != 1 {
			t.Errorf("invalid page. Expected: 1, got: %d", page)
		}
		if perPage != 10 {
			t.Errorf("invalid perPage. Expected: 10, got: %d", perPage)
		}
	})
	t.Run("should get invalid perPage number", func(t *testing.T) {
		pp := os.Getenv(per_page)
		os.Setenv(per_page, "")
		defer func() {
			os.Setenv(per_page, pp)
		}()
		log := &LogMock{}
		client := &ClientSetMock{}
		input := &FindInputDtoMock{}
		input.Status = "invalid"
		service := NewFind(log, client)
		_, perPage, _, _, _, _, _ := service.getAll(input)
		if perPage != 10 {
			t.Errorf("invalid perPage. Expected: 10, got: %d", perPage)
		}
	})
}
