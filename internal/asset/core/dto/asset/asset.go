package dto

type AssetCreateIn struct {
	ID        string `json:"id"`
	ClassID   string `json:"class_id"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
