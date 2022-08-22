package service

import (
	"inwind-blog-server-v3/dao"
	"inwind-blog-server-v3/model"
	"inwind-blog-server-v3/serializer"
	"time"
)

type UserListService struct {
	PageSize  int `form:"pageSize"  json:"pageSize" binding:"required"`
	PageIndex int `form:"pageIndex" json:"pageIndex" binding:"required"`
}

type CreateUserService struct {
	Username  string    `form:"username" json:"userName" binding:"required"`
	Password  string    `form:"password" json:"passWord"  binding:"required"`
	RoleId    int       `form:"roleId" json:"role_id"  binding:"required"`
	CreatedAt time.Time `form:"createdAt" json:"created_at"`
	UpdatedAt time.Time `form:"updatedAt" json:"updated_at"`
}

func (c *CreateUserService) CreateUser() serializer.PromptResponse {
	user := model.User{
		Username: c.Username,
		RoleId:   c.RoleId,
		BasicModel: model.BasicModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := user.EncryptPassword(c.Password); err != nil {
		return serializer.BuildErrorResponse(406, "密码加密失败")
	}

	result := dao.CreateUser(user)
	if result.Error != nil {
		return serializer.BuildErrorResponse(400, "新增用户失败")
	}

	return serializer.BuildSuccessResponse()

}

func (u *UserListService) GetUserList() serializer.Response {
	user, total := dao.SelectUserList(u.PageSize, u.PageIndex)

	return serializer.BuildListResponse(serializer.BuildUsers(user), total)
}
