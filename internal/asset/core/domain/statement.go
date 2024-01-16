package domain

import "time"

var (
	// AssetStatementTypeMap is a map that represents the asset statement type within the statement
	HistoryMap = map[string]string{
		"FLOW":      "Contribution or withdrawal of the asset",
		"VALUATION": "Change in the value of the asset",
		"DIVIDENDS": "Dividends received from the asset",
		"DCLOSE":    "Daily close the asset",
	}
)

// AssetMovement is a struct that represents the asset movement
type Statement struct {
	ID        string    `gorm:"type:varchar(50); primaryKey"`
	AssetID   string    `gorm:"type:varchar(50); not null"`
	Asset     *Asset    `gorm:"foreignKey:AssetID;associationForeignKey:ID"`
	Date      time.Time `gorm:"type:date; not null"`
	History   string    `gorm:"type:varchar(50); not null"`
	Value     float64   `gorm:"type:decimal(20, 2); not null"`
	Comment   string    `gorm:"type:varchar(50); not null"`
	BalanceID string    `gorm:"type:varchar(50); not null"`
	Balance   *Balance  `gorm:"foreignKey:AssetBalanceID;associationForeignKey:ID"`
}
