package domain

// AssetPortfolio is a struct that represents the asset portfolio
type AssetPortfolio struct {
	ID   string `json:"id"   gorm:"type:varchar(50); primaryKey"`
	Name string `json:"name" gorm:"type:varchar(50); not null"`
}

// AssetPortfolioItem is a struct that represents the asset portfolio item
type AssetPortfolioItem struct {
	AssetPortfolioID string          `json:"asset_portfolio_id" gorm:"type:varchar(50); primaryKey"`
	AssetPortfolio   *AssetPortfolio `json:"asset_portfolio"    gorm:"foreignKey:AssetPortfolioID;associationForeignKey:ID"`
	AssetID          string          `json:"asset_id"           gorm:"type:varchar(50); primaryKey"`
	Asset            *Asset          `json:"asset"              gorm:"foreignKey:AssetID;associationForeignKey:ID"`
}
