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
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password"  binding:"required"`
	RoleId    int       `json:"roleId"  binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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
