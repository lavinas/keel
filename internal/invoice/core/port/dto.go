package port

type InsertInputDto interface {
	Validate() error
}

type InsertOutputDto interface {
	Load(status string, reference string)
}
