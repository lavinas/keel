package domain

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/lavinas/keel/invoice/internal/core/port"
)

// Invoice represents an invoice - main model
type Invoice struct {
	Base
	ClientID      string         `json:"client_id"      gorm:"type:varchar(50); not null"`
	Client        *Client        `json:"-"              gorm:"foreignKey:BusinessID,ClientID;associationForeignKey:BusinessID,ID"`
	DateStr       string         `json:"date"           gorm:"-"`
	Date          time.Time      `json:"-"              gorm:"type:date; not null"`
	DueStr        string         `json:"due"            gorm:"-"`
	Due           time.Time      `json:"-"              gorm:"type:date; not null"`
	AmountStr     string         `json:"amount"         gorm:"-"`
	Amount        float64        `json:"-"              gorm:"type:decimal(20, 2); not null"`
	InstructionID string         `json:"instruction_id" gorm:"type:varchar(50)"`
	Instruction   *Instruction   `json:"-"              gorm:"foreignKey:BusinessID,InstructionID;associationForeignKey:BusinessID,ID"`
	Items         []*InvoiceItem `json:"items"          gorm:"foreignKey:InvoiceID;associationForeignKey:ID"`
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
	for _, f := range []func() error{
		i.MarshalDate,
		i.MarshalDue,
		i.MarshalAmount,
		i.MarshalInvoiceItem,
	} {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

// ValidateClient validates the client of the invoice
func (i *Invoice) ValidateClient(repo port.Repository) error {
	if i.ClientID == "" && i.Client == nil {
		return errors.New(ErrInvoiceClientIsRequired)
	} else if i.ClientID != "" && i.Client != nil {
		return errors.New(ErrInvoiceClientInformedTwice)
	} else if i.ClientID != "" {
		return i.ValidateClientID(repo)
	} else if err := i.Client.Validate(repo); err != nil {
		return errors.New("client " + err.Error())
	}
	return nil
}

// ValidateClientID validates the client id of the invoice
func (i *Invoice) ValidateClientID(repo port.Repository) error {
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
	if i.InstructionID == "" {
		return errors.New(ErrInvoiceInstructionIDIsRequired)
	}
	var instruction Instruction
	instruction.ID = i.InstructionID
	if err := instruction.ValidateID(repo); err != nil {
		return errors.New("instruction " + err.Error())
	}
	if !repo.Exists(&instruction, i.InstructionID) {
		return errors.New(ErrInvoiceInstructionNotFound)
	}
	return nil
}

// validateInvoiceItem validates the invoice item
func (i *Invoice) ValidateInvoiceItem(repo port.Repository) error {
	if len(i.Items) == 0 {
		return nil
	}
	count := 0
	sum := 0.0
	for _, item := range i.Items {
		count++
		if err := item.Validate(repo); err != nil {
			return fmt.Errorf("item %d: %s", count, err.Error())
		}
		sum += item.GetAmount()
	}
	if sum != i.Amount {
		return errors.New(ErrInvoiceAmountUnmatch)
	}
	return nil
}

// MarshalInvoiceItem marshals the invoice item
func (i *Invoice) MarshalInvoiceItem() error {
	if i.Items == nil {
		return nil
	}
	for _, item := range i.Items {
		if err := item.Marshal(); err != nil {
			return err
		}
	}
	return nil
}
