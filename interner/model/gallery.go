package model

type Gallery struct {
	Id   int64  `json:"id" form:"id"`
	Url  string `json:"url" form:"url"`
	Name string `json:"name" form:"name"`
}
