package port

type Service interface {
	ClientInsert(input ClientInsertInputDto, output ClientInserOutputDto) error
	ClientList(input ClientListInputDto, output ClientListOutputDto) error
	ClientUpdate(id string, input ClientInsertInputDto, output ClientInserOutputDto) error
	ClientGet(param string, input ClientInsertInputDto, output ClientInserOutputDto) error
}
