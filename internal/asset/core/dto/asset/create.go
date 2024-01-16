package dto

import (
	"time"

	"github.com/lavinas/keel/pkg/kerror"
	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/internal/asset/core/domain"
	
)

const (
	ErrorAssetIDRequired        = "Asset ID is required"
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
	if a.ID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorAssetIDRequired)
	}
	if a.ClassID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorAssetClassIDRequired)
	}

	if a.Name == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorAssetNameRequired)
	}
	if a.StartDate == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorAssetStartDateRequired)
	}
	if _, err := time.Parse("2006-01-02", a.StartDate); err != nil {
		return kerror.NewKError(kerror.BadRequest, ErrorAssetStartDateInvalid)
	}
	if exists, error := repo.Exists(domain.Class{}, a.ClassID); error != nil {
		return kerror.NewKError(kerror.Internal, error.Error())
	} else if !exists {
		return kerror.NewKError(kerror.BadRequest, ErrorAssetClassIDNotFound)
	}
	return nil
}
