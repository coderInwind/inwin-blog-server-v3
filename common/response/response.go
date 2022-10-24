package response

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/errcode"
	"net/http"
)

type Response struct {
	ctx *gin.Context
}

type CommonResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type ListData struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
}

type ErrorResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Detail string `json:"detail"`
}

func NewResponse(c *gin.Context) *Response {
	return &Response{ctx: c}
}

func (r *Response) Ok() {
	r.ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "operate successfully!",
	})
}

func (r *Response) OkWithData(data any) {
	r.ctx.JSON(http.StatusOK, CommonResponse{
		Code: 0,
		Data: data,
		Msg:  "operate successfully!",
	})
}

func (r *Response) OkWithList(data any, total int64) {
	r.ctx.JSON(http.StatusOK, CommonResponse{
		Code: 0,
		Data: ListData{
			List:  data,
			Total: total,
		},
		Msg: "operate successfully!",
	})
}

func (r *Response) OkWithMsg() {
	r.ctx.JSON(http.StatusOK, CommonResponse{
		Code: 0,
		Msg:  "operate successfully!",
	})
}

func (r *Response) FailWithMsg(err *errcode.Error) {
	r.ctx.JSON(err.StatusCode(), ErrorResponse{
		Code:   err.GetCode(),
		Msg:    err.GetMsg(),
		Detail: err.GetDetail(),
	})
}
