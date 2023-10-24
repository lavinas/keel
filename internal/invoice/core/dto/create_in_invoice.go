package dto

import (
	"errors"
	"strconv"
	"time"
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

// Validate validates the InsertInputDto
func (i CreateInputDto) Validate() error {
	validationMap := map[string]func() error{
		"reference":        i.ValidateReference,
		"business_nickname": i.ValidateBusinessNickname,
		"customer_nickname": i.ValidateCustomerNickname,
		"amount":           i.ValidateAmount,
		"date":             i.ValidateDate,
		"due":              i.ValidateDue,
		"items":			i.ValidateItems,
	}
	var message string = ""
	for _, v := range validationMap {
		if v() != nil {
			message += message + v().Error() + " | "
		}
	}
	if message != "" {
		message = message[:len(message)-3]
		return errors.New(message)
	}
	return nil
}

// Validate reference validates the reference
func (i CreateInputDto) ValidateReference() error {
	if i.Reference == "" {
		return errors.New(ErrReferenceEmpty)
	}
	return nil
}

// ValidateBusinessNickname validates the business nickname
func (i CreateInputDto) ValidateBusinessNickname() error {
	if i.BusinessNickname == "" {
		return errors.New(ErrReferenceEmpty)
	}
	return nil
}

// ValidateCustomerNickname validates the customer nickname
func (i CreateInputDto) ValidateCustomerNickname() error {
	if i.CustomerNickname == "" {
		return errors.New(ErrReferenceEmpty)
	}
	return nil
}

// ValidateAmount validates the amount
func (i CreateInputDto) ValidateAmount() error {
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
func (i CreateInputDto) ValidateDate() error {
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
func (i CreateInputDto) ValidateDue() error {
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
func (i CreateInputDto) ValidateItems() error {
	message := ""
	if len(i.Items) != 0 {
		for _, item := range i.Items {
			if item.Validate() != nil {
				message += item.Reference + ": " + item.Validate().Error() + " | "
			}
		}
	}
	if message != "" {
		message = message[:len(message)-3]
		return errors.New(message)
	}
	return nil
}
