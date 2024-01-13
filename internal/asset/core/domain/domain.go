package domain

import "time"

// TaxItem is a struct that represents the asset tax item per period
type AssetTaxItem struct {
	ID          string
	AssetTaxID  string
	PeriodType  string
	PeriodUntil int
	Tax         float64
}

// AssetTax is a struct that represents the asset tax
type AssetTax struct {
	ID   string
	Name string
	Tax  float64
}

// AssetType is a struct that represents the asset type
type AssetType struct {
	ID         string
	Name       string
	AssetTaxID string
}

// Asset is a struct that represents the asset
type Asset struct {
	ID          string
	AssetTypeID string
	Name        string
	StartDate   time.Time
	EndDate     time.Time
}

// AssetValue is a struct that represents the asset value in a specific date
type AssetHistory struct {
	ID             string
	AssetID        string
	Date           time.Time
	HistoryType    string
	HistoryValue   float64
	PrincipalValue float64
	ReturnValue    float64
	GrossValue     float64
	AsseTaxItemID  string
	TaxValue       float64
	NetValue       float64
}

// AssetPortfolio is a struct that represents the asset portfolio
type AssetPortfolio struct {
	ID   string
	Name string
}

// AssetPortfolioItem is a struct that represents the asset portfolio item
type AssetPortfolioItem struct {
	ID               string
	AssetPortfolioID string
	AssetID          string
}
