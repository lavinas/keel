package dto

import (
	"fmt"

	"github.com/lavinas/keel/internal/asset/core/domain"
	"github.com/lavinas/keel/internal/asset/core/port"

	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorClassIDRequired         = "Class ID is required"
	ErrorClassIDLength           = "Class ID must have %d characters"
	ErrorClassIDDuplicated       = "Class ID is duplicated"
	ErrorClassNameRequired       = "Class Name is required"
	ErrorClassNameLength         = "Class Name must have %d characters"
	ErrorClassAssetTaxIDRequired = "Class Asset Tax ID is required"
	ErrorClassAssetTaxIDLength   = "Class Asset Tax ID must have %d characters"
	ErrorClassAssetTaxIDNotFound = "Class Asset Tax ID is not found"
	ErrorClassDomainInvalid      = "Domain is not an asset class"
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
		a.validateTaxID,
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

// GetDomain returns the asset class domain for input
func (a *ClassCreateIn) GetDomain() (port.Domain, *kerror.KError) {
	class := domain.NewClass(a.ID, a.Name, a.TaxID)
	return class, nil
}

// validateID validates the id asset class dto for input
func (a *ClassCreateIn) validateID(repo port.Repository) *kerror.KError {
	if a.ID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorClassIDRequired)
	}
	if len(a.ID) > domain.LengthClassID {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorClassIDLength, domain.LengthClassID))
	}
	if exists, err := repo.Exists(&domain.Class{}, a.ID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if exists {
		return kerror.NewKError(kerror.BadRequest, ErrorClassIDDuplicated)
	}
	return nil
}

// validateName validates the name asset class dto for input
func (a *ClassCreateIn) validateName(repo port.Repository) *kerror.KError {
	if a.Name == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorClassNameRequired)
	}
	if len(a.Name) > domain.LengthClassName {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorClassNameLength, domain.LengthClassName))
	}
	return nil
}

// validateAssetTaxID validates the asset tax id asset class dto for input
func (a *ClassCreateIn) validateTaxID(repo port.Repository) *kerror.KError {
	if a.TaxID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorClassAssetTaxIDRequired)
	}
	if len(a.TaxID) > domain.LengthTaxID {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorClassAssetTaxIDLength, domain.LengthTaxID))
	}
	if exists, err := repo.Exists(&domain.Tax{}, a.TaxID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if !exists {
		return kerror.NewKError(kerror.BadRequest, ErrorClassAssetTaxIDNotFound)
	}
	return nil
}

// SetDomain sets the asset class domain for output
func (a *ClassCreateOut) SetDomain(d port.Domain) *kerror.KError {
	class, ok := d.(*domain.Class)
	if !ok {
		return kerror.NewKError(kerror.Internal, ErrorClassDomainInvalid)
	}
	a.ID = class.ID
	a.Name = class.Name
	a.TaxID = class.TaxID
	a.TaxName = class.Tax.Name
	return nil
}
