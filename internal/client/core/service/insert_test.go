package service

import (
	"fmt"
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
	// should validate error
	t.Run("should validate error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "ok"
		input := InsertInputDtoMock{}
		input.Status = "invalid"
		output := InsertOutputDtoMock{}
		s := NewInsert(&log, &client, &input, &output)
		fmt.Println(s)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if log.mtype != "Info" {
			t.Errorf("Expected Info, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "bad request") {
			t.Errorf("Expected 'bad request', got '%s'", log.msg)
		}
	})
	// should not load client
	t.Run("should not load client", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "loaderror"
		input := InsertInputDtoMock{}
		input.Status = "ok"
		output := InsertOutputDtoMock{}
		s := NewInsert(&log, &client, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if log.mtype != "Error" {
			t.Errorf("Expected Error, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "error") {
			t.Errorf("Expected 'error', got '%s'", log.msg)
		}
	})
	// should document duplicity check error
	t.Run("should document duplicity check error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "duplicitydocumenterror"
		input := InsertInputDtoMock{}
		input.Status = "ok"
		output := InsertOutputDtoMock{}
		s := NewInsert(&log, &client, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if log.mtype != "Error" {
			t.Errorf("Expected Error, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "error") {
			t.Errorf("Expected 'internal error', got '%s'", log.msg)
		}
	})

	// should document duplicity message
	t.Run("should document duplicity message", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "duplicitydocument"
		input := InsertInputDtoMock{}
		input.Status = "ok"
		output := InsertOutputDtoMock{}
		s := NewInsert(&log, &client, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if log.mtype != "Info" {
			t.Errorf("Expected Error, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "conflict") {
			t.Errorf("Expected 'error', got '%s'", log.msg)
		}
	})
	// should email duplicity check error
	t.Run("should email duplicity check error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "duplicityemailerror"
		input := InsertInputDtoMock{}
		input.Status = "ok"
		output := InsertOutputDtoMock{}
		s := NewInsert(&log, &client, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if log.mtype != "Error" {
			t.Errorf("Expected Error, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "error") {
			t.Errorf("Expected 'error', got '%s'", log.msg)
		}
	})

	// should document duplicity message
	t.Run("should email duplicity message", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "duplicityemail"
		input := InsertInputDtoMock{}
		input.Status = "ok"
		output := InsertOutputDtoMock{}
		s := NewInsert(&log, &client, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if log.mtype != "Info" {
			t.Errorf("Expected Error, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "conflict") {
			t.Errorf("Expected 'error', got '%s'", log.msg)
		}
	})
	// should nick duplicity check error
	t.Run("should nick duplicity check error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "duplicitynickerror"
		input := InsertInputDtoMock{}
		input.Status = "ok"
		output := InsertOutputDtoMock{}
		s := NewInsert(&log, &client, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if log.mtype != "Error" {
			t.Errorf("Expected Error, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "error") {
			t.Errorf("Expected 'error', got '%s'", log.msg)
		}
	})
	// should nick duplicity message
	t.Run("should nick duplicity message", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "duplicitynick"
		input := InsertInputDtoMock{}
		input.Status = "ok"
		output := InsertOutputDtoMock{}
		s := NewInsert(&log, &client, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if log.mtype != "Info" {
			t.Errorf("Expected Error, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "conflict") {
			t.Errorf("Expected 'error', got '%s'", log.msg)
		}
	})
	// should save error message
	t.Run("should save error message", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "saveerror"
		input := InsertInputDtoMock{}
		input.Status = "ok"
		output := InsertOutputDtoMock{}
		s := NewInsert(&log, &client, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error should not be nil")
		}
		if log.mtype != "Error" {
			t.Errorf("Expected Error, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "error") {
			t.Errorf("Expected 'error', got '%s'", log.msg)
		}
	})



}