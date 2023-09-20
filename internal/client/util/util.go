package util

import (
	"regexp"
	"github.com/lavinas/keel/pkg/cpf_cnpj"
)

type Util struct {
}

func NewUtil() *Util {
	return &Util{}
}

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

func (d *Util) ClearNumber(document string) string {
	re := regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(document, "")
}
