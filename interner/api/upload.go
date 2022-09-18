package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/response"
	"inwind-blog-server-v3/interner/service"
)

type UploadApi struct {
}

func (UploadApi) GetUploadUrl(c *gin.Context) {
	res := response.NewResponse(c)
	token := service.ServiceGroupApp.GetUploadService()

	res.OkWithData(token)
}
