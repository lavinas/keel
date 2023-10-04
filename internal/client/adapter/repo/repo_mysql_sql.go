package repo

const (
	clientSaveQuery              = `INSERT INTO client (id, name, nickname, document, phone, email) VALUES (?, ?, ?, ?, ?, ?)`
	clientDocumentDuplicityQuery = `SELECT COUNT(*) count FROM client WHERE document = ?`
	clientEmailDuplicityQuery    = `SELECT COUNT(*) count FROM client WHERE email = ?`
	clientTruncateQuery          = `TRUNCATE TABLE client`
	clientGetAll                 = `SELECT id, name, nickname, document, phone, email FROM client`
)
