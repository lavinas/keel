package dto

import (
	"fmt"

	"github.com/lavinas/keel/internal/asset/core/domain"
	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorPortfolioIDRequired          = "Portfolio ID is required"
	ErrorPortfolioIDLength            = "Portfolio ID must have %d characters"
	ErrorPortfolioIDDuplicated        = "Portfolio ID is duplicated"
	ErrorPortfolioNameRequired        = "Portfolio Name is required"
	ErrorPortfolioNameLength          = "Portfolio Name must have %d characters"
	ErrorPortfolioItemsRequired       = "Portfolio Items is required"
	ErrorPortfolioItemIDRequired      = "Portfolio Item ID is required"
	ErrorPortfolioItemIDLength        = "Portfolio Item ID must have %d characters"
	ErrorPortfolioItemIDNotFound      = "Portfolio Item ID is not found"
	ErrorPortfolioItemAssetIDRequired = "Portfolio Item Asset ID is required"
	ErrorPortfolioItemAssetIDLength   = "Portfolio Item Asset ID must have %d characters"
	ErrorPortfolioItemAssetIDNotFound = "Portfolio Item Asset ID is not found"
	ErrorPortfolioDomainInvalid       = "Domain is not a portfolio"
)

// PortfolioCreateIn is a struct that represents the asset portfolio dto for input
type PortfolioCreateIn struct {
	ID             string                `json:"id"`
	Name           string                `json:"name"`
	PortfolioItems []PortfolioItemCreate `json:"portfolio_items"`
}

// PortfolioCreateOut is a struct that represents the asset portfolio dto for output
type PortfolioCreateOut struct {
	ID             string                `json:"id"`
	Name           string                `json:"name"`
	PortfolioItems []PortfolioItemCreate `json:"portfolio_items"`
}

// PortfolioItemCreate is a struct that represents the asset portfolio item dto for input and output
type PortfolioItemCreate struct {
	AssetID string `json:"asset_id"`
}

// Validate validates the asset portfolio dto for input
func (a *PortfolioCreateIn) Validate(repo port.Repository) *kerror.KError {
	valMap := []func(repo port.Repository) *kerror.KError{
		a.validateID,
		a.validateName,
		a.validatePortfolioItems,
	}
	ret := kerror.NewKError(kerror.None, "")
	for _, val := range valMap {
		if err := val(repo); err != nil {
			ret.JoinKError(err)
		}
	}
	if !ret.IsEmpty() {
		return ret
	}
	return nil
}

// GetDomain returns the asset portfolio domain for input
func (a *PortfolioCreateIn) GetDomain() (port.Domain, *kerror.KError) {
	items := make([]*domain.PortfolioItem, 0)
	for _, item := range a.PortfolioItems {
		portfolioItem := domain.NewPortfolioItem(a.ID, item.AssetID)
		items = append(items, portfolioItem)
	}
	return domain.NewPortfolio(a.ID, a.Name, items), nil
}

// validateID validates the id asset portfolio dto for input
func (a *PortfolioCreateIn) validateID(repo port.Repository) *kerror.KError {
	if a.ID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorPortfolioIDRequired)
	}
	if len(a.ID) > domain.LengthPortfolioID {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorPortfolioIDLength, domain.LengthPortfolioID))
	}
	if exists, err := repo.Exists(&domain.Portfolio{}, a.ID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if exists {
		return kerror.NewKError(kerror.BadRequest, ErrorPortfolioIDDuplicated)
	}
	return nil
}

// validateName validates the name asset portfolio dto for input
func (a *PortfolioCreateIn) validateName(repo port.Repository) *kerror.KError {
	if a.Name == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorPortfolioNameRequired)
	}
	if len(a.Name) > domain.LengthPortfolioName {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorPortfolioNameLength, domain.LengthPortfolioName))
	}
	return nil
}

// validatePortfolioItems validates the portfolio items asset portfolio dto for input
func (a *PortfolioCreateIn) validatePortfolioItems(repo port.Repository) *kerror.KError {
	if len(a.PortfolioItems) == 0 {
		return kerror.NewKError(kerror.BadRequest, ErrorPortfolioItemsRequired)
	}
	for _, item := range a.PortfolioItems {
		if err := item.Validate(repo); err != nil {
			return err
		}
	}
	return nil
}

// Validate validates the asset portfolio item dto for input
func (api *PortfolioItemCreate) Validate(repo port.Repository) *kerror.KError {
	valMap := []func(repo port.Repository) *kerror.KError{
		api.validateAssetID,
	}
	ret := kerror.NewKError(kerror.None, "")
	for _, val := range valMap {
		if err := val(repo); err != nil {
			ret.JoinKError(err)
		}
	}
	if !ret.IsEmpty() {
		return ret
	}
	return nil
}

// validateAssetID validates the asset id asset portfolio item dto for input
func (api *PortfolioItemCreate) validateAssetID(repo port.Repository) *kerror.KError {
	if api.AssetID == "" {
		return kerror.NewKError(kerror.BadRequest, ErrorPortfolioItemAssetIDRequired)
	}
	if len(api.AssetID) > domain.LengthAssetID {
		return kerror.NewKError(kerror.BadRequest, fmt.Sprintf(ErrorPortfolioItemAssetIDLength, domain.LengthAssetID))
	}
	if exists, err := repo.Exists(&domain.Asset{}, api.AssetID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if !exists {
		return kerror.NewKError(kerror.BadRequest, ErrorPortfolioItemAssetIDNotFound)
	}
	return nil
}

// SetDomain sets the asset portfolio domain for output
func (a *PortfolioCreateOut) SetDomain(d port.Domain) *kerror.KError {
	portfolio, ok := d.(*domain.Portfolio)
	if !ok {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioDomainInvalid)
	}
	items := make([]PortfolioItemCreate, 0)
	for _, item := range portfolio.PortfolioItems {
		items = append(items, PortfolioItemCreate{
			AssetID: item.AssetID,
		})
	}
	a.ID = portfolio.ID
	a.Name = portfolio.Name
	a.PortfolioItems = items
	return nil
}
