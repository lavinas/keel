package dto

import (
	"fmt"

	"github.com/lavinas/keel/internal/asset/core/domain"
	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorTaxIDRequired           = "Tax ID is required"
	ErrorTaxIDLength             = "Tax ID must have less than %d characters"
	ErrorTaxIDDuplicated         = "Tax ID is duplicated"
	ErrorTaxNameRequired         = "Tax Name is required"
	ErrorTaxNameLength           = "Tax Name must have less than %d characters"
	ErrorTaxPeriodRequired       = "Tax Period is required"
	ErrorTaxPeriodLength         = "Tax Period must have less than %d characters"
	ErrorTaxPeriodInvalid        = "Tax Period is invalid"
	ErrorTaxItensRequired        = "Tax Itens is required"
	ErrorTaxItemIDRequired       = "Tax Item ID is required"
	ErrorTaxItemIDLength         = "Tax Item ID must have less than %d characters"
	ErrorTaxItemIDDuplicated     = "Tax Item ID is duplicated"
	ErrorTaxItemUntilInvalid     = "Tax Item Until is invalid"
	ErrorTaxItemValueInvalid     = "Tax Item Value is invalid"
	ErrorTaxItemTaxInvalidDomain = "Domain is not a tax"
)

var (
	// PeriodTypeMap is a map that represents the period type within the asset tax item
	PeriodMap = map[string]string{
		"Y": "Yearly",
		"M": "Monthly",
		"D": "Daily",
		"F": "Full",
		"N": "N/A",
	}
)

// AssetTaxItemIn is a struct that represents the asset tax item dto for input
type TaxItemCreate struct {
	ID    string  `json:"id"`
	Until int     `json:"until"`
	Value float64 `json:"value"`
}

// AssetTaxIn is a struct that represents the asset tax dto for input
type TaxCreateIn struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Period   string          `json:"period"`
	TaxItens []TaxItemCreate `json:"tax_itens"`
}

// AssetTaxOut is a struct that represents the asset tax dto for output
type TaxCreateOut struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	TaxItens []TaxItemCreate `json:"tax_itens"`
}

// Validate validates the asset tax dto for input
func (a *TaxCreateIn) Validate(repo port.Repository) *kerror.KError {
	valMap := []func(repo port.Repository) *kerror.KError{
		a.validateID,
		a.validateName,
		a.validatePeriod,
		a.validateTaxItens,
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

// GetDomain returns the asset tax domain for input
func (a *TaxCreateIn) GetDomain() (port.Domain, *kerror.KError) {
	taxItens := make([]*domain.TaxItem, 0)
	for _, taxItem := range a.TaxItens {
		taxItens = append(taxItens, domain.NewTaxItem(taxItem.ID, a.ID, taxItem.Until, taxItem.Value))
	}
	return domain.NewTax(a.ID, a.Name, a.Period, taxItens), nil
}

// validateID validates the id asset tax dto for input
func (a *TaxCreateIn) validateID(repo port.Repository) *kerror.KError {
	if a.ID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorTaxIDRequired)
	}
	if len(a.ID) > domain.LengthTaxID {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorTaxIDLength, domain.LengthTaxID))
	}
	if exists, err := repo.Exists(&domain.Tax{}, a.ID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if exists {
		return kerror.NewKError(kerror.BadRequest, ErrorTaxIDDuplicated)
	}
	return nil
}

// validateName validates the name asset tax dto for input
func (a *TaxCreateIn) validateName(repo port.Repository) *kerror.KError {
	if a.Name == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorTaxNameRequired)
	}
	if len(a.Name) > domain.LengthTaxName {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorTaxNameLength, domain.LengthTaxName))
	}
	return nil
}

// validatePeriod validates the period asset tax dto for input
func (a *TaxCreateIn) validatePeriod(repo port.Repository) *kerror.KError {
	if a.Period == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorTaxPeriodRequired)
	}
	if len(a.Period) > domain.LengthTaxPeriod {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorTaxPeriodLength, domain.LengthTaxPeriod))
	}
	if _, ok := PeriodMap[a.Period]; !ok {
		return kerror.NewKError(kerror.BadRequest, ErrorTaxPeriodInvalid)
	}
	return nil
}

// validateTaxItens validates the tax itens asset tax dto for input
func (a *TaxCreateIn) validateTaxItens(repo port.Repository) *kerror.KError {
	if len(a.TaxItens) == 0 {
		return kerror.NewKError(kerror.BadRequest, ErrorTaxItensRequired)
	}
	for _, taxItem := range a.TaxItens {
		if err := taxItem.validate(repo); err != nil {
			return err
		}
	}
	return nil
}

// validate validates the tax item asset tax dto for input
func (a *TaxItemCreate) validate(repo port.Repository) *kerror.KError {
	valMap := []func(repo port.Repository) *kerror.KError{
		a.validateID,
		a.validateUntil,
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

// validateID validates the id tax item asset tax dto for input
func (a *TaxItemCreate) validateID(repo port.Repository) *kerror.KError {
	if a.ID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorTaxItemIDRequired)
	}
	if len(a.ID) > domain.LengthTaxItemID {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorTaxItemIDLength, domain.LengthTaxItemID))
	}
	if exists, err := repo.Exists(&domain.TaxItem{}, a.ID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if exists {
		return kerror.NewKError(kerror.BadRequest, ErrorTaxItemIDDuplicated)
	}
	return nil
}

// validatePeriodUntil validates the period until tax item asset tax dto for input
func (a *TaxItemCreate) validateUntil(repo port.Repository) *kerror.KError {
	if a.Until <= 0 {
		return kerror.NewKError(kerror.BadRequest, ErrorTaxItemUntilInvalid)
	}
	return nil
}

// validateValue validates the value tax item asset tax dto for input
func (a *TaxItemCreate) validateValue(repo port.Repository) *kerror.KError {
	if a.Value < 0 || a.Value > 1 {
		return kerror.NewKError(kerror.BadRequest, ErrorTaxItemValueInvalid)
	}
	return nil
}

// SetDomain sets the asset tax domain for output
func (a *TaxCreateOut) SetDomain(d port.Domain) *kerror.KError {
	tax, ok := d.(*domain.Tax)
	if !ok {
		return kerror.NewKError(kerror.Internal, ErrorTaxItemTaxInvalidDomain)
	}
	taxItens := make([]TaxItemCreate, 0)
	for _, taxItem := range tax.TaxItems {
		taxItens = append(taxItens, TaxItemCreate{
			ID:    taxItem.ID,
			Until: taxItem.Until,
			Value: taxItem.Value,
		})
	}
	a.ID = tax.ID
	a.Name = tax.Name
	a.TaxItens = taxItens
	return nil
}
