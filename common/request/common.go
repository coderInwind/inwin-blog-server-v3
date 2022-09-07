package request

type PageRequest struct {
	PageSize  int `form:"pageSize"  json:"pageSize" binding:"required"`
	PageIndex int `form:"pageIndex" json:"pageIndex" binding:"required"`
}

type DetailRequest struct {
	Id int `form:"id"  json:"id" binding:"required"`
}
