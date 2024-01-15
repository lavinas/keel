package domain

import "time"

// AssetMovement is a struct that represents the asset movement
type AssetStatement struct {
	ID             string        `gorm:"type:varchar(50); primaryKey"`
	AssetID        string        `gorm:"type:varchar(50); not null"`
	Asset          *Asset        `gorm:"foreignKey:AssetID;associationForeignKey:ID"`
	Date           time.Time     `gorm:"type:date; not null"`
	StatementType  string        `gorm:"type:varchar(50); not null"`
	StatementValue float64       `gorm:"type:decimal(20, 2); not null"`
	StatementDesc  string        `gorm:"type:varchar(50); not null"`
	PrincipalValue float64       `gorm:"type:decimal(20, 2); not null"`
	ReturnValue    float64       `gorm:"type:decimal(20, 2); not null"`
	GrossValue     float64       `gorm:"type:decimal(20, 2); not null"`
	NetValue       float64       `gorm:"type:decimal(20, 2); not null"`
	TaxValue       float64       `gorm:"type:decimal(20, 2); not null"`
	AsseTaxItemID  string        `gorm:"type:varchar(50); not null"`
	AssetTaxItem   *AssetTaxItem `gorm:"foreignKey:AsseTaxItemID;associationForeignKey:ID"`
}
