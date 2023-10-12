package service

import (
	"strings"
	"testing"
)

func TestUpdateExecute(t *testing.T) {
	t.Run("should update ok", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "findbyid"
		input := UpdateInputDtoMock{}
		input.Status = "ok"
		output := UpdateOutputDtoMock{}
		id := "957134b5-8de1-4121-80e0-275bb16e1b11"
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if log.mtype != "Info" {
			t.Errorf("Expected Info, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "updated") {
			t.Errorf("Expected 'created', got '%s'", log.msg)
		}
	})
	t.Run("should generate blank id error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "findbyid"
		input := UpdateInputDtoMock{}
		input.Status = "ok"
		output := UpdateOutputDtoMock{}
		id := ""
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error: %s", err)
		}
		if log.mtype != "Info" {
			t.Errorf("Expected Info, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "bad request") {
			t.Errorf("Expected 'bad request', got '%s'", log.msg)
		}
	})
	t.Run("should generate invalid id error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "findbyid"
		input := UpdateInputDtoMock{}
		input.Status = "ok"
		output := UpdateOutputDtoMock{}
		id := "invalid"
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error: %s", err)
		}
		if log.mtype != "Info" {
			t.Errorf("Expected Info, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "bad request") {
			t.Errorf("Expected 'bad request', got '%s'", log.msg)
		}
	})
	t.Run("should generate blank input error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "findbyid"
		input := UpdateInputDtoMock{}
		input.Status = "blank"
		output := UpdateOutputDtoMock{}
		id := "957134b5-8de1-4121-80e0-275bb16e1b11"
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error: %s", err)
		}
		if log.mtype != "Info" {
			t.Errorf("Expected Info, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "bad request") {
			t.Errorf("Expected 'bad request', got '%s'", log.msg)
		}
	})
	t.Run("should generate invalid input error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "findbyid"
		input := UpdateInputDtoMock{}
		input.Status = "invalid"
		output := UpdateOutputDtoMock{}
		id := "957134b5-8de1-4121-80e0-275bb16e1b11"
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error: %s", err)
		}
		if log.mtype != "Info" {
			t.Errorf("Expected Info, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "bad request") {
			t.Errorf("Expected 'bad request', got '%s'", log.msg)
		}
	})
	t.Run("should generate not found error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "notfound"
		input := UpdateInputDtoMock{}
		input.Status = "ok"
		output := UpdateOutputDtoMock{}
		id := "957134b5-8de1-4121-80e0-275bb16e1b11"
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error: %s", err)
		}
		if log.mtype != "Info" {
			t.Errorf("Expected Info, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "not found") {
			t.Errorf("Expected 'not found', got '%s'", log.msg)
		}
	})
	t.Run("should generate format input error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "ok"
		input := UpdateInputDtoMock{}
		input.Status = "formaterror"
		output := UpdateOutputDtoMock{}
		id := "957134b5-8de1-4121-80e0-275bb16e1b11"
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error: %s", err)
		}
		if !strings.Contains(err.Error(), "internal") {
			t.Errorf("Expected 'internal error', got '%s'", err.Error())
		}
	})
	t.Run("should generate load by id error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "findbyiderror"
		input := UpdateInputDtoMock{}
		input.Status = "ok"
		output := UpdateOutputDtoMock{}
		id := "957134b5-8de1-4121-80e0-275bb16e1b11"
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error: %s", err)
		}
		if !strings.Contains(err.Error(), "internal") {
			t.Errorf("Expected 'internal error', got '%s'", err.Error())
		}
	})
	t.Run("should generate document duplicity error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "duplicitydocument"
		input := UpdateInputDtoMock{}
		input.Status = "ok"
		output := UpdateOutputDtoMock{}
		id := "957134b5-8de1-4121-80e0-275bb16e1b11"
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error: %s", err)
		}
		if log.mtype != "Info" {
			t.Errorf("Expected Info, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "conflict") {
			t.Errorf("Expected 'duplicity', got '%s'", log.msg)
		}
	})

	t.Run("should generate email duplicity error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "duplicityemail"
		input := UpdateInputDtoMock{}
		input.Status = "ok"
		output := UpdateOutputDtoMock{}
		id := "957134b5-8de1-4121-80e0-275bb16e1b11"
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error: %s", err)
		}
		if log.mtype != "Info" {
			t.Errorf("Expected Info, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "conflict") {
			t.Errorf("Expected 'duplicity', got '%s'", log.msg)
		}
	})
	t.Run("should generate nick duplicity error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "duplicitynick"
		input := UpdateInputDtoMock{}
		input.Status = "ok"
		output := UpdateOutputDtoMock{}
		id := "957134b5-8de1-4121-80e0-275bb16e1b11"
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error: %s", err)
		}
		if log.mtype != "Info" {
			t.Errorf("Expected Info, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "conflict") {
			t.Errorf("Expected 'duplicity', got '%s'", log.msg)
		}
	})

	t.Run("should generate document duplicity internal error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "duplicitydocumenterror"
		input := UpdateInputDtoMock{}
		input.Status = "ok"
		output := UpdateOutputDtoMock{}
		id := "957134b5-8de1-4121-80e0-275bb16e1b11"
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error: %s", err)
		}
		if log.mtype != "Error" {
			t.Errorf("Expected Info, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "error") {
			t.Errorf("Expected 'duplicity', got '%s'", log.msg)
		}
	})
	t.Run("should generate email duplicity internal error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "duplicityemailerror"
		input := UpdateInputDtoMock{}
		input.Status = "ok"
		output := UpdateOutputDtoMock{}
		id := "957134b5-8de1-4121-80e0-275bb16e1b11"
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error: %s", err)
		}
		if log.mtype != "Error" {
			t.Errorf("Expected Info, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "error") {
			t.Errorf("Expected 'duplicity', got '%s'", log.msg)
		}
	})
	t.Run("should generate nick duplicity internal error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "duplicitynickerror"
		input := UpdateInputDtoMock{}
		input.Status = "ok"
		output := UpdateOutputDtoMock{}
		id := "957134b5-8de1-4121-80e0-275bb16e1b11"
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Error: %s", err)
		}
		if log.mtype != "Error" {
			t.Errorf("Expected Info, got %s", log.mtype)
		}
		if !strings.Contains(log.msg, "error") {
			t.Errorf("Expected 'duplicity', got '%s'", log.msg)
		}
	})
	t.Run("should generate update error", func(t *testing.T) {
		log := LogMock{}
		client := ClientMock{}
		client.Status = "updateerror"
		input := UpdateInputDtoMock{}
		input.Status = "ok"
		output := UpdateOutputDtoMock{}
		id := "957134b5-8de1-4121-80e0-275bb16e1b11"
		s := NewUpdate(&log, &client, id, &input, &output)
		err := s.Execute()
		if err == nil {
			t.Errorf("Expected error, got %s", err)
		}
		if !strings.Contains(err.Error(), "internal") {
			t.Errorf("Expected 'internal error', got '%s'", err.Error())
		}
	})

}
