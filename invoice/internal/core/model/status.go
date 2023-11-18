package model

// InvoiceStatus represents a status of the invoice
type InvoiceStatus struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Validate validates the invoice status
func (i *InvoiceStatus) Validate() error {
	return nil
}

// PaymentStatus represents a status of the payment of the invoice
type PaymentStatus struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ValidatePaymentStatus validates the payment status
func (p *PaymentStatus) Validate() error {
	return nil
}