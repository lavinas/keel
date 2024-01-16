package domain

var (
	// PeriodTypeMap is a map that represents the period type within the asset tax item
	PeriodTypeMap = map[string]string{
		"YEARLY":  "Yearly",
		"MONTHLY": "Monthly",
		"DAILY":   "Daily",
	}
)

// AssetTax is a struct that represents the asset tax
type Tax struct {
	ID   string `gorm:"type:varchar(50); primaryKey"`
	Name string `gorm:"type:varchar(50); not null"`
}

// TaxItem is a struct that represents the asset tax item per period
type TaxItem struct {
	ID          string  `gorm:"type:varchar(50); primaryKey"`
	TaxID       string  `gorm:"type:varchar(50); not null"`
	Tax         *Tax    `gorm:"foreignKey:AssetTaxID;associationForeignKey:ID"`
	PeriodType  string  `gorm:"type:varchar(50); not null"`
	PeriodUntil int     `gorm:"type:int; null"`
	Value       float64 `gorm:"type:decimal(0, 4); not null"`
}
