package model

type Blog struct {
	BasicModel
	Title    string `form:"title" json:"title"`
	Content  string `form:"content" json:"content"`
	Src      string `form:"src" json:"src"`
	Tags     []Tag  `gorm:"many2many:blog_tags" json:"tags"`
	Overview string `form:"overview" json:"overview"`
	Pv       int    `form:"pv" json:"pv"`
	Like     int    `form:"like" json:"like"`
	Hidden   int    `json:"hidden"`
	AuthorId uint   `form:"authorId" json:"authorId"`
	Author   User   `json:"author" form:"author"`
}
