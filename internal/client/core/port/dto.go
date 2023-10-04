package port

type ClientCreateInputDto interface {
	Validate() error
	Format() error
	Get() (string, string, string, string, string)
}

type ClientCreateOutputDto interface {
	Fill(id, name, nick, doc, phone, email string)
}

type ClientListOutputDto interface {
	Append(id, name, nick, doc, phone, email string)
}
