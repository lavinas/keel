package port

type ApiConsumer interface {
	GetClientByNickname(nickname string) (GetClientByNicknameInputDto, error)
	GetId() string
	GetName() string
	GetNickname() string
	GetDocument() uint64
	GetPhone() uint64
	GetEmail() string
}