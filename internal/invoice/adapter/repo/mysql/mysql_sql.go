package mysql

// Queries
var (
	querieMap = map[string]string{
		"SaveInvoiceClient":     "INSERT INTO {DB}.invoice_client (id, nickname, client_id, name, document, phone, email) VALUES (?, ?, ?, ?, ?, ?, ?);",
		"SaveInvoice":           "INSERT INTO {DB}.invoice (id, reference, business_id, customer_id, amount, date, due, note_id, status_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
		"SaveInvoiceItem":       "INSERT INTO {DB}.invoice_item (id, invoice_id, service_reference, description, amount, quantity) VALUES (?, ?, ?, ?, ?, ?);",
		"TruncateInvoiceClient": "DELETE FROM {DB}.invoice_client;",
		"TruncateInvoice":       "DELETE FROM {DB}.invoice;",
		"TruncateInvoiceItem":   "DELETE FROM {DB}.invoice_item;",
	}
)
