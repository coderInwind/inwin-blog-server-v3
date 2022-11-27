package request

type GetRecordParams struct {
	Date []string ` form:"date" json:"date" binding:"required"`
}
