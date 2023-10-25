package repo

const (
	// SaveInvoiceClientQuery is the query to save a invoice client
	SaveInvoiceClient = "INSERT INTO invoice_client (id, nickname, client_id, name, document, phone, email) VALUES (?, ?, ?, ?, ?, ?, ?)"
	SaveInvoice       = "INSERT INTO invoice (id, reference, business_id, customer_id, amount, date, due, note_id, status_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	SaveInvoiceItem   = "INSERT INTO invoice_item (id, invoice_id, service_reference, description, amount, quantity) VALUES (?, ?, ?, ?, ?, ?)"
	TruncateInvoiceClient = "TRUNCATE TABLE invoice_client"
	TruncateInvoice       = "TRUNCATE TABLE invoice"
	TruncateInvoiceItem   = "TRUNCATE TABLE invoice_item"
)
