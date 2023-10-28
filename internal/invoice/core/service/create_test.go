package service

import (
	"testing"
)

func TestCreateExecute(t *testing.T) {
	t.Run("should create without errors", func(t *testing.T) {
		l := LogMock{}
		d := InvoiceMock{}
		i := CreateInputDtoMock{}
		o := CreateOutputDtoMock{}
		c := NewCreate(&l, &d, &i, &o)
		err := c.Execute()
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
