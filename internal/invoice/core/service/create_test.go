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
	})
}
