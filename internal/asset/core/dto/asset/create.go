package dto

import (
	"time"

	"github.com/lavinas/keel/internal/asset/core/domain"
	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorAssetIDRequired        = "Asset ID is required"
	ErrorAssetIDDuplicated      = "Asset ID is duplicated"
	ErrorAssetClassIDRequired   = "Asset Class ID is required"
	ErrorAssetNameRequired      = "Asset Name is required"
	ErrorAssetStartDateRequired = "Asset Start Date is required"
	ErrorAssetStartDateInvalid  = "Asset Start Date is invalid"
	ErrorAssetClassIDNotFound   = "Asset Class ID is not found"
)

// AssetCreateIn is a struct that represents the asset dto for input creation
type AssetCreateIn struct {
	ID        string `json:"id"`
	ClassID   string `json:"class_id"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// AssetCreateOut is a struct that represents the asset dto for output creation
type AssetCreateOut struct {
	ID        string `json:"id"`
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

// validateID validates the id asset dto for input creation
func (a *AssetCreateIn) validateID(repo port.Repository) *kerror.KError {
	if a.ID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorAssetIDRequired)
	}
	if exists, error := repo.Exists(domain.Asset{}, a.ID); error != nil {
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
	if exists, error := repo.Exists(domain.Class{}, a.ClassID); error != nil {
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

// ToDomain converts the asset dto for input creation to domain
func (a *AssetCreateIn) ToDomain() *domain.Asset {
	startDate, _ := time.Parse("2006-01-02", a.StartDate)
	var endDate *time.Time = nil
	if a.EndDate != "" {
		ed, _ := time.Parse("2006-01-02", a.EndDate)
		endDate = &ed
	}
	return &domain.Asset{
		ID:        a.ID,
		ClassID:   a.ClassID,
		Name:      a.Name,
		StartDate: startDate,
		EndDate:   endDate,
	}
}
