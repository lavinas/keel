package domain

// AssetPortfolio is a struct that represents the asset portfolio
type Portfolio struct {
	ID   string `gorm:"type:varchar(50); primaryKey"`
	Name string `gorm:"type:varchar(50); not null"`
}

// AssetPortfolioItem is a struct that represents the asset portfolio item
type AssetPortfolioItem struct {
	PortfolioID string     `gorm:"type:varchar(50); primaryKey"`
	Portfolio   *Portfolio `gorm:"foreignKey:AssetPortfolioID;associationForeignKey:ID"`
	AssetID     string     `gorm:"type:varchar(50); primaryKey"`
	Asset       *Asset     `gorm:"foreignKey:AssetID;associationForeignKey:ID"`
}
