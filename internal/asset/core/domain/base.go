package domain

func GetDomain() []interface{} {
	return []interface{}{
		&Tax{},
		&Class{},
		&Asset{},
		&Balance{},
		&Statement{},
		&Portfolio{},
	}
}
