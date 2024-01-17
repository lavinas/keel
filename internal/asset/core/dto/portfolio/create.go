package dto

import (
	"github.com/lavinas/keel/internal/asset/core/domain"
	"github.com/lavinas/keel/internal/asset/core/port"
	"github.com/lavinas/keel/pkg/kerror"
)

const (
	ErrorPortfolioIDRequired          = "Portfolio ID is required"
	ErrorPortfolioIDDuplicated        = "Portfolio ID is duplicated"
	ErrorPortfolioNameRequired        = "Portfolio Name is required"
	ErrorPortfolioItemsRequired       = "Portfolio Items is required"
	ErrorPortfolioItemIDRequired      = "Portfolio Item ID is required"
	ErrorPortfolioItemIDNotFound      = "Portfolio Item ID is not found"
	ErrorPortfolioItemAssetIDRequired = "Portfolio Item Asset ID is required"
	ErrorPortfolioItemAssetIDNotFound = "Portfolio Item Asset ID is not found"
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
	PortfolioID string `json:"portfolio_id"`
	AssetID     string `json:"asset_id"`
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

// validateID validates the id asset portfolio dto for input
func (a *PortfolioCreateIn) validateID(repo port.Repository) *kerror.KError {
	if a.ID == "" {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioIDRequired)
	}
	if exists, err := repo.Exists(&domain.Portfolio{}, a.ID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if exists {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioIDDuplicated)
	}
	return nil
}

// validateName validates the name asset portfolio dto for input
func (a *PortfolioCreateIn) validateName(repo port.Repository) *kerror.KError {
	if a.Name == "" {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioNameRequired)
	}
	return nil
}

// validatePortfolioItems validates the portfolio items asset portfolio dto for input
func (a *PortfolioCreateIn) validatePortfolioItems(repo port.Repository) *kerror.KError {
	if len(a.PortfolioItems) == 0 {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioItemsRequired)
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
		api.validatePortfolioID,
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

// validatePortfolioID validates the portfolio id asset portfolio item dto for input
func (api *PortfolioItemCreate) validatePortfolioID(repo port.Repository) *kerror.KError {
	if api.PortfolioID == "" {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioItemIDRequired)
	}
	if exists, err := repo.Exists(&domain.Portfolio{}, api.PortfolioID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if !exists {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioItemIDNotFound)
	}
	return nil
}

// validateAssetID validates the asset id asset portfolio item dto for input
func (api *PortfolioItemCreate) validateAssetID(repo port.Repository) *kerror.KError {
	if api.AssetID == "" {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioItemAssetIDRequired)
	}
	if exists, err := repo.Exists(&domain.Asset{}, api.AssetID); err != nil {
		return kerror.NewKError(kerror.Internal, err.Error())
	} else if !exists {
		return kerror.NewKError(kerror.Internal, ErrorPortfolioItemAssetIDNotFound)
	}
	return nil
}
