package domain

import (
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorTaxIDRequired        = "Tax ID is required"
	ErrorTaxNameRequired      = "Tax Name is required"
	ErrorTaxPeriodRequired    = "Tax Period is required"
	ErrorTaxPeriodInvalid     = "Tax Period is invalid"
	ErrorTaxItemIDRequired    = "Tax Item ID is required"
	ErrorTaxItemTaxIDRequired = "Tax Item Tax ID is required"
	ErrorTaxItemValueInvalid  = "Tax Item Value is invalid"
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

// AssetTax is a struct that represents the asset tax
type Tax struct {
	ID       string     `gorm:"type:varchar(25); primaryKey"`
	Name     string     `gorm:"type:varchar(50); not null"`
	Period   string     `gorm:"type:varchar(2); not null"`
	TaxItems []*TaxItem `gorm:"foreignKey:TaxID;associationForeignKey:ID"`
}

// TaxItem is a struct that represents the asset tax item per period
type TaxItem struct {
	ID    string  `gorm:"type:varchar(25); primaryKey"`
	TaxID string  `gorm:"type:varchar(25); not null"`
	Until int `gorm:"type:int; null"`
	Value float64 `gorm:"type:decimal(0, 4); not null"`
}

// NewTax creates a new asset tax
func NewTax(id, name, period string, taxItens []*TaxItem) *Tax {
	return &Tax{
		ID:       id,
		Name:     name,
		Period:   period,
		TaxItems: taxItens,
	}
}

// NewTaxItem creates a new asset tax item
func NewTaxItem(id, taxID string, until int, value float64) *TaxItem {
	return &TaxItem{
		ID:    id,
		TaxID: taxID,
		Until: until,
		Value: value,
	}
}

// Validate validates the asset tax
func (t *Tax) Validate() *kerror.KError {
	if t.ID == "" {
		return kerror.NewKError(kerror.Internal, ErrorTaxIDRequired)
	}
	if t.Name == "" {
		return kerror.NewKError(kerror.Internal, ErrorTaxNameRequired)
	}
	if t.Period == "" {
		return kerror.NewKError(kerror.Internal, ErrorTaxPeriodRequired)
	}
	if _, ok := PeriodMap[t.Period]; !ok {
		return kerror.NewKError(kerror.Internal, ErrorTaxPeriodInvalid)
	}
	for _, ti := range t.TaxItems {
		if err := ti.Validate(t.ID); err != nil {
			return err
		}
	}
	return nil
}

// Validate validates the asset tax item
func (ti *TaxItem) Validate(tax_id string) *kerror.KError {
	if ti.ID == "" {
		return kerror.NewKError(kerror.Internal, ErrorTaxItemIDRequired)
	}
	if ti.TaxID != tax_id {
		return kerror.NewKError(kerror.Internal, ErrorTaxItemTaxIDRequired)
	}
	if ti.Value < 0 || ti.Value > 1 {
		return kerror.NewKError(kerror.Internal, ErrorTaxItemValueInvalid)
	}
	return nil
}
