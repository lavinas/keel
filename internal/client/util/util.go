package util

import (
	"net/mail"
	"regexp"

	"github.com/lavinas/keel/pkg/cpf_cnpj"
)

// Util is the util functions to be used in the client services
type Util struct {
}

// NewUtil creates a new util
func NewUtil() *Util {
	return &Util{}
}

// ValidateDocument validates a document
func (d *Util) ValidateDocument(document string) bool {
	doc := d.ClearNumber(document)
	if cpf_cnpj.ValidateCPF(doc) {
		return true
	}
	if cpf_cnpj.ValidateCNPJ(doc) {
		return true
	}
	return false
}

// ValidateEmail validates an email address
func (d *Util) ValidateEmail(email string) bool {
	if e := d.ClearEmail(email); e != "" {
		return true
	}
	return false
}

// ClearNumber clears a number
func (d *Util) ClearNumber(document string) string {
	re := regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(document, "")
}

// ClearEmail clears an email address
func (d *Util) ClearEmail(email string) string {
	e, err := mail.ParseAddress(email)
	if err != nil {
		return ""
	}
	return e.Address
}
