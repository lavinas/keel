package formatter

import (
	"testing"
)

func TestFormatName(t *testing.T) {
	t.Run("should format normal name", func(t *testing.T) {
		r, err := FormatName("   Jose   da Silva  ")
		if err != nil {
			t.Errorf("Invalid result: %v", err)
		}
		if r != "Jose Da Silva" {
			t.Errorf("Invalid result: %v", r)
		}
	})
	t.Run("should format name with spaces", func(t *testing.T) {
		r, err := FormatName("   joSe   da Silva  ")
		if err != nil {
			t.Errorf("Invalid result: %v", err)
		}
		if r != "Jose Da Silva" {
			t.Errorf("Invalid result: %v", r)
		}
	})
	t.Run("should return name is blank error", func(t *testing.T) {
		r, err := FormatName("    ")
		if err.Error() != "name is blank" {
			t.Errorf("Invalid result: %v", err)
		}
		if r != "" {
			t.Errorf("Invalid result: %v", r)
		}
	})
	t.Run("should return name should have at least two parts error", func(t *testing.T) {
		r, err := FormatName("test")
		if err.Error() != "name should have at least two parts" {
			t.Errorf("Invalid result: %v", err)
		}
		if r != "" {
			t.Errorf("Invalid result: %v", r)
		}
	})
}

func TestFormatNickname(t *testing.T) {
	t.Run("should format spaced nickname", func(t *testing.T) {
		r, err := FormatNickname("   Jose   da Silva  222")
		if err != nil {
			t.Errorf("Invalid result: %v", err)
		}
		if r != "jose_da_silva_222" {
			t.Errorf("Invalid result: %v", r)
		}
	})
	t.Run("should return nickname is blank error", func(t *testing.T) {
		r, err := FormatNickname("    ")
		if err.Error() != "nickname is blank" {
			t.Errorf("Invalid result: %v", err)
		}
		if r != "" {
			t.Errorf("Invalid result: %v", r)
		}
	})
}

func TestFormatDocument(t *testing.T) {
	t.Run("should format spaced cpf", func(t *testing.T) {
		r, err := FormatDocument("206.  656.600-49")
		if err != nil {
			t.Errorf("Invalid result: %v", err)
		}
		if r != "20665660049" {
			t.Errorf("Invalid result: %v", r)
		}
	})
	t.Run("should format spaced cnpj", func(t *testing.T) {
		r, err := FormatDocument("28.183.859/0001-00")
		if err != nil {
			t.Errorf("Invalid result: %v", err)
		}
		if r != "28183859000100" {
			t.Errorf("Invalid result: %v", r)
		}
	})
	t.Run("should format cpf with wrong format", func(t *testing.T) {
		r, err := FormatDocument("044. 179328-24")
		if err != nil {
			t.Errorf("Invalid result: %v", err)
		}
		if r != "04417932824" {
			t.Errorf("Invalid result: %v", r)
		}
	})
	t.Run("should return document is blank error", func(t *testing.T) {
		if r, err := FormatDocument(""); err.Error() != "document is blank" {
			t.Errorf("Invalid result: %v", r)
		}
	})
	t.Run("should return invalid document error when no numeric input", func(t *testing.T) {
		if r, err := FormatDocument("dasdasdsa--asdasd"); err.Error() != "invalid document" {
			t.Errorf("Invalid result: %v", r)
		}
	})
	t.Run("should return invalid document error when cpf is invalid", func(t *testing.T) {
		if r, err := FormatDocument("206.656.600-50"); err.Error() != "invalid document" {
			t.Errorf("Invalid result: %v", r)
		}
	})
	t.Run("should retirn invalid document error when cnpj is invalid", func(t *testing.T) {
		if r, err := FormatDocument("04.417.932/0001-00"); err.Error() != "invalid document" {
			t.Errorf("Invalid result: %v", r)
		}
	})
}

func TestFormatPhone(t *testing.T) {
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

func TestFormatEmail(t *testing.T) {
	r, err := FormatEmail("test@test.com")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "test@test.com" {
		t.Errorf("Invalid result: %v", r)
	}
	if _, err := FormatEmail(""); err.Error() != "email is blank" {
		t.Errorf("Invalid result: %v", err)
	}
	if _, err := FormatEmail("test"); err.Error() != "invalid email" {
		t.Errorf("Invalid result: %v", err)
	}
}
