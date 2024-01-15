package dto

var (
	// AssetStatementTypeMap is a map that represents the asset statement type within the statement
	StatementTypeMap = map[string]string{
		"FLOW":    "Contribution or withdrawal of the asset",
		"INCOME":  "Income or outcome of the asset",
		"CLOSING": "Close the asset",
	}
)

// AssetStatementIn is a struct that represents the asset statement dto for input
type AssetStatementIn struct {
	ID             string  `json:"id"`
	AssetID        string  `json:"asset_id"`
	Date           string  `json:"date"`
	StatementType  string  `json:"statement_type"`
	StatementValue float64 `json:"statement_value"`
	StatementDesc  string  `json:"statement_desc"`
}
