package domain

import "time"

// TaxItem is a struct that represents the asset tax item per period
type AssetTaxItem struct {
	ID          string    `json:"id"           gorm:"type:varchar(50); primaryKey"`
	AssetTaxID  string    `json:"asset_tax_id" gorm:"type:varchar(50); not null"`
	AssetTax    *AssetTax `json:"asset_tax"    gorm:"foreignKey:AssetTaxID;associationForeignKey:ID"`
	PeriodType  string    `json:"period_type"  gorm:"type:varchar(50); not null"`
	PeriodUntil int       `json:"period_until" gorm:"type:int; null"`
	Tax         float64   `json:"tax"          gorm:"type:decimal(0, 4); not null"`
}

// AssetTax is a struct that represents the asset tax
type AssetTax struct {
	ID   string `json:"id"   gorm:"type:varchar(50); primaryKey"`
	Name string `json:"name" gorm:"type:varchar(50); not null"`
}

// AssetType is a struct that represents the asset type
type AssetType struct {
	ID         string    `json:"id"           gorm:"type:varchar(50); primaryKey"`
	Name       string    `json:"name"         gorm:"type:varchar(50); not null"`
	AssetTaxID string    `json:"asset_tax_id" gorm:"type:varchar(50); not null"`
	AssetTax   *AssetTax `json:"asset_tax"    gorm:"foreignKey:AssetTaxID;associationForeignKey:ID"`
}

// Asset is a struct that represents the asset
type Asset struct {
	ID          string     `json:"id"             gorm:"type:varchar(50); primaryKey"`
	AssetTypeID string     `json:"asset_type_id"  gorm:"type:varchar(50); not null"`
	AssetType   *AssetType `json:"asset_type"     gorm:"foreignKey:AssetTypeID;associationForeignKey:ID"`
	Name        string     `json:"name"           gorm:"type:varchar(50); not null"`
	StartDate   time.Time  `json:"start_date"     gorm:"type:date; not null"`
	EndDate     time.Time  `json:"end_date"       gorm:"type:date; null"`
}

// AssetMovement is a struct that represents the asset movement
type AssetMovement struct {
	ID             string        `json:"id"                gorm:"type:varchar(50); primaryKey"`
	AssetID        string        `json:"asset_id"          gorm:"type:varchar(50); not null"`
	Asset          *Asset        `json:"asset"             gorm:"foreignKey:AssetID;associationForeignKey:ID"`
	Date           time.Time     `json:"date"              gorm:"type:date; not null"`
	MovementType   string        `json:"movement_type"     gorm:"type:varchar(50); not null"`
	MovementValue  float64       `json:"movement_value"    gorm:"type:decimal(20, 2); not null"`
	PrincipalValue float64       `json:"principal_value"   gorm:"type:decimal(20, 2); not null"`
	ReturnValue    float64       `json:"return_value"      gorm:"type:decimal(20, 2); not null"`
	GrossValue     float64       `json:"gross_value"       gorm:"type:decimal(20, 2); not null"`
	NetValue       float64       `json:"net_value"         gorm:"type:decimal(20, 2); not null"`
	TaxValue       float64       `json:"tax_value"         gorm:"type:decimal(20, 2); not null"`
	AsseTaxItemID  string        `json:"asset_tax_item_id" gorm:"type:varchar(50); not null"`
	AssetTaxItem   *AssetTaxItem `json:"asset_tax_item"    gorm:"foreignKey:AsseTaxItemID;associationForeignKey:ID"`
}

// AssetPortfolio is a struct that represents the asset portfolio
type AssetPortfolio struct {
	ID   string `json:"id"   gorm:"type:varchar(50); primaryKey"`
	Name string `json:"name" gorm:"type:varchar(50); not null"`
}

// AssetPortfolioItem is a struct that represents the asset portfolio item
type AssetPortfolioItem struct {
	AssetPortfolioID string          `json:"asset_portfolio_id" gorm:"type:varchar(50); primaryKey"`
	AssetPortfolio   *AssetPortfolio `json:"asset_portfolio"    gorm:"foreignKey:AssetPortfolioID;associationForeignKey:ID"`
	AssetID          string          `json:"asset_id"           gorm:"type:varchar(50); primaryKey"`
	Asset            *Asset          `json:"asset"              gorm:"foreignKey:AssetID;associationForeignKey:ID"`
}
