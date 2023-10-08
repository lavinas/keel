package port

type ClientInsertInputDto interface {
	IsBlank() bool
	Validate() error
	ValidateUpdate() error
	Format() error
	FormatUpdate() error
	Get() (string, string, string, string, string)
}

type ClientInserOutputDto interface {
	Fill(id, name, nick, doc, phone, email string)
}

type ClientListInputDto interface {
	Validate() error
	Get() (string, string, string, string, string, string)
}

type ClientListOutputDto interface {
	SetPage(page, perPage uint64)
	Append(id, name, nick, doc, phone, email string)
	Count() int
}
