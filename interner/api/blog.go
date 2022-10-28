package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/errcode"
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/common/response"
	"inwind-blog-server-v3/common/serializer"
	"inwind-blog-server-v3/interner/service"
	"inwind-blog-server-v3/utils"
)

type BlogApi struct{}

func (b *BlogApi) GetBlogList(c *gin.Context) {
	var params request.PageRequest
	res := response.NewResponse(c)
	//校验入参
	if err := c.ShouldBind(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	//判断当前用户的权限等级
	power, exists := c.Get("power")
	if !exists {
		res.FailWithMsg(errcode.ServerError)
	}
	// 超管和次管
	if power.(int) >= 9 {
		list, total, err := service.ServiceGroupApp.GetBlogList(params)
		//查询
		if err != nil {
			res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
			return
		}

		res.OkWithList(list, total)
	} else {
		list, total, err := service.ServiceGroupApp.GetOwnBlogList(params)
		//查询
		if err != nil {
			res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
			return
		}

		res.OkWithList(list, total)

	}

}

func (b *BlogApi) GetBlogDetail(c *gin.Context) {
	var params request.SelectBlogRequest
	res := response.NewResponse(c)

	if err := c.ShouldBind(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	detail, err := service.ServiceGroupApp.GetBlogDetail(params)
	if err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}

	//res.OkWithData(serializer.BuildBlog(detail))
	res.OkWithData(detail)
}

func (b *BlogApi) EditBlog(c *gin.Context) {
	var params request.EditBlog
	res := response.NewResponse(c)

	if err := c.ShouldBind(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	if err := service.ServiceGroupApp.EditBlog(params); err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}

	res.Ok()
}

func (b *BlogApi) CreateBlog(c *gin.Context) {
	var params request.CreateBlog
	res := response.NewResponse(c)
	if err := c.ShouldBind(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	//取出当前用户的信息，将blog.author赋值为username
	claims, exists := c.Get("claims")
	if !exists {
		res.FailWithMsg(errcode.ServerError)
	}
	id := claims.(*utils.Claims).Id

	// 序列化参数
	clearedParams := serializer.BuildCreateBlogParams(params, id)

	id, err := service.ServiceGroupApp.CreateBlog(clearedParams)
	if err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}

	res.OkWithData(id)
}

func (b *BlogApi) DeleteBlog(c *gin.Context) {
	var params request.SelectBlogRequest
	res := response.NewResponse(c)
	if err := c.ShouldBind(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	if err := service.ServiceGroupApp.DeleteBlog(params.Id); err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}

	res.Ok()
}
