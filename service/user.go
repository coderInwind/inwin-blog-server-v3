package service

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common"
	"inwind-blog-server-v3/dao"
	"inwind-blog-server-v3/serializer"
)

type UserListParams struct {
	PageSize  int `form:"pageSize"  json:"pageSize" binding:"required"`
	PageIndex int `form:"pageIndex" json:"pageIndex" binding:"required"`
}

func CreateUser(c *gin.Context, params common.UserParams) {
	dao.CreateUser(c, params)
}

func (u *UserListParams) GetUserListService() serializer.Response {
	user, total := dao.SelectUserList(u.PageSize, u.PageIndex)

	return serializer.BuildListResponse(serializer.BuildUsers(user), total)
}
