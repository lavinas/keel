package dto

type AssetIn struct {
	ID           string `json:"id"`
	AssetClassID string `json:"asset_class_id"`
	Name         string `json:"name"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}
