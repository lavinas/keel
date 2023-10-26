package mysql

var (
	dbQueries = map[string]string{
		"clientSaveQuery":              `INSERT INTO {DB}.client (id, name, nickname, document, phone, email) VALUES (?, ?, ?, ?, ?, ?)`,
		"clientDocumentDuplicityQuery": `SELECT COUNT(1) count FROM {DB}.client WHERE document = ? and id != ?`,
		"clientEmailDuplicityQuery":    `SELECT COUNT(1) count FROM {DB}.client WHERE email = ? and id != ?`,
		"clientNickDuplicityQuery":     `SELECT COUNT(1) count FROM {DB}.client WHERE nickname = ? and id != ?`,
		"clientTruncateQuery":          `TRUNCATE TABLE {DB}.client`,
		"clientSetBase":                `SELECT id, name, nickname, document, phone, email FROM {DB}.client where 1 = 1`,
		"clientSetFilterName":          ` AND name LIKE ?`,
		"clientSetFilterNick":          ` AND nickname LIKE ?`,
		"clientSetFilterDoc":           ` AND document LIKE ?`,
		"clientSetFilterPhone":         ` AND phone LIKE ?`,
		"clientSetFilterEmail":         ` AND email LIKE ?`,
		"clientSetPagination":          ` LIMIT ? OFFSET ?`,
		"clientUpdateQuery":            `UPDATE {DB}.client SET name = ?, nickname = ?, document = ?, phone = ?, email = ? WHERE id = ?`,
		"clientGetById":                `SELECT id, name, nickname, document, phone, email FROM {DB}.client WHERE id = ?`,
		"clientGetByNick":              `SELECT id, name, nickname, document, phone, email FROM {DB}.client WHERE nickname = ?`,
		"clientGetByEmail":             `SELECT id, name, nickname, document, phone, email FROM {DB}.client WHERE email = ?`,
		"clientGetByDoc":               `SELECT id, name, nickname, document, phone, email FROM {DB}.client WHERE document = ?`,
		"clientGetByPhone":             `SELECT id, name, nickname, document, phone, email FROM {DB}.client WHERE phone = ?`,
	}
)
