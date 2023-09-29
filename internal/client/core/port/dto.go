package port

type CreateInputDto interface {
	Validate() error
	Format() error
	GetName() string
	GetNickname() string
	GetDocument() string
	GetPhone() string
	GetEmail() string
}

type CreateOutputDto interface {
	Fill(id, name, nick, doc, phone, email string)
}
