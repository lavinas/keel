package service

import (
	"testing"
)

func TestServiceInsert(t *testing.T){
	t.Run("should insert", func(t *testing.T) {
		domain := &DomainMock{}
		config := &ConfigMock{}
		log := &LogMock{}
		repo := &RepoMock{}
		input := &InsertInputDtoMock{}
		input.Status = "ok"
		output := &InsertOutputDtoMock{}
		service := NewService(domain, config, log, repo)
		err := service.Insert(input, output)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	t.Run("should not insert", func(t *testing.T) {
		domain := &DomainMock{}
		config := &ConfigMock{}
		log := &LogMock{}
		repo := &RepoMock{}
		input := &InsertInputDtoMock{}
		input.Status = "invalid"
		output := &InsertOutputDtoMock{}
		service := NewService(domain, config, log, repo)
		err := service.Insert(input, output)
		if err == nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
}

func TestServiceFind(t *testing.T){
	t.Run("should find", func(t *testing.T) {
		domain := &DomainMock{}
		config := &ConfigMock{}
		log := &LogMock{}
		repo := &RepoMock{}
		input := &FindInputDtoMock{}
		input.Status = "ok"
		output := &FindOutputDtoMock{}
		service := NewService(domain, config, log, repo)
		err := service.Find(input, output)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	t.Run("should not find", func(t *testing.T) {
		domain := &DomainMock{}
		config := &ConfigMock{}
		log := &LogMock{}
		repo := &RepoMock{}
		input := &FindInputDtoMock{}
		input.Status = "invalid"
		output := &FindOutputDtoMock{}
		service := NewService(domain, config, log, repo)
		err := service.Find(input, output)
		if err == nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
}

func TestServiceUpdate(t *testing.T){
	t.Run("should update", func(t *testing.T) {
		domain := &DomainMock{}
		domain.Status = "findbyid"
		config := &ConfigMock{}
		log := &LogMock{}
		repo := &RepoMock{}
		input := &UpdateInputDtoMock{}
		input.Status = "ok"
		output := &UpdateOutputDtoMock{}
		service := NewService(domain, config, log, repo)
		err := service.Update("1", input, output)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
	t.Run("should not update", func(t *testing.T) {
		domain := &DomainMock{}
		config := &ConfigMock{}
		log := &LogMock{}
		repo := &RepoMock{}
		input := &UpdateInputDtoMock{}
		input.Status = "invalid"
		output := &UpdateOutputDtoMock{}
		service := NewService(domain, config, log, repo)
		err := service.Update("1", input, output)
		if err == nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
}