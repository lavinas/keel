package service

import (
	_ "reflect"
	"strings"
	"testing"

)

func TestInsertExecute(t *testing.T) {
	t.Run("should insert", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "ok"
		input := InsertInputDtoMock{}
		input.Status = "ok"
		output := InsertOutputDtoMock{}
		output.Status = "ok"
		s := NewInsert(&log, &client, &input, &output)
		err := s.Execute()
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if log.mtype != "Info" {
			t.Errorf("Expected Info, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "created") {
			t.Errorf("Expected 'created', got '%s'", log.msg)
		}
	})
}