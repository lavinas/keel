package dto

import (
	"strings"
	"testing"
)

func TestItemValidate(t *testing.T){
	t.Run("should return nil when all fields are valid", func(t *testing.T){
		dto := CreateInputItemDto{
			Reference:   " ref ",
			Description: "desc ",
			Quantity:    " 100 ",
			Price:       "1.23",
		}
		err := dto.Validate()
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}
	})

	t.Run("should return error when reference is empty", func(t *testing.T){
		dto := CreateInputItemDto{
			Reference:   "",
			Description: "desc",
			Quantity:    "1",
			Price:       "1",
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrItemReferenceEmpty {
			t.Errorf("Expected error %s, got %s", ErrItemReferenceEmpty, err.Error())
		}
	})

	t.Run("should return error when description is empty", func(t *testing.T){
		dto := CreateInputItemDto{
			Reference:   "ref",
			Description: "",
			Quantity:    "1",
			Price:       "1",
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrItemDescriptionEmpty {
			t.Errorf("Expected error %s, got %s", ErrItemDescriptionEmpty, err.Error())
		}
	})

	t.Run("should return error when quantity is empty", func(t *testing.T){
		dto := CreateInputItemDto{
			Reference:   "ref",
			Description: "desc",
			Quantity:    "",
			Price:       "1",
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrItemQuantityEmpty {
			t.Errorf("Expected error %s, got %s", ErrItemQuantityEmpty, err.Error())
		}
	})

	t.Run("should return error when quantity is not numeric", func(t *testing.T){
		dto := CreateInputItemDto{
			Reference:   "ref",
			Description: "desc",
			Quantity:    "a",
			Price:       "1",
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrItemQuantityNotNumeric {
			t.Errorf("Expected error %s, got %s", ErrItemQuantityNotNumeric, err.Error())
		}
	})

	t.Run("should return error when quantity is zero", func(t *testing.T){
		dto := CreateInputItemDto{
			Reference:   "ref",
			Description: "desc",
			Quantity:    "0",
			Price:       "1",
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrItemQuantidyZeroNegative {
			t.Errorf("Expected error %s, got %s", ErrItemQuantidyZeroNegative, err.Error())
		}
	})
	t.Run("should return error when quantity is negative", func(t *testing.T){
		dto := CreateInputItemDto{
			Reference:   "ref",
			Description: "desc",
			Quantity:    "-1",
			Price:       "1",
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrItemQuantidyZeroNegative {
			t.Errorf("Expected error %s, got %s", ErrItemQuantidyZeroNegative, err.Error())
		}
	})
	t.Run("should return error when quantity is too big", func(t *testing.T){
		dto := CreateInputItemDto{
			Reference:   "ref",
			Description: "desc",
			Quantity:    "100000000000",
			Price:       "1",
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrItemQuantityMax {
			t.Errorf("Expected error %s, got %s", ErrItemQuantityMax, err.Error())
		}
	})

	t.Run("should return error when price is empty", func(t *testing.T){
		dto := CreateInputItemDto{
			Reference:   "ref",
			Description: "desc",
			Quantity:    "1",
			Price:       "",
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrItemPriceEmpty {
			t.Errorf("Expected error %s, got %s", ErrItemPriceEmpty, err.Error())
		}
	})
	
	t.Run("should return error when price is not numeric", func(t *testing.T){
		dto := CreateInputItemDto{
			Reference:   "ref",
			Description: "desc",
			Quantity:    "1",
			Price:       "a",
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrItemPriceNotNumeric {
			t.Errorf("Expected error %s, got %s", ErrItemPriceNotNumeric, err.Error())
		}
	})

	t.Run("should return error when price is zero", func(t *testing.T){
		dto := CreateInputItemDto{
			Reference:   "ref",
			Description: "desc",
			Quantity:    "1",
			Price:       "0",
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrItemPriceZeroNegative {
			t.Errorf("Expected error %s, got %s", ErrItemPriceZeroNegative, err.Error())
		}

	})
	t.Run("should return error when price is negative", func(t *testing.T){
		dto := CreateInputItemDto{
			Reference:   "ref",
			Description: "desc",
			Quantity:    "1",
			Price:       "-1",
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrItemPriceZeroNegative {
			t.Errorf("Expected error %s, got %s", ErrItemPriceZeroNegative, err.Error())
		}
	})
	t.Run("should return error when price is too big", func(t *testing.T){
		dto := CreateInputItemDto{
			Reference:   "ref",
			Description: "desc",
			Quantity:    "1",
			Price:       "100000000000",
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err != nil && err.Error() != ErrItemPriceMax {
			t.Errorf("Expected error %s, got %s", ErrItemPriceMax, err.Error())
		}
	})
	t.Run("should return a combination of errors", func(t *testing.T){
		dto := CreateInputItemDto{
			Reference:   "",
			Description: "",
			Quantity:    "",
			Price:       "",
		}
		err := dto.Validate()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if strings.Contains(err.Error(), ErrItemReferenceEmpty) == false {
			t.Errorf("Expected error %s, got %s", ErrItemReferenceEmpty, err.Error())
		}
		if strings.Contains(err.Error(), ErrItemDescriptionEmpty) == false {
			t.Errorf("Expected error %s, got %s", ErrItemDescriptionEmpty, err.Error())
		}
		if strings.Contains(err.Error(), ErrItemQuantityEmpty) == false {
			t.Errorf("Expected error %s, got %s", ErrItemQuantityEmpty, err.Error())
		}
		if strings.Contains(err.Error(), ErrItemPriceEmpty) == false {
			t.Errorf("Expected error %s, got %s", ErrItemPriceEmpty, err.Error())
		}
	})
}
