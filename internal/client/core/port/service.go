package port

type Service interface {
	ClientCreate(input CreateInputDto, output CreateOutputDto) error
}
