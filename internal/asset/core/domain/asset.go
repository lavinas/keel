package domain

import (
	"fmt"
	"time"

	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorAssetIDRequired        = "Asset ID is required"
	ErrorAssetIDLength          = "Asset ID must have %d characters"
	ErrorAssetClassIDInvalid    = "Asset Class ID is invalid"
	ErrorAssetClassIDRequired   = "Asset Class ID is required"
	ErrorAssetClassIDLength     = "Asset Class ID must have %d characters"
	ErrorAssetNameRequired      = "Asset Name is required"
	ErrorAssetNameLength        = "Asset Name must have %d characters"
	ErrorAssetStartDateRequired = "Asset Start Date is required"
	LengthAssetID               = 25
	LengthAssetName             = 50
)

// Asset is a struct that represents the asset
type Asset struct {
	ID            string     `gorm:"type:varchar(25); primaryKey"`
	ClassID       string     `gorm:"type:varchar(25); not null"`
	Class         *Class     `gorm:"foreignKey:ClassID;associationForeignKey:ID"`
	Name          string     `gorm:"type:varchar(50); not null"`
	StartDate     time.Time  `gorm:"type:date; not null"`
	EndDate       *time.Time `gorm:"type:date; null"`
	LastBalanceID string     `gorm:"type:varchar(25); null"`
}

// NewAsset creates a new asset
func NewAsset(id, classID, name string, startDate time.Time, endDate *time.Time) *Asset {
	return &Asset{
		ID:        id,
		ClassID:   classID,
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
	}
}

// SetCreate sets the asset create fields on create operation
func (a *Asset) SetCreate(repo port.Repository) *kerror.KError {
	if a.ID == "" {
		return kerror.NewKError(kerror.Internal, ErrorAssetIDRequired)
	}
	if ex, err := repo.GetByID(a.Class, a.ClassID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if !ex {
		return kerror.NewKError(kerror.Internal, ErrorAssetClassIDInvalid)
	}
	return nil
}

// Validate validates the asset
func (a *Asset) Validate() *kerror.KError {
	if a.ID == "" {
		return kerror.NewKError(kerror.Internal, ErrorAssetIDRequired)
	}
	if len(a.ID) > LengthAssetID {
		return kerror.NewKError(kerror.Internal, fmt.Sprintf(ErrorAssetIDLength, LengthAssetID))
	}
	if a.ClassID == "" {
		return kerror.NewKError(kerror.Internal, ErrorAssetClassIDRequired)
	}
	if a.Class == nil {
		return kerror.NewKError(kerror.Internal, ErrorAssetClassIDInvalid)
	}
	if len(a.ClassID) > LengthClassID {
		return kerror.NewKError(kerror.Internal, fmt.Sprintf(ErrorAssetClassIDLength, LengthClassID))
	}
	if a.Name == "" {
		return kerror.NewKError(kerror.Internal, ErrorAssetNameRequired)
	}
	if len(a.Name) > LengthAssetName {
		return kerror.NewKError(kerror.Internal, fmt.Sprintf(ErrorAssetNameLength, LengthAssetName))
	}
	if a.StartDate.IsZero() {
		return kerror.NewKError(kerror.Internal, ErrorAssetStartDateRequired)
	}
	return nil
}

// TableName returns the table name for gorm
func (b *Asset) TableName() string {
	return "asset"
}
