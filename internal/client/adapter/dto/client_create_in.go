package dto

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

// ClientCreateInputDto is the input DTO used to create a client
type ClientCreateInputDto struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Document string `json:"document"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// IsBlank checks if the input DTO is blank
func (c *ClientCreateInputDto) IsBlank() bool {
	return strings.Trim(c.Name, " ") == "" && strings.Trim(c.Nickname, " ") == "" &&
		strings.Trim(c.Document, " ") == "" && strings.Trim(c.Phone, " ") == "" &&
		strings.Trim(c.Email, " ") == ""
}

// Validate validates the input DTO
func (c *ClientCreateInputDto) Validate() error {
	msg := ""
	if _, err := formatName(c.Name); err != nil {
		msg += err.Error() + " | "
	}
	if _, err := formatNickname(c.Nickname); err != nil {
		msg += err.Error() + " | "
	}
	if _, err := formatDocument(c.Document); err != nil {
		msg += err.Error() + " | "
	}
	if _, err := formatPhone(c.Phone); err != nil {
		msg += err.Error() + " | "
	}
	if _, err := formatEmail(c.Email); err != nil {
		msg += err.Error() + " | "
	}
	if msg == "" {
		return nil
	}
	msg = strings.Trim(msg, " |")
	return errors.New(msg)
}

// Validate validates the input DTO for update values
func (c *ClientCreateInputDto) ValidateUpdate() error {
	msg := ""
	if strings.Trim(c.Name, " ") != "" {
		if _, err := formatName(c.Name); err != nil {
			msg += err.Error() + " | "
		}
	}
	if strings.Trim(c.Nickname, " ") != "" {
		if _, err := formatNickname(c.Nickname); err != nil {
			msg += err.Error() + " | "
		}
	}
	if strings.Trim(c.Document, " ") != "" {
		if _, err := formatDocument(c.Document); err != nil {
			msg += err.Error() + " | "
		}
	}
	if strings.Trim(c.Phone, " ") != "" {
		if _, err := formatPhone(c.Phone); err != nil {
			msg += err.Error() + " | "
		}
	}
	if strings.Trim(c.Email, " ") != "" {
		if _, err := formatEmail(c.Email); err != nil {
			msg += err.Error() + " | "
		}
	}
	if msg == "" {
		return nil
	}
	msg = strings.Trim(msg, " |")
	return errors.New(msg)
}

// Format clears all fields (name, nickname, document, phone, email)
func (c *ClientCreateInputDto) Format() error {
	var err error
	var name, nick, doc, phone, email string
	if name, err = formatName(c.Name); err != nil {
		return err
	}
	if nick, err = formatNickname(c.Nickname); err != nil {
		return err
	}
	if doc, err = formatDocument(c.Document); err != nil {
		return err
	}
	if phone, err = formatPhone(c.Phone); err != nil {
		return err
	}
	if email, err = formatEmail(c.Email); err != nil {
		return err
	}
	c.Name = name
	c.Nickname = nick
	c.Document = doc
	c.Phone = phone
	c.Email = email
	return nil
}

// FormatUpdate clears all fields (name, nickname, document, phone, email) for update values
func (c *ClientCreateInputDto) FormatUpdate() error {
	var err error
	var name, nick, doc, phone, email string
	if strings.Trim(c.Name, " ") != "" {
		if name, err = formatName(c.Name); err != nil {
			return err
		}
	}
	if strings.Trim(c.Nickname, " ") != "" {
		if nick, err = formatNickname(c.Nickname); err != nil {
			return err
		}
	}
	if strings.Trim(c.Document, " ") != "" {
		if doc, err = formatDocument(c.Document); err != nil {
			return err
		}
	}
	if strings.Trim(c.Phone, " ") != "" {
		if phone, err = formatPhone(c.Phone); err != nil {
			return err
		}
	}
	if strings.Trim(c.Email, " ") != "" {
		if email, err = formatEmail(c.Email); err != nil {
			return err
		}
	}
	c.Name = name
	c.Nickname = nick
	c.Document = doc
	c.Phone = phone
	c.Email = email
	return nil
}

// Get returns all fields (name, nickname, document, phone, email)
func (c *ClientCreateInputDto) Get() (string, string, string, string, string) {
	return c.Name, c.Nickname, c.Document, c.Phone, c.Email
}

// clearName clears a name
func formatName(name string) (string, error) {
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
func formatNickname(nickname string) (string, error) {
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
func formatDocument(document string) (string, error) {
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
func formatEmail(email string) (string, error) {
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
func formatPhone(number string) (string, error) {
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
