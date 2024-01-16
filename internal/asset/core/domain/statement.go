package domain

import "time"

var (
	// AssetStatementTypeMap is a map that represents the asset statement type within the statement
	StatementTypes = map[string]string{
		"FLOW":    "Contribution or withdrawal of the asset",
		"INCOME":  "Income or outcome of the asset",
		"CLOSING": "Close the asset",
	}
)

// AssetMovement is a struct that represents the asset movement
type Statement struct {
	ID             string    `gorm:"type:varchar(50); primaryKey"`
	AssetID        string    `gorm:"type:varchar(50); not null"`
	Asset          *Asset    `gorm:"foreignKey:AssetID;associationForeignKey:ID"`
	Date           time.Time `gorm:"type:date; not null"`
	StatementType  string    `gorm:"type:varchar(50); not null"`
	StatementValue float64   `gorm:"type:decimal(20, 2); not null"`
	StatementDesc  string    `gorm:"type:varchar(50); not null"`
	PrincipalValue float64   `gorm:"type:decimal(20, 2); not null"`
	ReturnValue    float64   `gorm:"type:decimal(20, 2); not null"`
	GrossValue     float64   `gorm:"type:decimal(20, 2); not null"`
	NetValue       float64   `gorm:"type:decimal(20, 2); not null"`
	TaxValue       float64   `gorm:"type:decimal(20, 2); not null"`
	TaxItemID      string    `gorm:"type:varchar(50); not null"`
	TaxItem        *TaxItem  `gorm:"foreignKey:AsseTaxItemID;associationForeignKey:ID"`
}
