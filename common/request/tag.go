package request

type EditTag struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

type CreateTag struct {
	Name string `json:"name" form:"name"`
}

type DeleteTag struct {
	Id uint `json:"id" form:"id"`
}
