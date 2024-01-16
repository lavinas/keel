package domain

import (
	"time"

	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorStatementIDRequired        = "Statement ID is required"
	ErrorStatementAssetIDRequired   = "Statement Asset ID is required"
	ErrorStatementDateRequired      = "Statement Date is required"
	ErrorStatementHistoryRequired   = "Statement History is required"
	ErrorStatementHistoryInvalid    = "Statement History is invalid"
	ErrorStatementValueRequired     = "Statement Value is required"
	ErrorStatementBalanceIDRequired = "Statement Balance ID is required"
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
	Balance   *Balance  `gorm:"foreignKey:AssetBalanceID;associationForeignKey:ID"`
}

// Validate validates the asset movement
func (s *Statement) Validate() *kerror.KError {
	if s.ID == "" {
		return kerror.NewKError(kerror.Internal, ErrorStatementIDRequired)
	}
	if s.AssetID == "" {
		return kerror.NewKError(kerror.Internal, ErrorStatementAssetIDRequired)
	}
	if s.Date.IsZero() {
		return kerror.NewKError(kerror.Internal, ErrorStatementDateRequired)
	}
	if s.History == "" {
		return kerror.NewKError(kerror.Internal, ErrorStatementHistoryRequired)
	}
	if _, ok := HistoryMap[s.History]; !ok {
		return kerror.NewKError(kerror.Internal, ErrorStatementHistoryInvalid)
	}
	if s.Value == 0 {
		return kerror.NewKError(kerror.Internal, ErrorStatementValueRequired)
	}
	if s.BalanceID == "" {
		return kerror.NewKError(kerror.Internal, ErrorStatementBalanceIDRequired)
	}
	if s.Balance != nil {
		return s.Balance.Validate()
	}
	return nil
}
