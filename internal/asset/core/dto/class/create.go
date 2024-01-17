package dto

import (
	"github.com/lavinas/keel/internal/asset/core/domain"
	"github.com/lavinas/keel/internal/asset/core/port"

	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorClassIDRequired         = "Class ID is required"
	ErrorClassIDDuplicated       = "Class ID is duplicated"
	ErrorClassNameRequired       = "Class Name is required"
	ErrorClassAssetTaxIDRequired = "Class Asset Tax ID is required"
	ErrorClassAssetTaxIDNotFound = "Class Asset Tax ID is not found"
)

// AssetClassIn is a struct that represents the asset class dto for input
type ClassCreateIn struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	TaxID string `json:"tax_id"`
}

type ClassCreateOut struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	TaxID   string `json:"tax_id"`
	TaxName string `json:"tax_name"`
}

// Validate validates the asset class dto for input
func (a *ClassCreateIn) Validate(repo port.Repository) *kerror.KError {
	valMap := []func(repo port.Repository) *kerror.KError{
		a.validateID,
		a.validateName,
		a.validateAssetTaxID,
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

// validateID validates the id asset class dto for input
func (a *ClassCreateIn) validateID(repo port.Repository) *kerror.KError {
	if a.ID == "" {
		return kerror.NewKError(kerror.Internal, ErrorClassIDRequired)
	}
	if exists, err := repo.Exists(&domain.Class{}, a.ID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if exists {
		return kerror.NewKError(kerror.Internal, ErrorClassIDDuplicated)
	}
	return nil
}

// validateName validates the name asset class dto for input
func (a *ClassCreateIn) validateName(repo port.Repository) *kerror.KError {
	if a.Name == "" {
		return kerror.NewKError(kerror.Internal, ErrorClassNameRequired)
	}
	return nil
}

// validateAssetTaxID validates the asset tax id asset class dto for input
func (a *ClassCreateIn) validateAssetTaxID(repo port.Repository) *kerror.KError {
	if a.TaxID == "" {
		return kerror.NewKError(kerror.Internal, ErrorClassAssetTaxIDRequired)
	}
	if exists, err := repo.Exists(&domain.Tax{}, a.TaxID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if !exists {
		return kerror.NewKError(kerror.Internal, ErrorClassAssetTaxIDNotFound)
	}
	return nil
}

// ToDomain converts the asset class dto to domain
func (a *ClassCreateIn) ToDomain() *domain.Class {
	return &domain.Class{
		ID:    a.ID,
		Name:  a.Name,
		TaxID: a.TaxID,
	}
}
