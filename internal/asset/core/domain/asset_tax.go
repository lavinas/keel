package domain

// AssetTax is a struct that represents the asset tax
type AssetTax struct {
	ID   string `json:"id"   gorm:"type:varchar(50); primaryKey"`
	Name string `json:"name" gorm:"type:varchar(50); not null"`
}

// TaxItem is a struct that represents the asset tax item per period
type AssetTaxItem struct {
	ID          string    `json:"id"           gorm:"type:varchar(50); primaryKey"`
	AssetTaxID  string    `json:"asset_tax_id" gorm:"type:varchar(50); not null"`
	AssetTax    *AssetTax `json:"asset_tax"    gorm:"foreignKey:AssetTaxID;associationForeignKey:ID"`
	PeriodType  string    `json:"period_type"  gorm:"type:varchar(50); not null"`
	PeriodUntil int       `json:"period_until" gorm:"type:int; null"`
	Tax         float64   `json:"tax"          gorm:"type:decimal(0, 4); not null"`
}
