package port

type InsertInputDto interface {
	IsBlank() bool
	Validate() error
	Format() error
	Get() (string, string, string, string, string)
}

type UpdateInputDto interface {
	IsBlank() bool
	Validate() error
	Format() error
	Get() (string, string, string, string, string)
}

type InsertOutputDto interface {
	Fill(id, name, nick, doc, phone, email string)
}

type UpdateOutputDto interface {
	Fill(id, name, nick, doc, phone, email string)
}

type FindInputDto interface {
	Validate() error
	Get() (string, string, string, string, string, string)
}

type FindOutputDto interface {
	SetPage(page, perPage uint64)
	Append(id, name, nick, doc, phone, email string)
	Count() int
}
