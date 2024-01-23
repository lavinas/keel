package domain

import (
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorPortfolioIDRequired          = "Portfolio ID is required"
	ErrorPortfolioNameRequired        = "Portfolio Name is required"
	ErrorPortfolioItemIDRequired      = "Portfolio Item ID is required"
	ErrorPortfolioItemAssetIDRequired = "Portfolio Item Asset ID is required"
)

// AssetPortfolio is a struct that represents the asset portfolio
type Portfolio struct {
	ID             string           `gorm:"type:varchar(25); primaryKey"`
	Name           string           `gorm:"type:varchar(50); not null"`
	PortfolioItems []*PortfolioItem `gorm:"foreignKey:PortfolioID;associationForeignKey:ID"`
}

// AssetPortfolioItem is a struct that represents the asset portfolio item
type PortfolioItem struct {
	PortfolioID string     `gorm:"type:varchar(50); primaryKey"`
	Portfolio   *Portfolio `gorm:"foreignKey:AssetPortfolioID;associationForeignKey:ID"`
	AssetID     string     `gorm:"type:varchar(50); primaryKey"`
	Asset       *Asset     `gorm:"foreignKey:AssetID;associationForeignKey:ID"`
}

// NewAssetPortfolio creates a new asset portfolio
func NewPortfolio(id, name string, items []*PortfolioItem) *Portfolio {
	return &Portfolio{
		ID:             id,
		Name:           name,
		PortfolioItems: items,
	}
}

// NewAssetPortfolioItem creates a new asset portfolio item
func NewPortfolioItem(portfolioID, assetID string) *PortfolioItem {
	return &PortfolioItem{
		PortfolioID: portfolioID,
		AssetID:     assetID,
	}
}

// Validate validates the asset portfolio
func (p *Portfolio) Validate() *kerror.KError {
	if p.ID == "" {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioIDRequired)
	}
	if p.Name == "" {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioNameRequired)
	}
	return nil
}

// Validate validates the asset portfolio item
func (api *PortfolioItem) Validate() *kerror.KError {
	if api.PortfolioID == "" {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioItemIDRequired)
	}
	if api.AssetID == "" {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioItemAssetIDRequired)
	}
	return nil
}
