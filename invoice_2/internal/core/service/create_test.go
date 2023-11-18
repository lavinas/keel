package service

import (
	"strings"
	"testing"
)

func TestCreateExecute(t *testing.T) {
	t.Run("should create without errors", func(t *testing.T) {
		p := RepoMock{}
		l := LogMock{}
		r := RestConsumerMock{}
		d := InvoiceMock{}
		i := CreateInputDtoMock{}
		o := CreateOutputDtoMock{}
		c := NewCreate(&p, &l, &r, &d)
		err := c.Execute(&i, &o)
		if err != nil {
			t.Errorf("expected no errors, got %v", err)
		}
		if o.status != "created" {
			t.Errorf("expected status to be created, got %v", o.status)
		}
		if o.reference != "" {
			t.Errorf("expected reference to be empty, got %v", o.reference)
		}
		if !strings.Contains(l.Message, "created") {
			t.Errorf("expected log message to be created, got %v", l.Message)
		}
		if l.Type != "info" {
			t.Errorf("expected log type to be info, got %v", l.Type)
		}
	})
	t.Run("should return error when input is invalid", func(t *testing.T) {
		p := RepoMock{}
		l := LogMock{}
		r := RestConsumerMock{}
		d := InvoiceMock{}
		i := CreateInputDtoMock{Status: "format error"}
		o := CreateOutputDtoMock{}
		c := NewCreate(&p, &l, &r, &d)
		err := c.Execute(&i, &o)
		if err == nil {
			t.Errorf("expected errors, got %v", err)
		}
		if !strings.Contains(o.status, "bad request") {
			t.Errorf("expected status to be bad request, got %v", o.status)
		}
		if o.reference != "" {
			t.Errorf("expected reference to be empty, got %v", o.reference)
		}
	})
	t.Run("should return error when domain fails to load", func(t *testing.T) {
		p := RepoMock{}
		l := LogMock{}
		r := RestConsumerMock{}
		d := InvoiceMock{Status: "load error"}
		i := CreateInputDtoMock{}
		o := CreateOutputDtoMock{}
		c := NewCreate(&p, &l, &r, &d)
		err := c.Execute(&i, &o)
		if err == nil {
			t.Errorf("expected errors, got %v", err)
		}
		if !strings.Contains(o.status, "internal error") {
			t.Errorf("expected status to be internal error, got %v", o.status)
		}
		if o.reference != "" {
			t.Errorf("expected reference to be empty, got %v", o.reference)
		}
	})
	t.Run("should return error when domain fails to save", func(t *testing.T) {
		p := RepoMock{}
		l := LogMock{}
		r := RestConsumerMock{}
		d := InvoiceMock{Status: "save error"}
		i := CreateInputDtoMock{}
		o := CreateOutputDtoMock{}
		c := NewCreate(&p, &l, &r, &d)
		err := c.Execute(&i, &o)
		if err == nil {
			t.Errorf("expected errors, got %v", err)
		}
		if !strings.Contains(o.status, "internal error") {
			t.Errorf("expected status to be internal error, got %v", o.status)
		}
		if o.reference != "" {
			t.Errorf("expected reference to be empty, got %v", o.reference)
		}
	})
	t.Run("should return error when has duplicity error", func(t *testing.T) {
		p := RepoMock{}
		l := LogMock{}
		r := RestConsumerMock{}
		d := InvoiceMock{Status: "duplicity error"}
		i := CreateInputDtoMock{}
		o := CreateOutputDtoMock{}
		c := NewCreate(&p, &l, &r, &d)
		err := c.Execute(&i, &o)
		if err == nil {
			t.Errorf("expected errors, got %v", err)
		}
		if !strings.Contains(o.status, "internal error") {
			t.Errorf("expected bad request, got %v", o.status)
		}
		if o.reference != "" {
			t.Errorf("expected reference to be empty, got %v", o.reference)
		}
	})
	t.Run("should return error when has duplicity", func(t *testing.T) {
		p := RepoMock{}
		l := LogMock{}
		r := RestConsumerMock{}
		d := InvoiceMock{Status: "duplicity"}
		i := CreateInputDtoMock{}
		o := CreateOutputDtoMock{}
		c := NewCreate(&p, &l, &r, &d)
		err := c.Execute(&i, &o)
		if err == nil {
			t.Errorf("expected errors, got %v", err)
		}
		if !strings.Contains(o.status, "conflict") {
			t.Errorf("expected conflict, got %v", o.status)
		}
		if o.reference != "" {
			t.Errorf("expected reference to be empty, got %v", o.reference)
		}
	})
	t.Run("should return error when consumer fails to load", func(t *testing.T) {
		p := RepoMock{}
		l := LogMock{}
		r := RestConsumerMock{Status: "get client error"}
		d := InvoiceMock{}
		i := CreateInputDtoMock{}
		o := CreateOutputDtoMock{}
		c := NewCreate(&p, &l, &r, &d)
		err := c.Execute(&i, &o)
		if err == nil {
			t.Errorf("expected errors, got %v", err)
		}
		if !strings.Contains(o.status, "internal error") {
			t.Errorf("expected internal error, got %v", o.status)
		}
		if o.reference != "" {
			t.Errorf("expected reference to be empty, got %v", o.reference)
		}
	})
	t.Run("should return nil when consumer is not found", func(t *testing.T) {
		p := RepoMock{}
		l := LogMock{}
		r := RestConsumerMock{}
		r.Status = "get client not found"
		d := InvoiceMock{}
		i := CreateInputDtoMock{}
		o := CreateOutputDtoMock{}
		c := NewCreate(&p, &l, &r, &d)
		err := c.Execute(&i, &o)
		if err != nil {
			t.Errorf("expected no errors, got %v", err)
		}
		if o.status != "created" {
			t.Errorf("expected status to be created, got %v", o.status)
		}
		if o.reference != "" {
			t.Errorf("expected reference to be empty, got %v", o.reference)
		}
	})
}