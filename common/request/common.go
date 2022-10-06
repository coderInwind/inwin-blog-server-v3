package request

type PageRequest struct {
	PageSize  int `form:"pageSize"  json:"pageSize" binding:"required"`
	PageIndex int `form:"pageIndex" json:"pageIndex" binding:"required"`
}

type AllowEmptyPageRequest struct {
	PageSize  int `form:"pageSize"  json:"pageSize"`
	PageIndex int `form:"pageIndex" json:"pageIndex"`
}

type SelectBlogRequest struct {
	Id uint `form:"id"  json:"id" binding:"required"`
}
