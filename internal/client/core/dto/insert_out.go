// Package dto is the package that defines the DTOs for the client adapter
package dto

// InsertOutputDto is the output DTO used to create a client
type InsertOutputDto struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Document string `json:"document" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func NewInsertOutputDto() *InsertOutputDto {
	return &InsertOutputDto{
		Id:       "",
		Name:     "",
		Nickname: "",
		Document: "",
		Phone:    "",
		Email:    "",
	}
}

// InsertInputDto is the input DTO used to create a client
func (o *InsertOutputDto) Fill(id, name, nick, doc, phone, email string) {
	o.Id = id
	o.Name = name
	o.Nickname = nick
	o.Document = doc
	o.Phone = phone
	o.Email = email
}

// Get returns the values of the output DTO
func (o *InsertOutputDto) Get() (string, string, string, string, string, string) {
	return o.Id, o.Name, o.Nickname, o.Document, o.Phone, o.Email
}