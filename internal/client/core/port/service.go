package port

type Service interface {
	ClientCreate(input ClientCreateInputDto, output ClientCreateOutputDto) error
	ClientList(input ClientListInputDto, output ClientListOutputDto) error
}
