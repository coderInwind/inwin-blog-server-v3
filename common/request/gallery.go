package request

type DeletePicture struct {
	Id int `json:"id" form:"id" binding:"required"`
}
