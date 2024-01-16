package domain

import "time"

// Asset is a struct that represents the asset
type Asset struct {
	ID        string    `gorm:"type:varchar(50); primaryKey"`
	ClassID   string    `gorm:"type:varchar(50); not null"`
	Class     *Class    `gorm:"foreignKey:AssetTypeID;associationForeignKey:ID"`
	Name      string    `gorm:"type:varchar(50); not null"`
	StartDate time.Time `gorm:"type:date; not null"`
	EndDate   time.Time `gorm:"type:date; null"`
	BalanceID string    `gorm:"type:varchar(50); null"`
	Balance   *Balance  `gorm:"foreignKey:AssetBalanceID;associationForeignKey:ID"`
}
