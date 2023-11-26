package cpf_cnpj

import (
	"errors"
	"regexp"
	"strconv"
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

// ParseUint returns a uint64 from a cpf or cnpj string
func ParseUint(s string) (uint64, error) {
	re := regexp.MustCompile(`[^0-9]`)
	doc := re.ReplaceAllString(s, "")
	if doc == "" {
		return 0, errors.New("invalid document")
	}
	i, err := strconv.ParseUint(doc, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}
