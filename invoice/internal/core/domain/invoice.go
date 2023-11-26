package domain

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Invoice represents an invoice - main model
type Invoice struct {
	Base
	Client      string       `json:"client" gorm:"type:varchar(100)"`
	DateStr     string       `json:"date" gorm:"-"`
	Date        time.Time    `json:"-" gorm:"type:date"`
	DueStr      string       `json:"due" gorm:"-"`
	Due         time.Time    `json:"-"  gorm:"type:date"`
	AmountStr   string       `json:"amount" gorm:"-"`
	Amount      float64      `json:"-" gorm:"type:decimal(20, 2)"`
	Instruction *Instruction `json:"instruction" gorm:"type:varchar(100)"`
}

// Validate validates the invoice
func (p *Invoice) Validate() error {
	return ValidateLoop([]func() error{
		p.Base.Validate,
		p.ValidateDate,
		p.ValidateDue,
		p.ValidateAmount,
		p.ValidateInstruction,
	})
}

// Marshal marshals the invoice
func (p *Invoice) Marshal() error {
	if err := p.MarshalDate(); err != nil {
		return err
	}
	if err := p.MarshalDue(); err != nil {
		return err
	}
	if err := p.MarshalAmount(); err != nil {
		return err
	}
	return nil
}

// ValidateClient validates the client of the invoice
func (p *Invoice) ValidateClient() error {
	if p.Client == "" {
		return errors.New(ErrInvoiceClientIsRequired)
	}
	if len(strings.Split(p.Client, " ")) < 2 {
		return errors.New(ErrClientNameLength)
	}
	if p.Client != strings.ToLower(p.Client) {
		return errors.New(ErrClientIDNotLower)
	}
	return nil
}

// ValidateDate validates the Date of the invoice
func (p *Invoice) ValidateDate() error {
	if p.DateStr == "" {
		return errors.New(ErrInvoiceDateIsRequired)
	}
	if _, err := time.Parse("2006-01-02", p.DateStr); err != nil {
		return errors.New(ErrInvoiceDateIsInvalid)
	}
	return nil
}

// MarshalDate marshals the Date of the invoice
func (p *Invoice) MarshalDate() error {
	var err error
	if p.Date, err = time.Parse("2006-01-02", p.DateStr); err != nil {
		return err
	}
	return nil
}

// ValidateDue validates the Due Date of the invoice
func (p *Invoice) ValidateDue() error {
	if p.DueStr == "" {
		return errors.New(ErrInvoiceDueIsRequired)
	}
	if _, err := time.Parse("2006-01-02", p.DueStr); err != nil {
		return errors.New(ErrInvoiceDateIsInvalid)
	}
	return nil
}

// MarshalDue marshals the Due Date of the invoice
func (p *Invoice) MarshalDue() error {
	var err error
	if p.Due, err = time.Parse("2006-01-02", p.DueStr); err != nil {
		return errors.New(ErrInvoiceDueIsInvalid)
	}
	return nil
}

// ValidateAmount validates the amount of the invoice
func (p *Invoice) ValidateAmount() error {
	if p.AmountStr == "" {
		return errors.New(ErrInvoiceAmountIsRequired)
	}
	if v, err := strconv.ParseFloat(p.AmountStr, 64); err != nil {
		return errors.New(ErrInvoiceAmountIsInvalid)
	} else if v <= 0 {
		return errors.New(ErrInvoiceAmountIsInvalid)
	}
	return nil
}

// MarshalAmount marshals the amount of the invoice
func (p *Invoice) MarshalAmount() error {
	var err error
	if p.Amount, err = strconv.ParseFloat(p.AmountStr, 64); err != nil {
		return errors.New(ErrInvoiceAmountIsInvalid)
	}
	return nil
}

// ValidateInstruction validates the instruction of the invoice
func (p *Invoice) ValidateInstruction() error {
	if p.Instruction == nil {
		return nil
	}
	return p.Instruction.Validate()
}
