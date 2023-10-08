package dto

// ClientInserOutputDto is the output DTO used to create a client
type ClientInserOutputDto struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Document string `json:"document" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func NewClientInserOutputDto() *ClientInserOutputDto {
	return &ClientInserOutputDto{
		Id:       "",
		Name:     "",
		Nickname: "",
		Document: "",
		Phone:    "",
		Email:    "",
	}
}

// ClientInsertInputDto is the input DTO used to create a client
func (o *ClientInserOutputDto) Fill(id, name, nick, doc, phone, email string) {
	o.Id = id
	o.Name = name
	o.Nickname = nick
	o.Document = doc
	o.Phone = phone
	o.Email = email
}
