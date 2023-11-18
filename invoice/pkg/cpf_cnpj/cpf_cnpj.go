package cpf_cnpj

import (
	"regexp"
)

// IsValid returns if a string is a valid CPF or CNPJ document
func ValidateCPFOrCNPJ(s string) bool {
	re := regexp.MustCompile(`[^0-9]`)
	doc := re.ReplaceAllString(s, "")
	if doc == "" {
		return false
	}
	if ValidateCPF(doc) {
		return true
	}
	if ValidateCNPJ(doc) {
		return true
	}
	return false
}
