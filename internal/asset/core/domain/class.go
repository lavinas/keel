package domain

import (
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorClassIDRequired    = "Class ID is required"
	ErrorClassNameRequired  = "Class Name is required"
	ErrorClassTaxIDRequired = "Class Tax ID is required"
)

// Class is a struct that represents the class of an asset
type Class struct {
	ID    string `gorm:"type:varchar(25); primaryKey"`
	Name  string `gorm:"type:varchar(50); not null"`
	TaxID string `gorm:"type:varchar(25); not null"`
	Tax   *Tax   `gorm:"foreignKey:AssetTaxID;associationForeignKey:ID"`
}

// Validate validates the asset type
func (c *Class) Validate() *kerror.KError {
	if c.ID == "" {
		return kerror.NewKError(kerror.Internal, ErrorClassIDRequired)
	}
	if c.Name == "" {
		return kerror.NewKError(kerror.Internal, ErrorClassNameRequired)
	}
	if c.TaxID == "" {
		return kerror.NewKError(kerror.Internal, ErrorClassTaxIDRequired)
	}
	return nil
}
