package dto

import (
	"time"

	"github.com/lavinas/keel/internal/asset/core/domain"
	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorStatementIDRequired      = "Statement ID is required"
	ErrorStatementIDDuplicated    = "Statement ID is duplicated"
	ErrorStatementAssetIDRequired = "Statement Asset ID is required"
	ErrorStatementAssetIDNotFound = "Statement Asset ID is not found"
	ErrorStatementDateRequired    = "Statement Date is required"
	ErrorStatementDateInvalid     = "Statement Date is invalid"
	ErrorStatementHistoryRequired = "Statement History is required"
	ErrorStatementHistoryInvalid  = "Statement History is invalid"
	ErrorStatementValueRequired   = "Statement Value is required"
)

var (
	// AssetStatementTypeMap is a map that represents the asset statement type within the statement
	HistoryMap = map[string]string{
		"FW": "Contribution or withdrawal of the asset",
		"VL": "Change in the value of the asset",
		"DV": "Dividends received from the asset",
		"CD": "Daily close of asset",
	}
)

// AssetStatementIn is a struct that represents the asset statement dto for input
type StatementCreateIn struct {
	ID      string  `json:"id"`
	AssetID string  `json:"asset_id"`
	Date    string  `json:"date"`
	History string  `json:"history"`
	Value   float64 `json:"value"`
	Comment string  `json:"comment"`
}

// AssetStatementOut is a struct that represents the asset statement dto for output
type StatementCreateOut struct {
	ID        string  `json:"id"`
	AssetID   string  `json:"asset_id"`
	AssetName string  `json:"asset_name"`
	Date      string  `json:"date"`
	History   string  `json:"history"`
	Value     float64 `json:"value"`
	Comment   string  `json:"comment"`
}

// Validate validates the asset statement dto for input
func (a *StatementCreateIn) Validate(repo port.Repository) *kerror.KError {
	valMap := []func(repo port.Repository) *kerror.KError{
		a.validateID,
		a.validateAssetID,
		a.validateDate,
		a.validateHistory,
		a.validateValue,
	}
	ret := kerror.NewKError(kerror.None, "")
	for _, val := range valMap {
		if err := val(repo); err != nil {
			ret.JoinKError(err)
		}
	}
	if !ret.IsEmpty() {
		return ret
	}
	return nil
}

// validateID validates the id asset statement dto for input
func (a *StatementCreateIn) validateID(repo port.Repository) *kerror.KError {
	if a.ID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorStatementIDRequired)
	}
	if exists, err := repo.Exists(&domain.Statement{}, a.ID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if exists {
		return kerror.NewKError(kerror.Internal, ErrorStatementIDDuplicated)
	}
	return nil
}

// validateAssetID validates the asset id asset statement dto for input
func (a *StatementCreateIn) validateAssetID(repo port.Repository) *kerror.KError {
	if a.AssetID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorStatementAssetIDRequired)
	}
	if exists, err := repo.Exists(&domain.Asset{}, a.AssetID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if !exists {
		return kerror.NewKError(kerror.BadRequest, ErrorStatementAssetIDNotFound)
	}
	return nil
}

// validateDate validates the date asset statement dto for input
func (a *StatementCreateIn) validateDate(repo port.Repository) *kerror.KError {
	if a.Date == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorStatementDateRequired)
	}
	if _, err := time.Parse("2006-01-02", a.Date); err != nil {
		return kerror.NewKError(kerror.BadRequest, ErrorStatementDateInvalid)
	}
	return nil
}

// validateHistory validates the history asset statement dto for input
func (a *StatementCreateIn) validateHistory(repo port.Repository) *kerror.KError {
	if a.History == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorStatementHistoryRequired)
	}
	if _, ok := HistoryMap[a.History]; !ok {
		return kerror.NewKError(kerror.BadRequest, ErrorStatementHistoryInvalid)
	}
	return nil
}

// validateValue validates the value asset statement dto for input
func (a *StatementCreateIn) validateValue(repo port.Repository) *kerror.KError {
	if a.Value == 0 {
		return kerror.NewKError(kerror.BadRequest, ErrorStatementValueRequired)
	}
	return nil
}

// ToDomain converts the asset statement dto to domain
func (a *StatementCreateIn) ToDomain(repo port.Repository) *domain.Statement {
	date, _ := time.Parse("2006-01-02", a.Date)
	
	return &domain.Statement{
		ID:      a.ID,
		AssetID: a.AssetID,
		Date:    date,
		History: a.History,
		Value:   a.Value,
		Comment: a.Comment,
	}
}
