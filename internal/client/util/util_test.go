package util

import (
	"testing"
)

func TestClearDocument(t *testing.T) {
	d := Util{}
	r := d.ClearNumber("206.656.600-49")

	if r != "20665660049" {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidateDocumentWithValidCPF(t *testing.T) {
	d := Util{}
	r := d.ValidateDocument("206.656.600-49")

	if !r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidateDocumentWithValidCPF2(t *testing.T) {
	d := Util{}
	r := d.ValidateDocument("20665660049")

	if !r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidateWithInvalidCPF(t *testing.T) {
	d := Util{}
	r := d.ValidateDocument("206.656.600-50")

	if r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidateWithValidCNPJ(t *testing.T) {
	d := Util{}
	r := d.ValidateDocument("53.931.154/0001-63")

	if !r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidateWithInvalidCNPJ(t *testing.T) {
	d := Util{}
	r := d.ValidateDocument("53.931.154/0001-62")

	if r {
		t.Errorf("Invalid result: %v", r)
	}
}
