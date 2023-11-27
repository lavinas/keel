package domain

import (
	"errors"
	"strconv"
	"time"

	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Invoice represents an invoice - main model
type Invoice struct {
	Base
	ClientID    string       `json:"client_id"   gorm:"type:varchar(50); not null"`
	Client      Client       `json:"-"           gorm:"foreignKey:BusinessID,ClientID;associationForeignKey:BusinessID,ID"`
	DateStr     string       `json:"date"        gorm:"-"`
	Date        time.Time    `json:"-"           gorm:"type:date; not null"`
	DueStr      string       `json:"due"         gorm:"-"`
	Due         time.Time    `json:"-"           gorm:"type:date; not null"`
	AmountStr   string       `json:"amount"      gorm:"-"`
	Amount      float64      `json:"-"           gorm:"type:decimal(20, 2); not null"`
	Instruction *Instruction `json:"instruction" gorm:"type:varchar(100)"`
}

// Validate validates the invoice
func (i *Invoice) Validate(repo port.Repository) error {
	return ValidateLoop([]func(repo port.Repository) error{
		i.Base.Validate,
		i.ValidateClient,
		i.ValidateDate,
		i.ValidateDue,
		i.ValidateAmount,
		i.ValidateInstruction,
	}, repo)
}

// Marshal marshals the invoice
func (i *Invoice) Marshal() error {
	if err := i.MarshalDate(); err != nil {
		return err
	}
	if err := i.MarshalDue(); err != nil {
		return err
	}
	if err := i.MarshalAmount(); err != nil {
		return err
	}
	return nil
}

// ValidateClient validates the client of the invoice
func (i *Invoice) ValidateClient(repo port.Repository) error {
	if i.ClientID == "" {
		return errors.New(ErrInvoiceClientIsRequired)
	}
	var client Client
	client.ID = i.ClientID
	if err := client.ValidateID(repo); err != nil {
		return errors.New("client " + err.Error())
	}
	if !repo.Exists(&client, i.ClientID) {
		return errors.New(ErrInvoiceClientNotFound)
	}
	return nil
}

// ValidateDate validates the Date of the invoice
func (i *Invoice) ValidateDate(repo port.Repository) error {
	if i.DateStr == "" {
		return errors.New(ErrInvoiceDateIsRequired)
	}
	if _, err := time.Parse("2006-01-02", i.DateStr); err != nil {
		return errors.New(ErrInvoiceDateIsInvalid)
	}
	return nil
}

// MarshalDate marshals the Date of the invoice
func (i *Invoice) MarshalDate() error {
	var err error
	if i.Date, err = time.Parse("2006-01-02", i.DateStr); err != nil {
		return err
	}
	return nil
}

// ValidateDue validates the Due Date of the invoice
func (i *Invoice) ValidateDue(repo port.Repository) error {
	if i.DueStr == "" {
		return errors.New(ErrInvoiceDueIsRequired)
	}
	if _, err := time.Parse("2006-01-02", i.DueStr); err != nil {
		return errors.New(ErrInvoiceDateIsInvalid)
	}
	if i.Due.Before(i.Date) {
		return errors.New(ErrInvoiceDueBeforeDate)
	}
	return nil
}

// MarshalDue marshals the Due Date of the invoice
func (i *Invoice) MarshalDue() error {
	var err error
	if i.Due, err = time.Parse("2006-01-02", i.DueStr); err != nil {
		return errors.New(ErrInvoiceDueIsInvalid)
	}
	return nil
}

// ValidateAmount validates the amount of the invoice
func (i *Invoice) ValidateAmount(repo port.Repository) error {
	if i.AmountStr == "" {
		return errors.New(ErrInvoiceAmountIsRequired)
	}
	if v, err := strconv.ParseFloat(i.AmountStr, 64); err != nil {
		return errors.New(ErrInvoiceAmountIsInvalid)
	} else if v <= 0 {
		return errors.New(ErrInvoiceAmountIsInvalid)
	}
	return nil
}

// MarshalAmount marshals the amount of the invoice
func (i *Invoice) MarshalAmount() error {
	var err error
	if i.Amount, err = strconv.ParseFloat(i.AmountStr, 64); err != nil {
		return errors.New(ErrInvoiceAmountIsInvalid)
	}
	return nil
}

// ValidateInstruction validates the instruction of the invoice
func (i *Invoice) ValidateInstruction(repo port.Repository) error {
	if i.Instruction == nil {
		return nil
	}
	return i.Instruction.Validate(repo)
}
