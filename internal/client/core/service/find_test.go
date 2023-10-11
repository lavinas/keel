package service

import (
	"testing"
)

func TestFindExecute(t *testing.T) {
	t.Run("should find clients", func(t *testing.T) {
		config := &ConfigMock{}
		log := &LogMock{}
		client := &ClientSetMock{}
		input := &FindInputDtoMock{}
		output := &FindOutputDtoMock{}
		service := NewFind(config, log, client, input, output)
		err := service.Execute()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
}