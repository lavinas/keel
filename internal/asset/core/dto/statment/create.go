package dto

var (
	// AssetStatementTypeMap is a map that represents the asset statement type within the statement
	HistoryMap = map[string]string{
		"FLOW":      "Contribution or withdrawal of the asset",
		"VALUATION": "Change in the value of the asset",
		"DIVIDENDS": "Dividends received from the asset",
		"DCLOSE":    "Daily close the asset",
	}
)

// AssetStatementIn is a struct that represents the asset statement dto for input
type StatementCreateIn struct {
	ID      string  `json:"id"`
	AssetID string  `json:"asset_id"`
	Date    string  `json:"date"`
	History string  `json:"statement_type"`
	Value   float64 `json:"statement_value"`
	Comment string  `json:"statement_desc"`
}
