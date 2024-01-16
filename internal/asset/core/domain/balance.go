package domain

import (
	"time"
)

type Balance struct {
	ID             string    `gorm:"type:varchar(50); primaryKey"`
	Date           time.Time `gorm:"type:date; not null"`
	PrincipalValue float64   `gorm:"type:decimal(20, 2); not null"`
	ReturnValue    float64   `gorm:"type:decimal(20, 2); not null"`
	GrossValue     float64   `gorm:"type:decimal(20, 2); not null"`
	NetValue       float64   `gorm:"type:decimal(20, 2); not null"`
	TaxValue       float64   `gorm:"type:decimal(20, 2); not null"`
	TaxItemID      string    `gorm:"type:varchar(50); not null"`
	TaxItem        *TaxItem  `gorm:"foreignKey:AsseTaxItemID;associationForeignKey:ID"`
}
