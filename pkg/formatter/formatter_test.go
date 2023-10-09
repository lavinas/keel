package formatter

import (
	"testing"
)

func TestClearName(t *testing.T) {
	r, err := FormatName("   Jose   da Silva  ")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "Jose Da Silva" {
		t.Errorf("Invalid result: %v", r)
	}
	r, err = FormatName("   joSe   da Silva  ")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "Jose Da Silva" {
		t.Errorf("Invalid result: %v", r)
	}
	r, err = FormatName("    ")
	if err.Error() != "name is blank" {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "" {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestClearNickname(t *testing.T) {
	r, err := FormatNickname("   Jose   da Silva  222")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "jose_da_silva_222" {
		t.Errorf("Invalid result: %v", r)
	}
	r, err = FormatNickname("    ")
	if err.Error() != "nickname is blank" {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "" {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestClearDocument(t *testing.T) {
	r, err := FormatDocument("206.  656.600-49")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "20665660049" {
		t.Errorf("Invalid result: %v", r)
	}

	r, err = FormatDocument("044. 179328-24")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "04417932824" {
		t.Errorf("Invalid result: %v", r)
	}

	if r, err := FormatDocument(""); err.Error() != "document is blank" {
		t.Errorf("Invalid result: %v", r)
	}
	if r, err := FormatDocument("dasdasdsa--asdasd"); err.Error() != "invalid document" {
		t.Errorf("Invalid result: %v", r)
	}
	if _, err := FormatDocument("206.656.600-50"); err.Error() != "invalid document" {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestClearPhone(t *testing.T) {
	r, err := FormatPhone("(013) 9-9999-9999")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "5513999999999" {
		t.Errorf("Invalid result: %v", r)
	}
	if _, err := FormatPhone(""); err.Error() != "phone is blank" {
		t.Errorf("Invalid result: %v", err)
	}
}
