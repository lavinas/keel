package port

type ClientCreateInputDto interface {
	Validate() error
	Format() error
	Get() (string, string, string, string, string)
}

type ClientCreateOutputDto interface {
	Fill(id, name, nick, doc, phone, email string)
}

type ClientListInputDto interface {
	Validate() error
	Get() (string, string, string, string, string, string, string)
}

type ClientListOutputDto interface {
	SetPage(page, perPage uint64)
	Append(id, name, nick, doc, phone, email string)
	Count() int
}
