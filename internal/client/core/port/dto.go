package port

type ClientCreateInputDto interface {
	Validate() error
	Format() error
	GetName() string
	GetNickname() string
	GetDocument() string
	GetPhone() string
	GetEmail() string
}

type ClientCreateOutputDto interface {
	Fill(id, name, nick, doc, phone, email string)
}
