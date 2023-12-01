package domain

import (
	"strconv"
	"time"

	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/pkg/kerror"
)

// Invoice represents an invoice - main model
type Invoice struct {
	Base
	ClientID      string       `json:"client_id"      gorm:"type:varchar(50); not null"`
	Client        *Client      `json:"client"         gorm:"foreignKey:BusinessID,ClientID;associationForeignKey:BusinessID,ID"`
	DateStr       string       `json:"date"           gorm:"-"`
	Date          time.Time    `json:"-"              gorm:"type:date; not null"`
	DueStr        string       `json:"due"            gorm:"-"`
	Due           time.Time    `json:"-"              gorm:"type:date; not null"`
	AmountStr     string       `json:"amount"         gorm:"-"`
	Amount        float64      `json:"-"              gorm:"type:decimal(20, 2); not null"`
	InstructionID string       `json:"instruction_id" gorm:"type:varchar(50)"`
	Instruction   *Instruction `json:"instruction"    gorm:"foreignKey:BusinessID,InstructionID;associationForeignKey:BusinessID,ID"`
	Item          []*Item      `json:"items"          gorm:"foreignKey:BusinessID,InvoiceID;associationForeignKey:BusinessID,ID"`
}

// Validate validates the invoice
func (i *Invoice) Validate(repo port.Repository) *kerror.KError {
	return ValidateLoop([]func(repo port.Repository) *kerror.KError{
		i.Base.Validate,
		i.ValidateClient,
		i.ValidateDate,
		i.ValidateDue,
		i.ValidateAmount,
		i.ValidateInstruction,
		i.ValidateItem,
		i.ValidateDuplicity,
	}, repo)
}

// Fit fits the invoice information received
func (i *Invoice) Fit() {
	i.Base.Fit()
	if i.Client != nil {
		i.Client.SetBusinessID(i.BusinessID)
		i.Client.SetCreatedAt(i.Created_at)
		i.Client.SetUpdatedAt(i.Updated_at)
		i.Client.Fit()
	}
	if i.Instruction != nil {
		i.Instruction.SetBusinessID(i.BusinessID)
		i.Instruction.SetCreatedAt(i.Created_at)
		i.Instruction.SetUpdatedAt(i.Updated_at)
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
func (i *Invoice) ValidateClient(repo port.Repository) *kerror.KError {
	if i.ClientID == "" && i.Client == nil {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceClientIsRequired)
	} else if i.ClientID != "" && i.Client != nil {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceClientInformedTwice)
	} else if i.ClientID != "" {
		return i.ValidateClientID(repo)
	} else if err := i.Client.Validate(repo); err != nil {
		return kerror.NewKError(kerror.BadRequest, "client "+err.Error())
	}
	return nil
}

// ValidateClientID validates the client id of the invoice
func (i *Invoice) ValidateClientID(repo port.Repository) *kerror.KError {
	if i.ClientID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceClientIsRequired)
	}
	var client Client
	client.ID = i.ClientID
	if err := client.ValidateID(repo); err != nil {
		return kerror.NewKError(kerror.BadRequest, "client "+err.Error())
	}
	if exists, err := repo.Exists(&client, i.BusinessID, i.ClientID); err != nil {
		return kerror.NewKError(kerror.Internal, ErrInvoiceClientNotFound)
	} else if !exists {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceClientNotFound)
	}
	return nil
}

// ValidateDate validates the Date of the invoice
func (i *Invoice) ValidateDate(repo port.Repository) *kerror.KError {
	if i.DateStr == "" {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceDateIsRequired)
	}
	if _, err := time.Parse("2006-01-02", i.DateStr); err != nil {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceDateIsInvalid)
	}
	return nil
}

// ValidateDue validates the Due Date of the invoice
func (i *Invoice) ValidateDue(repo port.Repository) *kerror.KError {
	if i.DueStr == "" {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceDueIsRequired)
	}
	if _, err := time.Parse("2006-01-02", i.DueStr); err != nil {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceDateIsInvalid)
	}
	if i.Due.Before(i.Date) {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceDueBeforeDate)
	}
	return nil
}

// ValidateAmount validates the amount of the invoice
func (i *Invoice) ValidateAmount(repo port.Repository) *kerror.KError {
	if i.AmountStr == "" {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceAmountIsRequired)
	}
	if v, err := strconv.ParseFloat(i.AmountStr, 64); err != nil {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceAmountIsInvalid)
	} else if v <= 0 {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceAmountIsInvalid)
	}
	return nil
}

// ValidateInstruction validates the instruction of the invoice
func (i *Invoice) ValidateInstruction(repo port.Repository) *kerror.KError {
	if i.InstructionID == "" && i.Instruction == nil {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceInstructionIDIsRequired)
	} else if i.InstructionID != "" && i.Instruction != nil {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceInstructionInformedTwice)
	} else if i.InstructionID != "" {
		return i.ValidateInstructionID(repo)
	} else if err := i.Instruction.Validate(repo); err != nil {
		return kerror.NewKError(kerror.BadRequest, "instruction "+err.Error())
	}
	return nil
}

// ValidateInstructionID validates the instruction id of the invoice
func (i *Invoice) ValidateInstructionID(repo port.Repository) *kerror.KError {
	if i.InstructionID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceInstructionIDIsRequired)
	}
	var instruction Instruction
	instruction.ID = i.InstructionID
	if err := instruction.ValidateID(repo); err != nil {
		return kerror.NewKError(kerror.BadRequest, "instruction "+err.Error())
	}
	if exists, err := repo.Exists(&instruction, i.BusinessID, i.InstructionID); err != nil {
		return kerror.NewKError(kerror.Internal, ErrInvoiceInstructionNotFound)
	} else if !exists {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceInstructionNotFound)
	}
	return nil
}

// validateItem validates the invoice item
func (i *Invoice) ValidateItem(repo port.Repository) *kerror.KError {
	if len(i.Item) == 0 {
		return nil
	}
	count := 0
	sum := 0.0
	for _, item := range i.Item {
		count++
		if err := item.Validate(repo); err != nil {
			return kerror.NewKError(kerror.BadRequest, "item "+err.Error())
		}
		sum += item.GetAmount()
	}
	if sum != i.Amount {
		return kerror.NewKError(kerror.BadRequest, ErrInvoiceAmountUnmatch)
	}
	return nil
}

// ValidateDuplicity validates the duplicity of the model
func (b *Invoice) ValidateDuplicity(repo port.Repository) *kerror.KError {
	return b.Base.ValidateDuplicity(b, repo)
}
