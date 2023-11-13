package dto

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lavinas/keel/invoice/internal/core/port"
	"github.com/lavinas/keel/invoice/pkg/ktools"
)

const (
	// ErrReferenceEmpty is the error message for an empty reference
	ErrReferenceEmpty = "reference is empty"
	// ErrBusinessNicknameEmpty is the error message for an empty business nickname
	ErrBusinessNicknameEmpty = "business nickname is empty"
	// ErrCustomerNicknameEmpty is the error message for an empty customer nickname
	ErrCustomerNicknameEmpty = "customer nickname is empty"
	// ErrAmountEmpty is the error message for an empty amount
	ErrAmountEmpty = "amount is empty"
	// ErrAmountInvalid is the error message for an invalid amount
	ErrAmountInvalid = "amount is invalid"
	// ErrAmountZeroOrNegative is the error message for a negative or zero invalid amount
	ErrAmountZeroOrNegative = "amount is negative or zero"
	// ErrDateEmpty is the error message for an empty date
	ErrDateEmpty = "date is empty"
	// ErrDateInvalid is the error message for an invalid date
	ErrDateInvalid = "date is invalid"
	// OldDate is a base date for comparison if the date is in too far in the past
	OldDate = "2000-01-01"
	// ErrOldDate is the error message for a date too old
	ErrOldDate = "date is too old"
	// ErrDueDateEmpty is the error message for an empty due date
	ErrDueDateEmpty = "due date is empty"
	// ErrDueDateInvalid is the error message for an invalid due date
	ErrDueDateInvalid = "due date is invalid"
	// ErrDueDateOlderThanDate is the error message for a due date older than the date
	ErrDueDateOlderThanDate = "due date is older than date"
)

// InsertInputDto is the DTO for the crate a new invoice
type CreateInputDto struct {
	Reference        string               `json:"reference"`
	BusinessNickname string               `json:"business_nickname"`
	CustomerNickname string               `json:"customer_nickname"`
	Amount           string               `json:"amount"`
	Date             string               `json:"date"`
	Due              string               `json:"due"`
	NoteReference    string               `json:"note_name"`
	Items            []CreateInputItemDto `json:"items"`
}

// Validate validates the InsertInputDto
func (i *CreateInputDto) Validate() error {
	execOrder := []func() error{
		i.ValidateReference,
		i.ValidateBusinessNickname,
		i.ValidateCustomerNickname,
		i.ValidateAmount,
		i.ValidateDate,
		i.ValidateDue,
		i.ValidateItems,
	}
	var errs []error
	for _, v := range execOrder {
		errs = append(errs, v())
	}
	if err := ktools.MergeError(errs...); err != nil {
		return err
	}
	return nil
}

// Format formats the InsertInputDto
func (i *CreateInputDto) Format() error {
	orderMap := []func() error{
		i.FormatReference,
		i.FormatBusinessNickname,
		i.FormatCustomerNickname,
		i.FormatAmount,
		i.FormatDate,
		i.FormatDue,
	}
	var errs []error
	for _, v := range orderMap {
		if err := v(); err != nil {
			errs = append(errs, v())
		}
	}
	if err := ktools.MergeError(errs...); err != nil {
		return err
	}
	return nil
}

// GetReference returns the reference
func (i *CreateInputDto) GetReference() string {
	return i.Reference
}

// GetBusinessNickname returns the business nickname
func (i *CreateInputDto) GetBusinessNickname() string {
	return i.BusinessNickname
}

// GetCustomerNickname returns the customer nickname
func (i *CreateInputDto) GetCustomerNickname() string {
	return i.CustomerNickname
}

// GetAmount returns the amount
func (i *CreateInputDto) GetAmount() (float64, error) {
	f, err := strconv.ParseFloat(i.Amount, 64)
	if err != nil {
		return 0, errors.New(ErrAmountInvalid)
	}
	return f, nil
}

// GetDate returns the date
func (i *CreateInputDto) GetDate() (time.Time, error) {
	date, err := time.Parse("2006-01-02", i.Date)
	if err != nil {
		return time.Now(), errors.New(ErrDateInvalid)
	}
	return date, nil
}

// GetDue returns the due
func (i *CreateInputDto) GetDue() (time.Time, error) {
	due, err := time.Parse("2006-01-02", i.Due)
	if err != nil {
		return time.Now(), errors.New(ErrDueDateInvalid)
	}
	return due, nil
}

// GetNoteReference returns the note reference
func (i *CreateInputDto) GetNoteReference() string {
	return i.NoteReference
}

// GetItems returns the items
func (i *CreateInputDto) GetItems() []port.CreateInputItemDto {
	var items []port.CreateInputItemDto
	if len(i.Items) == 0 {
		return nil
	}
	for _, item := range i.Items {
		items = append(items, item)
	}
	return items
}

// Validate reference validates the reference
func (i *CreateInputDto) ValidateReference() error {
	i.Reference = strings.Trim(i.Reference, " ")
	if i.Reference == "" {
		return errors.New(ErrReferenceEmpty)
	}
	return nil
}

// ValidateBusinessNickname validates the business nickname
func (i *CreateInputDto) ValidateBusinessNickname() error {
	i.BusinessNickname = strings.Trim(i.BusinessNickname, " ")
	if i.BusinessNickname == "" {
		return errors.New(ErrBusinessNicknameEmpty)
	}
	return nil
}

// ValidateCustomerNickname validates the customer nickname
func (i *CreateInputDto) ValidateCustomerNickname() error {
	i.CustomerNickname = strings.Trim(i.CustomerNickname, " ")
	if i.CustomerNickname == "" {
		return errors.New(ErrCustomerNicknameEmpty)
	}
	return nil
}

// ValidateAmount validates the amount
func (i *CreateInputDto) ValidateAmount() error {
	i.Amount = strings.Trim(i.Amount, " ")
	if i.Amount == "" {
		return errors.New(ErrAmountEmpty)
	}
	f, err := strconv.ParseFloat(i.Amount, 64)
	if err != nil {
		return errors.New(ErrAmountInvalid)
	}
	if f <= 0 {
		return errors.New(ErrAmountZeroOrNegative)
	}
	return nil
}

// ValidateDate validates the date
func (i *CreateInputDto) ValidateDate() error {
	i.Date = strings.Trim(i.Date, " ")
	if i.Date == "" {
		return errors.New(ErrDateEmpty)
	}
	if _, err := time.Parse("2006-01-02", i.Date); err != nil {
		return errors.New(ErrDateInvalid)
	}
	if i.Date < OldDate {
		return errors.New(ErrOldDate)
	}
	return nil
}

// ValidateDue validates the due
func (i *CreateInputDto) ValidateDue() error {
	i.Due = strings.Trim(i.Due, " ")
	if i.Due == "" {
		return errors.New(ErrDueDateEmpty)
	}
	if _, err := time.Parse("2006-01-02", i.Due); err != nil {
		return errors.New(ErrDueDateInvalid)
	}
	if i.ValidateDate() == nil && i.Date > i.Due {
		return errors.New(ErrDueDateOlderThanDate)
	}
	return nil
}

// ValidateItems validates all items if there are any
func (i *CreateInputDto) ValidateItems() error {
	message := ""
	if len(i.Items) != 0 {
		position := 1
		for _, item := range i.Items {
			if item.Validate() != nil {
				message += fmt.Sprintf("item %d: %s | ", position, item.Validate().Error())
			}
			position++
		}
	}
	if message != "" {
		message = message[:len(message)-3]
		return errors.New(message)
	}
	return nil
}

// FormatReference formats the reference
func (i *CreateInputDto) FormatReference() error {
	if err := i.ValidateReference(); err != nil {
		return err
	}
	i.Reference = strings.ToLower(strings.ReplaceAll(strings.Trim(i.Reference, " "), " ", "_"))
	return nil
}

// FormatBusinessNickname formats the business nickname
func (i *CreateInputDto) FormatBusinessNickname() error {
	if err := i.ValidateBusinessNickname(); err != nil {
		return err
	}
	i.BusinessNickname = strings.ToLower(strings.ReplaceAll(strings.Trim(i.BusinessNickname, " "), " ", "_"))
	return nil
}

// FormatCustomerNickname formats the customer nickname
func (i *CreateInputDto) FormatCustomerNickname() error {
	if err := i.ValidateCustomerNickname(); err != nil {
		return err
	}
	i.CustomerNickname = strings.ToLower(strings.ReplaceAll(strings.Trim(i.CustomerNickname, " "), " ", "_"))
	return nil
}

// FormatAmount formats the amount
func (i *CreateInputDto) FormatAmount() error {
	if err := i.ValidateAmount(); err != nil {
		return err
	}
	f, _ := strconv.ParseFloat(i.Amount, 64)
	i.Amount = fmt.Sprintf("%.2f", f)
	return nil
}

// FormatDate formats the date
func (i *CreateInputDto) FormatDate() error {
	if err := i.ValidateDate(); err != nil {
		return err
	}
	i.Date = strings.Trim(i.Date, " ")
	return nil
}

// FormatDue formats the due
func (i *CreateInputDto) FormatDue() error {
	if err := i.ValidateDue(); err != nil {
		return err
	}
	i.Due = strings.Trim(i.Due, " ")
	return nil
}
