package domain

// AssetType is a struct that represents the asset type
type AssetClass struct {
	ID         string    `json:"id"           gorm:"type:varchar(50); primaryKey"`
	Name       string    `json:"name"         gorm:"type:varchar(50); not null"`
	AssetTaxID string    `json:"asset_tax_id" gorm:"type:varchar(50); not null"`
	AssetTax   *AssetTax `json:"asset_tax"    gorm:"foreignKey:AssetTaxID;associationForeignKey:ID"`
}
