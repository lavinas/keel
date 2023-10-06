package repo

const (
	clientSaveQuery              = `INSERT INTO client (id, name, nickname, document, phone, email) VALUES (?, ?, ?, ?, ?, ?)`
	clientDocumentDuplicityQuery = `SELECT COUNT(*) count FROM client WHERE document = ?`
	clientEmailDuplicityQuery    = `SELECT COUNT(*) count FROM client WHERE email = ?`
	clientTruncateQuery          = `TRUNCATE TABLE client`
	clientListBase               = `SELECT id, name, nickname, document, phone, email FROM client where 1 = 1`
	clientListFilterName         = ` AND name LIKE ?`
	clientListFilterNick         = ` AND nickname LIKE ?`
	clientListFilterDoc          = ` AND document LIKE ?`
	clientListFilterEmail        = ` AND email LIKE ?`
	clientListPagination         = ` LIMIT ? OFFSET ?`
	clientLoadById               = `SELECT id, name, nickname, document, phone, email FROM client WHERE id = ?`
	clientUpdateQuery            = `UPDATE client SET name = ?, nickname = ?, document = ?, phone = ?, email = ? WHERE id = ?`
)
