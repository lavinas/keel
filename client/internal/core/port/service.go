package port

type Service interface {
	Insert(input InsertInputDto, output InsertOutputDto) error
	Find(input FindInputDto, output FindOutputDto) error
	Update(id string, input UpdateInputDto, output UpdateOutputDto) error
	Get(param string, output InsertOutputDto) error
}
