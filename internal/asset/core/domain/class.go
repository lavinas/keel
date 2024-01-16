package domain

// AssetType is a struct that represents the asset type
type Class struct {
	ID    string `gorm:"type:varchar(50); primaryKey"`
	Name  string `gorm:"type:varchar(50); not null"`
	TaxID string `gorm:"type:varchar(50); not null"`
	Tax   *Tax   `gorm:"foreignKey:AssetTaxID;associationForeignKey:ID"`
}
