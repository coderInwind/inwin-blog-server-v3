package response

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/errcode"
	"inwind-blog-server-v3/common/serializer"
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

type ErrorResponse struct {
	Code   int
	Msg    string
	Detail string
}

func NewResponse(c *gin.Context) *Response {
	return &Response{ctx: c}
}

func (r *Response) OkWithList(data []serializer.Blog) {
	r.ctx.JSON(http.StatusOK, CommonResponse{
		Code: 0,
		Data: data,
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
