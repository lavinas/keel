package domain

import "time"

// Asset is a struct that represents the asset
type Asset struct {
	ID           string      `json:"id"             gorm:"type:varchar(50); primaryKey"`
	AssetClassID string      `json:"asset_class_id" gorm:"type:varchar(50); not null"`
	AssetClass   *AssetClass `json:"asset_type"     gorm:"foreignKey:AssetTypeID;associationForeignKey:ID"`
	Name         string      `json:"name"           gorm:"type:varchar(50); not null"`
	StartDate    time.Time   `json:"start_date"     gorm:"type:date; not null"`
	EndDate      time.Time   `json:"end_date"       gorm:"type:date; null"`
}
