package model

type InvoiceItemCreate struct {
	ProductID   string  `json:"product_id"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
}

// InvoiceCreate is the DTO for creating an invoice.
type InvoiceCreate struct {
	BusinessID    string        `json:"business_id"`
	Number        string        `json:"number"`
	CustomerID    string        `json:"customer_id"`
	Date          string        `json:"date"`
	Due           string        `json:"due"`
	Amount        string        `json:"amount"`
	Items         []InvoiceItem `json:"items"`
	InstructionID string        `json:"instruction_id"`
}
