package service

type CreateInputOutputDto struct {
	Name string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Document string `json:"document" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type Create struct {
	

}