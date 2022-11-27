package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/api"
)

func InitSubmitRecord(router *gin.RouterGroup) {
	submitRecordRouter := router.Group("/submitRecord")
	submitRecord := api.SubmitRecordApi{}
	{
		submitRecordRouter.POST("/list", submitRecord.GetSubmitRecord)
	}
}
