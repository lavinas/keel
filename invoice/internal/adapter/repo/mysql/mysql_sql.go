package mysql

// Queries
var (
	querieMap = map[string]string{
		"SaveInvoiceClient":      "INSERT INTO {DB}.invoice_client (id, nickname, client_id, name, document, phone, email, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		"UpadateInvoiceClient":   "UPDATE {DB}.invoice_client SET nickname = ?, client_id = ?, name = ?, document = ?, phone = ?, email = ? WHERE id = ?;",
		"SaveInvoice":            "INSERT INTO {DB}.invoice (id, reference, business_id, customer_id, amount, date, due, note_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
		"SaveInvoiceItem":        "INSERT INTO {DB}.invoice_item (id, invoice_id, service_reference, description, amount, quantity) VALUES (?, ?, ?, ?, ?, ?);",
		"Truncate":               "DELETE FROM {DB}.{TABLE};",
		"IsDuplicatedInvoice":    "SELECT COUNT(*) count FROM {DB}.invoice WHERE reference = ?;",
		"GetInvoiceClient":       "SELECT id, nickname, client_id, name, document, phone, email, created_at FROM {DB}.invoice_client WHERE nickname = ? and created_at >= ? order by created_at desc limit 1;",
		"GetInvoiceVertex":       "SELECT class, id, name, description FROM {DB}.invoice_vertex;",
		"GetInvoiceEdge":         "SELECT class, from_invoice_vertex_id, to_invoice_vertex_id, description FROM {DB}.invoice_edge;",
		"CreateInvoiceStatusLog": "INSERT INTO {DB}.invoice_status_log (id, invoice_id, class, from_invoice_vertex_id, to_invoice_vertex_id, created_at, description, author) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		"CreateInvoiceStatus":    "INSERT INTO {DB}.invoice_status (invoice_id, invoice_vertex_class, invoice_vertex_id) VALUES (?, ?, ?);",
		"UpdateInvoiceStatus":    "UPDATE {DB}.invoice_status SET invoice_vertex_id = ? WHERE invoice_id = ? and invoice_vertex_class = ?;",
	}
)
