package domain

import (
	"time"

	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorAssetIDRequired        = "Asset ID is required"
	ErrorAssetClassIDRequired   = "Asset Class ID is required"
	ErrorAssetNameRequired      = "Asset Name is required"
	ErrorAssetStartDateRequired = "Asset Start Date is required"
)

// Asset is a struct that represents the asset
type Asset struct {
	ID        string     `gorm:"type:varchar(25); primaryKey"`
	ClassID   string     `gorm:"type:varchar(25); not null"`
	Class     *Class     `gorm:"foreignKey:AssetTypeID;associationForeignKey:ID"`
	Name      string     `gorm:"type:varchar(50); not null"`
	StartDate time.Time  `gorm:"type:date; not null"`
	EndDate   *time.Time `gorm:"type:date; null"`
	BalanceID string     `gorm:"type:varchar(25); null"`
	Balance   *Balance   `gorm:"foreignKey:AssetBalanceID;associationForeignKey:ID"`
}

// Validate validates the asset
func (a *Asset) Validate() *kerror.KError {
	if a.ID == "" {
		return kerror.NewKError(kerror.Internal, ErrorAssetIDRequired)
	}
	if a.ClassID == "" {
		return kerror.NewKError(kerror.Internal, ErrorAssetClassIDRequired)
	}
	if a.Name == "" {
		return kerror.NewKError(kerror.Internal, ErrorAssetNameRequired)
	}
	if a.StartDate.IsZero() {
		return kerror.NewKError(kerror.Internal, ErrorAssetStartDateRequired)
	}
	if a.Balance != nil {
		return a.Balance.Validate()
	}
	return nil
}
