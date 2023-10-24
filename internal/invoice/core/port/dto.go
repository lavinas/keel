package port

type InsertInputDto interface {
	Validate() error
}

type InsertOutputDto interface {
}
