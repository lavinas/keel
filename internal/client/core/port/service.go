package port

type Service interface {
	Insert(input InsertInputDto, output InsertOutputDto) error
	Find(input FindInputDto, output FindOutputDto) error
	Update(id string, input InsertInputDto, output InsertOutputDto) error
	Get(param string, input InsertInputDto, output InsertOutputDto) error
}
