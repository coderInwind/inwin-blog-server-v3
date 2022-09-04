package request

type PageInfo struct {
	PageSize  int `form:"pageSize"  json:"pageSize" binding:"required"`
	PageIndex int `form:"pageIndex" json:"pageIndex" binding:"required"`
}
