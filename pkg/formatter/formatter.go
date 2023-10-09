package formatter

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strconv"
	"strings"

	"github.com/lavinas/keel/pkg/cpf_cnpj"
	"github.com/lavinas/keel/pkg/phone"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	// countries is the list of countries for phone number validation
	countries = []string{"BR"}
)

// clearName clears a name
func FormatName(name string) (string, error) {
	n := strings.Trim(name, " ")
	if n == "" {
		return "", errors.New("name is blank")
	}
	r := regexp.MustCompile(`\s+`)
	n = r.ReplaceAllString(n, " ")
	if i := strings.Split(n, " "); len(i) < 2 {
		return "", errors.New("name should have at least two parts")
	}
	n = strings.ToLower(n)
	n = cases.Title(language.BrazilianPortuguese).String(n)
	return n, nil
}

// clearNickname clears a nickname
func FormatNickname(nickname string) (string, error) {
	nick := strings.Trim(nickname, " ")
	if nick == "" {
		return "", errors.New("nickname is blank")
	}
	r := regexp.MustCompile(`\s+`)
	nick = r.ReplaceAllString(nick, " ")
	nick = strings.Replace(nick, " ", "_", -1)
	if n := strings.Split(nick, " "); len(n) != 1 {
		return "", errors.New("invalid nickname")
	}
	nick = strings.ToLower(nick)
	return nick, nil
}

// clearNumber clears a number
func FormatDocument(document string) (string, error) {
	if strings.Trim(document, "") == "" {
		return "", errors.New("document is blank")
	}
	re := regexp.MustCompile(`[^0-9]`)
	doc := re.ReplaceAllString(document, "")
	if doc == "" {
		return "", errors.New("invalid document")
	}
	idoc, _ := strconv.ParseUint(doc, 10, 64)
	doc = fmt.Sprintf("%011d", idoc)
	if cpf_cnpj.ValidateCPF(doc) {
		return doc, nil
	}
	doc = fmt.Sprintf("%014d", idoc)
	if cpf_cnpj.ValidateCNPJ(doc) {
		return doc, nil
	}
	return "", errors.New("invalid document")
}

// clearEmail clears an email address
func FormatEmail(email string) (string, error) {
	if strings.Trim(email, "") == "" {
		return "", errors.New("email is blank")
	}
	e, err := mail.ParseAddress(email)
	if err != nil {
		return "", errors.New("invalid email")
	}
	return e.Address, nil
}

// clearPhone clears a phone number
func FormatPhone(number string) (string, error) {
	if strings.Trim(number, "") == "" {
		return "", errors.New("phone is blank")
	}
	p := ""
	for _, c := range countries {
		r := phone.Parse(number, c)
		if r != "" {
			p = r
			break
		}
	}
	if p == "" {
		return "", errors.New("invalid cell phone")
	}
	return p, nil
}
