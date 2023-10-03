package port

type Service interface {
	ClientCreate(input ClientCreateInputDto, output ClientCreateOutputDto) error
}
