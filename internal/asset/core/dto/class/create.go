package dto

// AssetClassIn is a struct that represents the asset class dto for input
type ClassCreateIn struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	AssetTaxID string `json:"asset_tax_id"`
}
