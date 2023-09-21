package util

import (
	"testing"
)

func TestBlankName(t *testing.T) {
	d := Util{}
	b, m := d.ValidateName(" ")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "name is blank" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestUncompleteName(t *testing.T) {
	d := Util{}
	b, m := d.ValidateName("Teste")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "name should have at least two parts" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestValidName(t *testing.T) {
	d := Util{}
	b, m := d.ValidateName("Teste xxx")
	if !b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestBlankNick(t *testing.T) {
	d := Util{}
	b, m := d.ValidateNickname("")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "nickname is blank" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestInvalidNick(t *testing.T) {
	d := Util{}
	b, m := d.ValidateNickname("jose agora")
	if !b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestValidNick(t *testing.T) {
	d := Util{}
	b, m := d.ValidateNickname(" teste_abril ")
	if !b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestBlankDocument(t *testing.T) {
	d := Util{}
	b, m := d.ValidateDocument("")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "invalid document" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestBlankAlphaDocument(t *testing.T) {
	d := Util{}
	b, m := d.ValidateDocument("abcdefX*&^")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "invalid document" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestValidateDocumentWithValidCPF(t *testing.T) {
	d := Util{}
	b, m := d.ValidateDocument("206. 656.600-49 ")
	if !b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestValidateDocumentWithValidCPF2(t *testing.T) {
	d := Util{}
	b, m := d.ValidateDocument("206656-60049")
	if !b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestValidateWithInvalidCPF(t *testing.T) {
	d := Util{}
	b, m := d.ValidateDocument("206.656.600-50")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m == "" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestValidateWithValidCNPJ(t *testing.T) {
	d := Util{}
	b, m := d.ValidateDocument("53.931.154/0001-63")
	if !b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestValidateWithInvalidCNPJ(t *testing.T) {
	d := Util{}
	b, m := d.ValidateDocument("53.931.154/0001-62")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m == "" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestBlankEmail(t *testing.T) {
	d := Util{}
	b, m := d.ValidateEmail("")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "invalid email" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestValidEmail(t *testing.T){
	d := Util{}
	b, m := d.ValidateEmail(" test@test.com.br ")
	if !b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestInvalidEmail(t *testing.T){
	d := Util{}
	b, m := d.ValidateEmail("test")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "invalid email" {
		t.Errorf("Invalid result: %v", m)
	}
	b, m = d.ValidateEmail("test @ mail.com")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "invalid email" {
		t.Errorf("Invalid result: %v", m)
	}
	b, m = d.ValidateEmail("test @mail")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "invalid email" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestBlankPhone(t *testing.T) {
	d := Util{}
	b, m := d.ValidatePhone("")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "invalid cell phone" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestValidPhone(t *testing.T) {
	d := Util{}
	b, m := d.ValidatePhone("(013)9-9999-9999")
	if !b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "" {
		t.Errorf("Invalid result: %v", m)
	}
	b, m = d.ValidatePhone("13999999999")
	if !b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestInvalidPhone(t *testing.T) {
	d := Util{}
	b, m := d.ValidatePhone("9-9999-9999")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "invalid cell phone" {
		t.Errorf("Invalid result: %v", m)
	}
	b, m = d.ValidatePhone("1122671111")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "invalid cell phone" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestValidAll(t *testing.T) {
	d := Util{}
	b, m := d.ValidateAll("Teste teste", "teste", "206.656.600-49", "13999999999", "teste@test.com.br")
	if !b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestInvalidAll(t *testing.T) {
	d := Util{}
	b, m := d.ValidateAll("", "teste", "206.656.600-49", "13999999999", "teste@test.com.br")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "name is blank" {
		t.Errorf("Invalid result: %v", m)
	}
	b, m = d.ValidateAll("", "teste", "206.656.600-49", "13999999999", "teste")
	if b {	
		t.Errorf("Invalid result: %v", b)
	}
	if m != "name is blank || invalid email" {
		t.Errorf("Invalid result: %v", m)
	}
}

func TestClearName(t *testing.T) {
	d := Util{}
	r, err := d.ClearName("   Jose   da Silva  ")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "Jose Da Silva" {
		t.Errorf("Invalid result: %v", r)
	}
	r, err = d.ClearName("   joSe   da Silva  ")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "Jose Da Silva" {
		t.Errorf("Invalid result: %v", r)
	}
	r, err = d.ClearName("    ")
	if err.Error() != "name is blank" {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "" {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestClearNickname(t *testing.T) {
	d := Util{}
	r, err := d.ClearNickname("   Jose   da Silva  222")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "jose_da_silva_222" {
		t.Errorf("Invalid result: %v", r)
	}
	r, err = d.ClearNickname("    ")
	if err.Error() != "nickname is blank" {
		t.Errorf("Invalid result: %v", err)
	}
	if r != "" {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestClearDocument(t *testing.T) {
	d := Util{}
	r, err := d.ClearDocument("206.  656.600-49")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != 20665660049 {
		t.Errorf("Invalid result: %v", r)
	}
	if r, err := d.ClearDocument(""); err.Error() != "invalid document" {
		t.Errorf("Invalid result: %v", r)
	}
	if r, err := d.ClearDocument("dasdasdsa--asdasd"); err.Error() != "invalid document" {
		t.Errorf("Invalid result: %v", r)
	}
	if _, err := d.ClearDocument("206.656.600-50"); err.Error() != "invalid document" {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestClearPhone(t *testing.T) {
	d := Util{}
	r, err := d.ClearPhone("(013) 9-9999-9999")
	if err != nil {
		t.Errorf("Invalid result: %v", err)
	}
	if r != 5513999999999 {
		t.Errorf("Invalid result: %v", r)
	}
	if _, err := d.ClearPhone(""); err.Error() != "invalid cell phone" {
		t.Errorf("Invalid result: %v", err)
	}
}

