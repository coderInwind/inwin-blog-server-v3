package model

type Blog struct {
	BasicModel
	Title    string `form:"title" json:"title"`
	Content  string `form:"content" json:"content"`
	Src      string `form:"src" json:"src"`
	TagId    int    `form:"tagId" json:"tagId"`
	Overview string `form:"overview" json:"overview"`
	Pv       int    `form:"pv" json:"pv"`
	Like     int    `form:"like" json:"like"`
}
