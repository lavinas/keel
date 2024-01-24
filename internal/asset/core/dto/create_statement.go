package dto

import (
	"fmt"
	"time"

	"github.com/lavinas/keel/internal/asset/core/domain"
	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorStatementIDRequired      = "Statement ID is required"
	ErrorStatementIDLenght        = "Statement ID must have %d characters"
	ErrorStatementIDDuplicated    = "Statement ID is duplicated"
	ErrorStatementAssetIDRequired = "Statement Asset ID is required"
	ErrorStatementAssedIDLength   = "Statement Asset ID must have %d characters"
	ErrorStatementAssetIDNotFound = "Statement Asset ID is not found"
	ErrorStatementDateRequired    = "Statement Date is required"
	ErrorStatementDateInvalid     = "Statement Date is invalid"
	ErrorStatementHistoryRequired = "Statement History is required"
	ErrorStatementHistoryLength   = "Statement History must have %d characters"
	ErrorStatementHistoryInvalid  = "Statement History is invalid"
	ErrorStatementValueRequired   = "Statement Value is required"
	ErrorStatementDomainInvalid   = "Domain is not a statement"
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

// GetDomain returns the asset statement domain for input
func (a *StatementCreateIn) GetDomain() (port.Domain, *kerror.KError) {
	date, err := time.Parse("2006-01-02", a.Date)
	if err != nil {
		return nil, kerror.NewKError(kerror.Internal, err.Error())
	}
	statement := domain.NewStatement(a.ID, a.AssetID, date, a.History, a.Value, a.Comment, "")
	return statement, nil
}

// validateID validates the id asset statement dto for input
func (a *StatementCreateIn) validateID(repo port.Repository) *kerror.KError {
	if a.ID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorStatementIDRequired)
	}
	if len(a.ID) > domain.LengthStatementID {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorStatementIDLenght, domain.LengthStatementID))
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
	if len(a.AssetID) > domain.LengthAssetID {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorStatementAssedIDLength, domain.LengthAssetID))
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
	if len(a.History) > domain.LengthStatementHistory {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorStatementHistoryLength, domain.LengthStatementHistory))
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

// SetDomain sets the asset statement domain for output
func (a *StatementCreateOut) SetDomain(d port.Domain) *kerror.KError {
	statement, ok := d.(*domain.Statement)
	if !ok {
		return kerror.NewKError(kerror.Internal, ErrorStatementDomainInvalid)
	}
	a.ID = statement.ID
	a.AssetID = statement.AssetID
	a.AssetName = statement.Asset.Name
	a.Date = statement.Date.Format("2006-01-02")
	a.History = statement.History
	a.Value = statement.Value
	a.Comment = statement.Comment
	return nil
}
