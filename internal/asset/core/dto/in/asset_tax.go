package dto

var (
	// PeriodTypeMap is a map that represents the period type within the asset tax item
	PeriodTypeMap = map[string]string{
		"YEARLY":  "Yearly",
		"MONTHLY": "Monthly",
		"DAILY":   "Daily",
	}
)

// AssetTaxIn is a struct that represents the asset tax dto for input
type AssetTaxIn struct {
	ID            string           `json:"id"`
	Name          string           `json:"name"`
	AssetTaxItens []AssetTaxItemIn `json:"asset_tax_itens"`
}

// AssetTaxItemIn is a struct that represents the asset tax item dto for input
type AssetTaxItemIn struct {
	ID          string  `json:"id"`
	PeriodType  string  `json:"period_type"`
	PeriodUntil int     `json:"period_until"`
	Tax         float64 `json:"tax"`
}
