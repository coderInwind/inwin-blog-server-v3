package service

import (
	"inwind-blog-server-v3/interner/dao"
	model2 "inwind-blog-server-v3/interner/model"
	serializer2 "inwind-blog-server-v3/interner/serializer"
	"time"
)

type UserListService struct {
	PageSize  int `form:"pageSize"  json:"pageSize" binding:"required"`
	PageIndex int `form:"pageIndex" json:"pageIndex" binding:"required"`
}

type CreateUserService struct {
	Username  string    `form:"username" json:"username" binding:"required"`
	Password  string    `form:"password" json:"password"  binding:"required"`
	RoleId    int       `form:"roleId" json:"roleId"  binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c *CreateUserService) CreateUser() serializer2.PromptResponse {
	user := model2.User{
		Username: c.Username,
		RoleId:   c.RoleId,
		BasicModel: model2.BasicModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := user.EncryptPassword(c.Password); err != nil {
		return serializer2.BuildErrorResponse(406, "密码加密失败")
	}

	result := dao.CreateUser(user)
	if result.Error != nil {
		return serializer2.BuildErrorResponse(400, "新增用户失败")
	}

	return serializer2.BuildSuccessResponse()

}

func (u *UserListService) GetUserList() serializer2.Response {
	user, total := dao.SelectUserList(u.PageSize, u.PageIndex)

	return serializer2.BuildListResponse(serializer2.BuildUsers(user), total)
}
