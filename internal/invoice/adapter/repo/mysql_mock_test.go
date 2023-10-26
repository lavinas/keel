package repo

type InvoiceClientMock struct {
}

func (i *InvoiceClientMock) Load(nickname, clientId, name, email string, phone, document uint64) {
}
func (i *InvoiceClientMock) Save() error {
	return nil
}
func (i *InvoiceClientMock) GetId() string {
	return "1"
}
func (i *InvoiceClientMock) GetNickname() string {
	return "nickname"
}
func (i *InvoiceClientMock) GetClientId() string {
	return "client_id"
}
func (i *InvoiceClientMock) GetName() string {
	return "name"
}
func (i *InvoiceClientMock) GetDocument() uint64 {
	return 1
}
func (i *InvoiceClientMock) GetPhone() uint64 {
	return 1
}
func (i *InvoiceClientMock) GetEmail() string {
	return "email"
}
