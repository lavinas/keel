package domain

import (
	"errors"
	"strings"
	"time"
)

// Invoice represents an invoice - main model
type Invoice struct {
	Base
	Client      string       `json:"client" gorm:"type:varchar(100)"`
	Date        time.Time    `json:"date" gorm:"type:date"`
	Due         time.Time    `json:"due"  gorm:"type:date"`
	Amount      float64      `json:"amount" gorm:"type:decimal(20, 2)"`
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
	if p.Date.IsZero() {
		return errors.New(ErrInvoiceDateIsRequired)
	}
	return nil
}

// ValidateDue validates the Due Date of the invoice
func (p *Invoice) ValidateDue() error {
	if p.Due.IsZero() {
		return errors.New(ErrInvoiceDueIsRequired)
	}
	return nil
}

// ValidateAmount validates the amount of the invoice
func (p *Invoice) ValidateAmount() error {
	if p.Amount <= 0 {
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
