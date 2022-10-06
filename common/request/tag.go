package request

type EditTag struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}
