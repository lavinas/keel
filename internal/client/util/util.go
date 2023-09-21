package util

import (
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"errors"

	"github.com/lavinas/keel/pkg/cpf_cnpj"
	"github.com/lavinas/keel/pkg/phone"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	// countries is the list of countries for phone number validation
	countries = []string{"BR"}
)

// Util is the util functions to be used in the client services
type Util struct {
}

// NewUtil creates a new util
func NewUtil() *Util {
	return &Util{}
}

// Validate name validates a name
func (d *Util) ValidateName(name string) (bool, string) {
	if _, err := d.ClearName(name); err != nil {
		return false, err.Error()
	}
	return true, ""
}

// ValidateNickname validates a nickname
func (d *Util) ValidateNickname(nickname string) (bool, string) {
	if _, err := d.ClearNickname(nickname); err != nil {
		return false, err.Error()
	}
	return true, ""
}

// ValidateDocument validates a document
func (d *Util) ValidateDocument(document string) (bool, string) {
	_, err := d.ClearDocument(document)
	if err != nil {
		return false, err.Error()
	}
	return true, ""
}

// ValidateEmail validates an email address
func (d *Util) ValidateEmail(email string) (bool, string) {
	if _, err := d.ClearEmail(email); err != nil {
		return false, err.Error()
	}
	return true, ""
}

// ValidatePhone validates a phone number
func (d *Util) ValidatePhone(number string) (bool, string) {
	_, err := d.ClearPhone(number)
	if err != nil {
		return false, err.Error()
	}
	return true, ""
}

// ValidateAll validates all fields (name, nickname, document, phone, email)
func (d *Util) ValidateAll(name, nickname, document, phone, email string) (bool, string) {
	message := ""
	if b, m := d.ValidateName(name); !b {
		message += m + " || "
	}
	if b, m := d.ValidateNickname(nickname); !b {
		message += m + " || "
	}
	if b, m := d.ValidateDocument(document); !b {
		message += m + " || "
	}
	if b, m := d.ValidatePhone(phone); !b {
		message += m + " || "
	}
	if b, m := d.ValidateEmail(email); !b {
		message += m + " || "
	}
	if message != "" {
		message = message[:len(message)-4]
		return false, message
	}
	return true, ""
}

//ClearName clears a name
func (d *Util) ClearName(name string) (string, error) {
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

//ClearNickname clears a nickname
func (d *Util) ClearNickname(nickname string) (string, error) {
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

// ClearNumber clears a number
func (d *Util) ClearDocument(document string) (uint64, error) {
	re := regexp.MustCompile(`[^0-9]`)
	doc := re.ReplaceAllString(document, "")
	if doc == "" {
		return 0, errors.New("invalid document")
	}
	idoc, _ := strconv.ParseUint(doc, 10, 64)
	doc = strconv.FormatUint(idoc, 10)
	if cpf_cnpj.ValidateCPF(doc) {
		return idoc, nil
	}
	if cpf_cnpj.ValidateCNPJ(doc) {
		return idoc, nil
	}
	return 0, errors.New("invalid document")
}

// ClearEmail clears an email address
func (d *Util) ClearEmail(email string) (string, error) {
	e, err := mail.ParseAddress(email)
	if err != nil {
		return "", errors.New("invalid email")
	}
	return e.Address, nil
}

//ClearPhone clears a phone number
func (d *Util) ClearPhone(number string) (uint64, error) {
	p := ""
	for _, c := range countries {
		r := phone.Parse(number, c)
		if r != "" {
			p = r
			break
		}
	}
	if p == "" {
		return 0, errors.New("invalid cell phone")
	}
	ip, _ := strconv.ParseUint(p, 10, 64) 
	return ip, nil
}

// ClearAll clears all fields (name, nickname, document, phone, email)
func (d *Util) ClearAll(name, nickname, document, phone, email string) (string, string, uint64, uint64, string, error) {
	iname, err := d.ClearName(name)
	if err != nil {
		return "", "", 0, 0, "", err
	}
	inick, err := d.ClearNickname(nickname)
	if err != nil {
		return "", "", 0, 0, "", err
	}
	idoc, err := d.ClearDocument(document)
	if err != nil {
		return "", "", 0, 0, "", err
	}
	iph, err := d.ClearPhone(phone)
	if err != nil {
		return "", "", 0, 0, "", err
	}
	iemail, err := d.ClearEmail(email)
	if err != nil {
		return "", "", 0, 0, "", err
	}
	return iname, inick, idoc, iph, iemail, nil
}
