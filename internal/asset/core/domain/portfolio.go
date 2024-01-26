package domain

import (
	"fmt"

	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorPortfolioIDRequired          = "Portfolio ID is required"
	ErrorPortfolioIDLength            = "Portfolio ID must have %d characters"
	ErrorPortfolioNameRequired        = "Portfolio Name is required"
	ErrorPortfolioNameLength          = "Portfolio Name must have %d characters"
	ErrorPortfolioItemIDRequired      = "Portfolio Item ID is required"
	ErrorPortfolioItemIDLength        = "Portfolio Item ID must have %d characters"
	ErrorPortfolioItemAssetIDInvalid  = "Portfolio Item Asset ID is invalid"
	ErrorPortfolioItemAssetIDRequired = "Portfolio Item Asset ID is required"
	ErrorPortfolioItemAssetIDLength   = "Portfolio Item Asset ID must have %d characters"
	LengthPortfolioID                 = 25
	LengthPortfolioName               = 50
	LengthPortfolioItemID             = 25
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
	Portfolio   *Portfolio `gorm:"foreignKey:PortfolioID;associationForeignKey:ID"`
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

// SetCreate sets the asset create fields on create operation
func (p *Portfolio) SetCreate(repo port.Repository) *kerror.KError {
	for _, item := range p.PortfolioItems {
		if err := item.SetCreate(repo); err != nil {
			return err
		}
	}
	return nil
}

// Validate validates the asset portfolio
func (p *Portfolio) Validate() *kerror.KError {
	if p.ID == "" {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioIDRequired)
	}
	if len(p.ID) > LengthPortfolioID {
		return kerror.NewKError(kerror.Internal, fmt.Sprintf(ErrorPortfolioIDLength, LengthPortfolioID))
	}
	if p.Name == "" {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioNameRequired)
	}
	if len(p.Name) > LengthPortfolioName {
		return kerror.NewKError(kerror.Internal, fmt.Sprintf(ErrorPortfolioNameLength, LengthPortfolioName))
	}
	return nil
}

// Validate validates the asset portfolio item
func (api *PortfolioItem) Validate() *kerror.KError {
	if api.PortfolioID == "" {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioItemIDRequired)
	}
	if len(api.PortfolioID) > LengthPortfolioItemID {
		return kerror.NewKError(kerror.Internal, fmt.Sprintf(ErrorPortfolioItemIDLength, LengthPortfolioItemID))
	}
	if api.AssetID == "" {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioItemAssetIDRequired)
	}
	if api.Asset == nil {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioItemAssetIDInvalid)
	}
	if len(api.AssetID) > LengthAssetID {
		return kerror.NewKError(kerror.Internal, fmt.Sprintf(ErrorPortfolioItemAssetIDLength, LengthAssetID))
	}
	return nil
}

// SetCreate sets the asset create fields on create operation
func (api *PortfolioItem) SetCreate(repo port.Repository) *kerror.KError {
	if api.AssetID == "" {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioItemIDRequired)
	}
	if ex, err := repo.GetByID(api.Asset, api.AssetID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if !ex {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioItemAssetIDInvalid)
	}
	return nil
}

// TableName returns the table name for gorm
func (p *Portfolio) TableName() string {
	return "portfolio"
}

// TableName returns the table name for gorm
func (p *PortfolioItem) TableName() string {
	return "portfolio_item"
}
