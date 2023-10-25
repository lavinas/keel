package dto

import (
	"fmt"
	"strings"
	"testing"
)

func TestInvoiceValidate(t *testing.T) {
	t.Run("should return nil when all fields are valid", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "1.35 ",
			Date:             " 2020-01-01 ",
			Due:              " 2020-01-02 ",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		err := dto.Validate()
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
		if err != nil && err.Error() != ErrItemReferenceEmpty {
			t.Errorf("Expected error %s, got %s", ErrItemReferenceEmpty, err.Error())
		}
	})

	t.Run("should return error when reference is empty", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "1",
			Date:             "2020-01-01",
			Due:              "2020-01-02",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "1",
					Price:       "1",
				},
			},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrReferenceEmpty {
			t.Errorf("Expected error %s, got %s", ErrReferenceEmpty, err.Error())
		}
	})
	t.Run("should return error when business nickname is empty", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "",
			CustomerNickname: "customer",
			Amount:           "1",
			Date:             "2020-01-01",
			Due:              "2020-01-02",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "1",
					Price:       "1",
				},
			},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrBusinessNicknameEmpty {
			t.Errorf("Expected error %s, got %s", ErrBusinessNicknameEmpty, err.Error())
		}
	})
	t.Run("should return error when customer nickname is empty", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "",
			Amount:           "1",
			Date:             "2020-01-01",
			Due:              "2020-01-02",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "1",
					Price:       "1",
				},
			},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrCustomerNicknameEmpty {
			t.Errorf("Expected error %s, got %s", ErrCustomerNicknameEmpty, err.Error())
		}
	})
	t.Run("should return error when amount is empty", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "",
			Date:             "2020-01-01",
			Due:              "2020-01-02",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "1",
					Price:       "1",
				},
			},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrAmountEmpty {
			t.Errorf("Expected error %s, got %s", ErrAmountEmpty, err.Error())
		}
	})
	t.Run("should return error when amount is not numeric", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "a",
			Date:             "2020-01-01",
			Due:              "2020-01-02",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "1",
					Price:       "1",
				},
			},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrAmountInvalid {
			t.Errorf("Expected error %s, got %s", ErrAmountInvalid, err.Error())
		}
	})
	t.Run("should return error when amount is zero", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "0",
			Date:             "2020-01-01",
			Due:              "2020-01-02",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "1",
					Price:       "1",
				},
			},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrAmountZeroOrNegative {
			t.Errorf("Expected error %s, got %s", ErrAmountZeroOrNegative, err.Error())
		}
	})
	t.Run("should return error when amount is negative", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "-1",
			Date:             "2020-01-01",
			Due:              "2020-01-02",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "1",
					Price:       "1",
				},
			},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrAmountZeroOrNegative {
			t.Errorf("Expected error %s, got %s", ErrAmountZeroOrNegative, err.Error())
		}
	})
	t.Run("should return nil when date is empty", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "1",
			Date:             "",
			Due:              "2020-01-02",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrDateEmpty {
			t.Errorf("Expected error %s, got %s", ErrDateEmpty, err.Error())
		}
	})
	t.Run("should return error when date is invalid", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "1",
			Date:             "2020-01-32",
			Due:              "2020-01-02",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrDateInvalid {
			t.Errorf("Expected error %s, got %s", ErrDateInvalid, err.Error())
		}
	})
	t.Run("should return error when date is too old", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "1",
			Date:             "1999-12-31",
			Due:              "2020-01-02",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrOldDate {
			t.Errorf("Expected error %s, got %s", ErrOldDate, err.Error())
		}
	})
	t.Run("should return error when due date is empty", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "1.22",
			Date:             "2020-01-01",
			Due:              "",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrDueDateEmpty {
			t.Errorf("Expected error %s, got %s", ErrDueDateEmpty, err.Error())
		}
	})
	t.Run("should return error when due date is invalid", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "aasdsd",
			Amount:           "111",
			Date:             "2020-01-01",
			Due:              "2020-01-32",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "sddds",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrDueDateInvalid {
			t.Errorf("Expected error %s, got %s", ErrDueDateInvalid, err.Error())
		}
	})
	t.Run("should return error when due date is older than date", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "dsadsdsd",
			Amount:           "1",
			Date:             "2020-01-02",
			Due:              "2020-01-01",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "ddsds",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrDueDateOlderThanDate {
			t.Errorf("Expected error %s, got %s", ErrDueDateOlderThanDate, err.Error())
		}
	})
	t.Run("should return error when item reference is empty", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "dsadsdsd",
			Amount:           "1",
			Date:             "2020-01-02",
			Due:              "2020-01-03",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "",
					Description: "sdddsd",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		message := fmt.Sprintf("item 1: %s", ErrItemReferenceEmpty)
		if err != nil && err.Error() != message {
			t.Errorf("Expected error %s, got %s", message, err.Error())
		}
	})
	t.Run("should return a concatenated error when multiple fields are invalid", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "",
			BusinessNickname: "",
			CustomerNickname: "",
			Amount:           "",
			Date:             "",
			Due:              "",
			NoteReference:    "",
			Items:            []CreateInputItemDto{},
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && strings.Contains(ErrReferenceEmpty, err.Error()) {
			t.Errorf("Expected error %s, got %s", ErrReferenceEmpty, err.Error())
		}
		if err != nil && strings.Contains(ErrBusinessNicknameEmpty, err.Error()) {
			t.Errorf("Expected error %s, got %s", ErrBusinessNicknameEmpty, err.Error())
		}
		if err != nil && strings.Contains(ErrCustomerNicknameEmpty, err.Error()) {
			t.Errorf("Expected error %s, got %s", ErrCustomerNicknameEmpty, err.Error())
		}
		if err != nil && strings.Contains(ErrAmountEmpty, err.Error()) {
			t.Errorf("Expected error %s, got %s", ErrAmountEmpty, err.Error())
		}
		if err != nil && strings.Contains(ErrDateEmpty, err.Error()) {
			t.Errorf("Expected error %s, got %s", ErrDateEmpty, err.Error())
		}
		if err != nil && strings.Contains(ErrDueDateEmpty, err.Error()) {
			t.Errorf("Expected error %s, got %s", ErrDueDateEmpty, err.Error())
		}
	})
}

func TestGetReference(t *testing.T) {
	t.Run("should return reference", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "1.35 ",
			Date:             " 2020-01-01 ",
			Due:              " 2020-01-02 ",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		if ret := dto.GetReference(); ret != "ref" {
			t.Errorf("Expected ref, got %s", ret)
		}
	})
}

func TestGetBusinessNickname(t *testing.T) {
	t.Run("should return business nickname", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "1.35 ",
			Date:             " 2020-01-01 ",
			Due:              " 2020-01-02 ",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		if ret := dto.GetBusinessNickname(); ret != "business" {
			t.Errorf("Expected business, got %s", ret)
		}
	})
}

func TestGetCustomerNickname(t *testing.T) {
	t.Run("should return customer nickname", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "1.35 ",
			Date:             " 2020-01-01 ",
			Due:              " 2020-01-02 ",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		if ret := dto.GetCustomerNickname(); ret != "customer" {
			t.Errorf("Expected customer, got %s", ret)
		}
	})
}

func TestGetAmount(t *testing.T) {
	t.Run("should return amount", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "1.35",
			Date:             " 2020-01-01 ",
			Due:              " 2020-01-02 ",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		if amount, _ := dto.GetAmount(); amount != 1.35 {
			t.Errorf("Expected 1.35, got %f", amount)
		}
		//Should return error when amount is not numeric
		dto.Amount = "a"
		_, err := dto.GetAmount()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrAmountInvalid {
			t.Errorf("Expected error %s, got %s", ErrAmountInvalid, err.Error())
		}
	})
}

func TestGetDate(t *testing.T) {
	t.Run("should return date", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "customer",
			Amount:           "1.35",
			Date:             "2020-01-01",
			Due:              " 2020-01-02 ",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		if date, _ := dto.GetDate(); date.Format("2006-01-02") != "2020-01-01" {
			t.Errorf("Expected 2020-01-01, got %s", date.Format("2006-01-02"))
		}
		//Should return error when date is invalid
		dto.Date = "2020-01-32"
		_, err := dto.GetDate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrDateInvalid {
			t.Errorf("Expected error %s, got %s", ErrDateInvalid, err.Error())
		}
	})
}

func TestGetDue(t *testing.T) {
	t.Run("should return due date", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "",
			Amount:           "1.35",
			Date:             "2020-01-01",
			Due:              "2020-01-02",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		if due, _ := dto.GetDue(); due.Format("2006-01-02") != "2020-01-02" {
			t.Errorf("Expected 2020-01-02, got %s", due.Format("2006-01-02"))
		}
		//Should return error when due date is invalid
		dto.Due = "2020-01-32"
		_, err := dto.GetDue()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrDueDateInvalid {
			t.Errorf("Expected error %s, got %s", ErrDueDateInvalid, err.Error())
		}
	})
}

func TestGetNoteReference(t *testing.T) {
	t.Run("should return note reference", func(t *testing.T) {
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "",
			Amount:           "1.35",
			Date:             "2020-01-01",
			Due:              "2020-01-02",
			NoteReference:    "note",
			Items: []CreateInputItemDto{
				{
					Reference:   "ref",
					Description: "desc",
					Quantity:    "100",
					Price:       "1.23",
				},
			},
		}
		if ret := dto.GetNoteReference(); ret != "note" {
			t.Errorf("Expected note, got %s", ret)
		}
	})
}

func TestGetItems(t *testing.T) {
	t.Run("should return items", func(t *testing.T) {
		item := CreateInputItemDto{
			Reference:   "ref",
			Description: "desc",
			Quantity:    "100",
			Price:       "1.23",
		}
		dto := CreateInputDto{
			Reference:        "ref",
			BusinessNickname: "business",
			CustomerNickname: "",
			Amount:           "1.35",
			Date:             "2020-01-01",
			Due:              "2020-01-02",
			NoteReference:    "note",
			Items:            []CreateInputItemDto{item},
		}
		if ret := dto.GetItems(); ret[0] != item {
			t.Errorf("Expected item, got %s", ret[0])
		}
	})
}
