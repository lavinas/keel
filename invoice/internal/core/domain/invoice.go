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
	ClientID      string         `json:"client_id"      gorm:"type:varchar(50); not null"`
	Client        *Client        `json:"client"         gorm:"foreignKey:BusinessID,ClientID;associationForeignKey:BusinessID,ID"`
	DateStr       string         `json:"date"           gorm:"-"`
	Date          time.Time      `json:"-"              gorm:"type:date; not null"`
	DueStr        string         `json:"due"            gorm:"-"`
	Due           time.Time      `json:"-"              gorm:"type:date; not null"`
	AmountStr     string         `json:"amount"         gorm:"-"`
	Amount        float64        `json:"-"              gorm:"type:decimal(20, 2); not null"`
	InstructionID string         `json:"instruction_id" gorm:"type:varchar(50)"`
	Instruction   *Instruction   `json:"instruction"    gorm:"foreignKey:BusinessID,InstructionID;associationForeignKey:BusinessID,ID"`
	Item          []*Item        `json:"items"          gorm:"foreignKey:BusinessID,InvoiceID;associationForeignKey:BusinessID,ID"`
	StatusID      string         `json:"-"              gorm:"type:varchar(50); not null"`
	InvoiceStatus *InvoiceStatus `json:"status"         gorm:"foreignKey:ID,StatusID;associationForeignKey:InvoiceID,ID"`
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

// SetCreate sets the created_at and updated_at of the model
func (i *Invoice) SetCreate(business_id string) {
	i.Base.SetCreate(business_id)
	i.InvoiceStatus = NewInvoiceStatus(i.Base.ID)
	i.InvoiceStatus.SetCreated()
}

// Fit fits the invoice information received
func (i *Invoice) Fit() {
	i.Base.Fit()
	if i.Client != nil {
		i.Client.SetCreate(i.BusinessID)
		i.Client.Fit()
	}
	if i.Instruction != nil {
		i.Instruction.SetCreate(i.BusinessID)
		i.Instruction.Fit()
	}
	for _, item := range i.Item {
		item.SetCreate(i.BusinessID)
		item.InvoiceID = i.ID
		item.Fit()
	}
	i.Date, _ = time.Parse("2006-01-02", i.DateStr)
	i.Due, _ = time.Parse("2006-01-02", i.DueStr)
	i.Amount, _ = strconv.ParseFloat(i.AmountStr, 64)
	if i.InvoiceStatus == nil {
		i.InvoiceStatus = NewInvoiceStatus(i.Base.ID)
	}
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
		err.SetPrefix("client ")
		return err
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
		err.SetPrefix("instruction ")
		return err
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
			err.SetPrefix("item " + strconv.Itoa(count) + " ")
			return err
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

// TableName returns the table name for gorm
func (b *Invoice) TableName() string {
	return "invoice"
}
