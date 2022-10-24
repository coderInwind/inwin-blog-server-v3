package request

type EditBlog struct {
	Id       uint   `form:"id" json:"id" binding:"required"`
	Title    string `form:"title" json:"title" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"`
	Src      string `form:"src" json:"src" binding:"required"`
	Tags     []uint `form:"tags" json:"tags" binding:"required"`
	Overview string `form:"overview" json:"overview" binding:"required"`
	Hidden   int    `from:"hidden" json:"hidden"`
}

type CreateBlog struct {
	Id       uint   `form:"id" json:"id"`
	Title    string `form:"title" json:"title" binding:"required"`
	Content  string `form:"content" json:"content"`
	Src      string `form:"src" json:"src" binding:"required"`
	Tags     []uint `form:"tags" json:"tags" binding:"required"`
	Overview string `form:"overview" json:"overview" binding:"required"`
}

//type EditBlog struct {
//	Id       int64  `form:"id" json:"id" binding:"required"`
//	Title    string `form:"title" json:"title" binding:"required"`
//	Content  string `form:"content" json:"content" binding:"required"`
//	Src      string `form:"src" json:"src" binding:"required"`
//	TagId    int    `form:"tagId" json:"tagId" binding:"required"`
//	Overview string `form:"overview" json:"overview" binding:"required"`
//}
