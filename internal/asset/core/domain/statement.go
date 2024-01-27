package domain

import (
	"fmt"
	"time"

	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorStatementIDRequired        = "Statement ID is required"
	ErrorStatementIDLenght          = "Statement ID must have %d characters"
	ErrorStatementAssetIDRequired   = "Statement Asset ID is required"
	ErrorStatementAssetIDLenght     = "Statement Asset ID must have %d characters"
	ErrorAssetIDInvalid             = "Asset ID is invalid"
	ErrorStatementDateRequired      = "Statement Date is required"
	ErrorStatementHistoryRequired   = "Statement History is required"
	ErrorStatementHistoryLength     = "Statement History must have %d characters"
	ErrorStatementHistoryInvalid    = "Statement History is invalid"
	ErrorStatementValueRequired     = "Statement Value is required"
	ErrorStatementBalanceIDRequired = "Statement Balance ID is required"
	ErrorStatementBalanceID         = "Statement Balance ID must have %d characters"
	LengthStatementID               = 25
	LengthStatementHistory          = 2
	LengthStatementComment          = 100
)

var (
	// AssetStatementTypeMap is a map that represents the asset statement type within the statement
	HistoryMap = map[string]string{
		"FW": "Contribution or withdrawal of the asset",
		"VL": "Change in the value of the asset",
		"DV": "Dividends received from the asset",
		"CD": "Daily close the asset",
	}
)

// AssetMovement is a struct that represents the asset movement
type Statement struct {
	ID        string    `gorm:"type:varchar(25); primaryKey"`
	AssetID   string    `gorm:"type:varchar(25); not null"`
	Asset     *Asset    `gorm:"foreignKey:AssetID;associationForeignKey:ID"`
	Date      time.Time `gorm:"type:date; not null"`
	History   string    `gorm:"type:varchar(2); not null"`
	Value     float64   `gorm:"type:decimal(17, 2); not null"`
	Comment   string    `gorm:"type:varchar(100); null"`
	BalanceID string    `gorm:"type:varchar(25); not null"`
	Balance   *Balance  `gorm:"foreignKey:BalanceID;associationForeignKey:ID"`
}

// NewStatement creates a new asset statement line
func NewStatement(id, assetID string, date time.Time, history string, value float64, comment, balanceID string) *Statement {
	return &Statement{
		ID:        id,
		AssetID:   assetID,
		Date:      date,
		History:   history,
		Value:     value,
		Comment:   comment,
		BalanceID: balanceID,
	}
}

// SetCreate sets the asset create fields on create operation
func (s *Statement) SetCreate(repo port.Repository) *kerror.KError {
	if s.AssetID == "" {
		return kerror.NewKError(kerror.Internal, ErrorStatementAssetIDRequired)
	}
	s.Asset = &Asset{ID: s.AssetID}
	if ex, err := repo.GetByID(s.Asset, s.AssetID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if !ex {
		return kerror.NewKError(kerror.Internal, ErrorAssetIDInvalid)
	}
	return nil
}

// Validate validates the asset movement
func (s *Statement) Validate() *kerror.KError {
	if s.ID == "" {
		return kerror.NewKError(kerror.Internal, ErrorStatementIDRequired)
	}
	if len(s.ID) > LengthStatementID {
		return kerror.NewKError(kerror.Internal, fmt.Sprintf(ErrorStatementIDLenght, LengthStatementID))
	}
	if s.AssetID == "" {
		return kerror.NewKError(kerror.Internal, ErrorStatementAssetIDRequired)
	}
	if s.Asset == nil {
		return kerror.NewKError(kerror.Internal, ErrorAssetIDInvalid)
	}
	if len(s.AssetID) > LengthAssetID {
		return kerror.NewKError(kerror.Internal, fmt.Sprintf(ErrorStatementAssetIDLenght, LengthAssetID))
	}
	if s.Date.IsZero() {
		return kerror.NewKError(kerror.Internal, ErrorStatementDateRequired)
	}
	if s.History == "" {
		return kerror.NewKError(kerror.Internal, ErrorStatementHistoryRequired)
	}
	if len(s.History) > LengthStatementHistory {
		return kerror.NewKError(kerror.Internal, fmt.Sprintf(ErrorStatementHistoryLength, LengthStatementHistory))
	}
	if _, ok := HistoryMap[s.History]; !ok {
		return kerror.NewKError(kerror.Internal, ErrorStatementHistoryInvalid)
	}
	if s.Value == 0 {
		return kerror.NewKError(kerror.Internal, ErrorStatementValueRequired)
	}
	/*
	if s.BalanceID == "" {
		return kerror.NewKError(kerror.Internal, ErrorStatementBalanceIDRequired)
	}
	if len(s.BalanceID) > LengthBalanceID {
		return kerror.NewKError(kerror.Internal, fmt.Sprintf(ErrorStatementBalanceID, LengthBalanceID))
	}
	if s.Balance != nil {
		return s.Balance.Validate()
	}
	*/
	return nil
}

// TableName returns the table name for gorm
func (b *Statement) TableName() string {
	return "statement"
}
