package service

import (
	"inwind-blog-server-v3/interner/dao"
	"inwind-blog-server-v3/interner/model"
	serializer2 "inwind-blog-server-v3/interner/serializer"
	"inwind-blog-server-v3/utils"
)

// CountTagRequest 定义了 Request 结构体作为接口入参的基准，而本项目由于并不会太复杂，所以直接放在了 service 层中便于使用
// 若后续业务不断增长，程序越来越复杂，service 也冗杂了，可以考虑将抽离一层接口校验层，便于解耦逻辑。

type BlogListSerivce struct {
	PageSize  int `form:"pageSize"  json:"pageSize" binding:"required"`
	PageIndex int `form:"pageIndex" json:"pageIndex" binding:"required"`
}

func (b *BlogListSerivce) GetBlogList() []model.Blog {

	//blogs, total, err := dao.SelectBlogList(b.PageSize, b.PageIndex)

	//if err != nil {
	//	return serializer2.BuildListErrorResponse(utils.ERROR, utils.GetMsg(utils.ERROR))
	//}

	//return serializer2.BuildListResponse(serializer2.BuildBlogs(blogs), total)

	return dao.SelectBlogList(b.PageSize, b.PageIndex)
}

type BlogDetailService struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

func (b *BlogDetailService) GetBlogDetail() serializer2.Response {
	blog, err := dao.SelectBlogDetail(b.Id)
	if err != nil {
		serializer2.BuildErrorResponse(utils.ERROR, utils.GetMsg(utils.ERROR))
	}
	return serializer2.Response{
		Code: utils.SUCCESS,
		Data: blog,
		Msg:  utils.GetMsg(utils.SUCCESS),
	}
}

type EditBlogService struct {
	Id       uint   `form:"id" json:"id" binding:"required"`
	Title    string `form:"title" json:"title" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"`
	Src      string `form:"src" json:"src" binding:"required"`
	TagId    string `form:"tagId" json:"tagId" binding:"required"`
	Overview string `form:"overview" json:"overview" binding:"required"`
}

func (e *EditBlogService) EditBlog() {
	//要不试试map？？
	//dao.UpdateBlog[*EditBlogService](e, e.Id)
	//global.DbEngine.Save(&e)
}
