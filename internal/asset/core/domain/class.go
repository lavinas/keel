package domain

import (
	"fmt"

	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorClassIDRequired    = "Class ID is required"
	ErrorClassIDLength      = "Class ID must have %d characters"
	ErrorClassNameRequired  = "Class Name is required"
	ErrorClassNameLength    = "Class Name must have %d characters"
	ErrorClassTaxIDRequired = "Class Tax ID is required"
	ErrorClassTaxIDLength   = "Class Tax ID must have %d characters"
	LengthClassID           = 25
	LengthClassName         = 50
)

// Class is a struct that represents the class of an asset
type Class struct {
	ID    string `gorm:"type:varchar(25); primaryKey"`
	Name  string `gorm:"type:varchar(50); not null"`
	TaxID string `gorm:"type:varchar(25); not null"`
	Tax   *Tax   `gorm:"foreignKey:TaxID;associationForeignKey:ID"`
}

// NewClass creates a new class
func NewClass(id, name, taxID string) *Class {
	return &Class{
		ID:    id,
		Name:  name,
		TaxID: taxID,
	}
}

// SetCreate sets the asset create fields on create operation
func (c *Class) SetCreate() *kerror.KError {
	return nil
}

// Validate validates the asset type
func (c *Class) Validate() *kerror.KError {
	if c.ID == "" {
		return kerror.NewKError(kerror.Internal, ErrorClassIDRequired)
	}
	if len(c.ID) > LengthClassID {
		return kerror.NewKError(kerror.Internal, fmt.Sprintf(ErrorClassIDLength, LengthClassID))
	}
	if c.Name == "" {
		return kerror.NewKError(kerror.Internal, ErrorClassNameRequired)
	}
	if len(c.Name) > LengthClassName {
		return kerror.NewKError(kerror.Internal, fmt.Sprintf(ErrorClassNameLength, LengthClassName))
	}
	if c.TaxID == "" {
		return kerror.NewKError(kerror.Internal, ErrorClassTaxIDRequired)
	}
	if len(c.TaxID) > LengthTaxID {
		return kerror.NewKError(kerror.Internal, fmt.Sprintf(ErrorClassTaxIDLength, LengthTaxID))
	}
	return nil
}

// TableName returns the table name for gorm
func (b *Class) TableName() string {
	return "class"
}
