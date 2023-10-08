package repo

const (
	clientSaveQuery              = `INSERT INTO client (id, name, nickname, document, phone, email) VALUES (?, ?, ?, ?, ?, ?)`
	clientDocumentDuplicityQuery = `SELECT COUNT(1) count FROM client WHERE document = ? and id != ?`
	clientEmailDuplicityQuery    = `SELECT COUNT(1) count FROM client WHERE email = ? and id != ?`
	clientNickDuplicityQuery     = `SELECT COUNT(1) count FROM client WHERE nickname = ? and id != ?`
	clientTruncateQuery          = `TRUNCATE TABLE client`
	clientSetBase                = `SELECT id, name, nickname, document, phone, email FROM client where 1 = 1`
	clientSetFilterName          = ` AND name LIKE ?`
	clientSetFilterNick          = ` AND nickname LIKE ?`
	clientSetFilterDoc           = ` AND document LIKE ?`
	clientSetFilterEmail         = ` AND email LIKE ?`
	clientSetPagination          = ` LIMIT ? OFFSET ?`
	clientUpdateQuery            = `UPDATE client SET name = ?, nickname = ?, document = ?, phone = ?, email = ? WHERE id = ?`
	clientGetById                = `SELECT id, name, nickname, document, phone, email FROM client WHERE id = ?`
	clientGetByNick              = `SELECT id, name, nickname, document, phone, email FROM client WHERE nickname = ?`
	clientGetByEmail             = `SELECT id, name, nickname, document, phone, email FROM client WHERE email = ?`
	clientGetByDoc               = `SELECT id, name, nickname, document, phone, email FROM client WHERE document = ?`
	clientGetByPhone             = `SELECT id, name, nickname, document, phone, email FROM client WHERE phone = ?`
)
