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
	ClientID      string       `json:"client_id"      gorm:"type:varchar(50); not null"`
	Client        *Client      `json:"-"              gorm:"foreignKey:BusinessID,ClientID;associationForeignKey:BusinessID,ID"`
	DateStr       string       `json:"date"           gorm:"-"`
	Date          time.Time    `json:"-"              gorm:"type:date; not null"`
	DueStr        string       `json:"due"            gorm:"-"`
	Due           time.Time    `json:"-"              gorm:"type:date; not null"`
	AmountStr     string       `json:"amount"         gorm:"-"`
	Amount        float64      `json:"-"              gorm:"type:decimal(20, 2); not null"`
	InstructionID string       `json:"instruction_id" gorm:"type:varchar(50)"`
	Instruction   *Instruction `json:"-"              gorm:"foreignKey:BusinessID,InstructionID;associationForeignKey:BusinessID,ID"`
	Item          []*Item      `json:"items"          gorm:"foreignKey:BusinessID,InvoiceID;associationForeignKey:BusinessID,ID"`
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
		i.ValidateItem,
	}, repo)
}

// Fit fits the invoice information received
func (i *Invoice) Fit() {
	i.Base.Fit()
	if i.Client != nil {
		i.Client.Fit()
	}
	if i.Instruction != nil {
		i.Instruction.Fit()
	}
	for _, item := range i.Item {
		item.SetBusinessID(i.BusinessID)
		item.SetCreatedAt(i.Created_at)
		item.SetUpdatedAt(i.Updated_at)
		item.InvoiceID = i.ID
		item.Fit()
	}
	i.Date, _ = time.Parse("2006-01-02", i.DateStr)
	i.Due, _ = time.Parse("2006-01-02", i.DueStr)
	i.Amount, _ = strconv.ParseFloat(i.AmountStr, 64)
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

// validateItem validates the invoice item
func (i *Invoice) ValidateItem(repo port.Repository) error {
	if len(i.Item) == 0 {
		return nil
	}
	count := 0; sum := 0.0
	for _, item := range i.Item {
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
