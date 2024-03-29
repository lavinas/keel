package dto

import (
	"fmt"
	"time"

	"github.com/lavinas/keel/internal/asset/core/domain"
	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorAssetIDRequired        = "Asset ID is required"
	ErrorAssetIDLength          = "Asset ID must have %d characters"
	ErrorAssetIDDuplicated      = "Asset ID is duplicated"
	ErrorAssetClassIDRequired   = "Asset Class ID is required"
	ErrorAssetClassIDLength     = "Asset Class ID must have %d characters"
	ErrorAssetNameRequired      = "Asset Name is required"
	ErrorAssetNameLength        = "Asset Name must have %d characters"
	ErrorAssetStartDateRequired = "Asset Start Date is required"
	ErrorAssetStartDateInvalid  = "Asset Start Date is invalid"
	ErrorAssetClassIDNotFound   = "Asset Class ID is not found"
	ErrorAssetDomainInvalid     = "Domain is not an asset"
)

// AssetCreateIn is a struct that represents the asset dto for input creation
type AssetCreateIn struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ClassID   string `json:"class_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// AssetCreateOut is a struct that represents the asset dto for output creation
type AssetCreateOut struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ClassID   string `json:"class_id"`
	ClassName string `json:"class_name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// Validate validates the asset dto for input creation
func (a *AssetCreateIn) Validate(repo port.Repository) *kerror.KError {
	valMap := []func(repo port.Repository) *kerror.KError{
		a.validateID,
		a.validateClassID,
		a.validateName,
		a.validateStartDate,
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

// GetDomain returns the asset domain for input creation
func (a *AssetCreateIn) GetDomain() (port.Domain, *kerror.KError) {
	startDate, err := time.Parse("2006-01-02", a.StartDate)
	if err != nil {
		return nil, kerror.NewKError(kerror.Internal, err.Error())
	}
	var endDate *time.Time
	if a.EndDate != "" {
		endDateValue, err := time.Parse("2006-01-02", a.EndDate)
		if err != nil {
			return nil, kerror.NewKError(kerror.Internal, err.Error())
		}
		endDate = &endDateValue
	}
	asset := domain.NewAsset(a.ID, a.ClassID, a.Name, startDate, endDate)
	return asset, nil
}

// validateID validates the id asset dto for input creation
func (a *AssetCreateIn) validateID(repo port.Repository) *kerror.KError {
	if a.ID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorAssetIDRequired)
	}
	if len(a.ID) > domain.LengthAssetID {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorAssetIDLength, domain.LengthAssetID))
	}
	if exists, error := repo.Exists(&domain.Asset{}, a.ID); error != nil {
		return kerror.NewKError(kerror.Internal, error.Error())
	} else if exists {
		return kerror.NewKError(kerror.BadRequest, ErrorAssetIDDuplicated)
	}
	return nil
}

// validateClassID validates the class id asset dto for input creation
func (a *AssetCreateIn) validateClassID(repo port.Repository) *kerror.KError {
	if a.ClassID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorAssetClassIDRequired)
	}
	if len(a.ClassID) > domain.LengthClassID {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorAssetClassIDLength, domain.LengthClassID))
	}
	if exists, error := repo.Exists(&domain.Class{}, a.ClassID); error != nil {
		return kerror.NewKError(kerror.Internal, error.Error())
	} else if !exists {
		return kerror.NewKError(kerror.BadRequest, ErrorAssetClassIDNotFound)
	}
	return nil
}

// validateName validates the name asset dto for input creation
func (a *AssetCreateIn) validateName(repo port.Repository) *kerror.KError {
	if a.Name == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorAssetNameRequired)
	}
	if len(a.Name) > domain.LengthAssetName {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorAssetNameLength, domain.LengthAssetName))
	}
	return nil
}

// validateStartDate validates the start date asset dto for input creation
func (a *AssetCreateIn) validateStartDate(repo port.Repository) *kerror.KError {
	if a.StartDate == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorAssetStartDateRequired)
	}
	if _, err := time.Parse("2006-01-02", a.StartDate); err != nil {
		return kerror.NewKError(kerror.BadRequest, ErrorAssetStartDateInvalid)
	}
	return nil
}

// SetDomain sets the asset domain for output creation
func (a *AssetCreateOut) SetDomain(d port.Domain) *kerror.KError {
	asset, ok := d.(*domain.Asset)
	if !ok {
		return kerror.NewKError(kerror.Internal, ErrorAssetDomainInvalid)
	}
	a.ID = asset.ID
	a.Name = asset.Name
	a.ClassID = asset.ClassID
	a.ClassName = asset.Class.Name
	a.StartDate = asset.StartDate.Format("2006-01-02")
	if asset.EndDate != nil {
		a.EndDate = asset.EndDate.Format("2006-01-02")
	}
	return nil
}
