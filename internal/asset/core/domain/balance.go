package domain

import (
	"time"

	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorBalanceIDRequired        = "Balance ID is required"
	ErrorBalanceDateRequired      = "Balance Date is required"
	ErrorBalanceAssetIDRequired   = "Balance Asset ID is required"
	ErrorBalanceTaxItemIDRequired = "Balance Tax Item ID is required"
	ErrorBalanceGrossValueInvalid = "Balance Principal Value + Return Value must be equal to Gross Value"
	ErrorBalanceNetValueInvalid   = "Balance Gross Value - Tax Value must be equal to Net Value"
)

type Balance struct {
	ID             string    `gorm:"type:varchar(25); primaryKey"`
	Date           time.Time `gorm:"type:date; not null"`
	AssetID        string    `gorm:"type:varchar(25); not null"`
	Asset          *Asset    `gorm:"foreignKey:AssetID;associationForeignKey:ID"`
	PrincipalValue float64   `gorm:"type:decimal(17, 2); not null"`
	ReturnValue    float64   `gorm:"type:decimal(17, 2); not null"`
	GrossValue     float64   `gorm:"type:decimal(17, 2); not null"`
	NetValue       float64   `gorm:"type:decimal(17, 2); not null"`
	TaxValue       float64   `gorm:"type:decimal(17, 2); not null"`
	TaxItemID      string    `gorm:"type:varchar(25); not null"`
	TaxItem        *TaxItem  `gorm:"foreignKey:AsseTaxItemID;associationForeignKey:ID"`
}

// NewBalance creates a new balance
func NewBalance(id string, date time.Time, assetID string, principalValue, returnValue, grossValue, netValue, taxValue float64, taxItemID string) *Balance {
	return &Balance{
		ID:             id,
		Date:           date,
		AssetID:        assetID,
		PrincipalValue: principalValue,
		ReturnValue:    returnValue,
		GrossValue:     grossValue,
		NetValue:       netValue,
		TaxValue:       taxValue,
		TaxItemID:      taxItemID,
	}
}

// Validate validates the balance
func (b *Balance) Validate() *kerror.KError {
	if b.ID == "" {
		return kerror.NewKError(kerror.Internal, ErrorBalanceIDRequired)
	}
	if b.Date.IsZero() {
		return kerror.NewKError(kerror.Internal, ErrorBalanceDateRequired)
	}
	if b.AssetID == "" {
		return kerror.NewKError(kerror.Internal, ErrorBalanceAssetIDRequired)
	}
	if b.TaxItemID == "" {
		return kerror.NewKError(kerror.Internal, ErrorBalanceTaxItemIDRequired)
	}
	if b.PrincipalValue+b.ReturnValue != b.GrossValue {
		return kerror.NewKError(kerror.Internal, ErrorBalanceGrossValueInvalid)
	}
	if b.GrossValue-b.TaxValue != b.NetValue {
		return kerror.NewKError(kerror.Internal, ErrorBalanceNetValueInvalid)
	}
	return nil
}
