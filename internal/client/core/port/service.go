package port

type Service interface {
	Create(input CreateInputDto, output CreateOutputDto) error
}
