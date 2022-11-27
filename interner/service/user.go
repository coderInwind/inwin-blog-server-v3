package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
	"inwind-blog-server-v3/utils"
	"inwind-blog-server-v3/utils/dto"
)

type UserService struct{}

func (UserService) GetUserList(params request.PageRequest) (users []model.User, total int64, err error) {

	pageIndex, pageSize := params.PageIndex, params.PageSize

	err = global.DB.Scopes(dto.Paginate(pageIndex, pageSize)).Preload("Role").Find(&users).Limit(-1).Offset(-1).Count(&total).Error

	return users, total, err
}

func (UserService) Login(params request.LoginRequest) (string, error) {
	username, password := params.Username, params.Password

	user := &model.User{}
	err := global.DB.Where("username = ?", username).First(&user).Error

	if err != nil {
		return "", err
	}

	// 将密码和数据库加密了的密码进行比对
	if ok := utils.BcryptCheck(password, user.Password); !ok {
		return "", errors.New("the user name or password is incorrect")
	}

	// 生成token
	jwt := utils.NewJWT()
	Claims := jwt.CreateClaims(user.Id)
	token, err := jwt.GenerateToken(Claims)
	if err != nil {
		return "", errors.New("token generation error")
	}

	// 将jwt存入redis
	if err := utils.SetRedisJwt(token, username); err != nil {
		fmt.Println(err)
		return "", errors.New(" 设置登录状态失败! ")
	}

	return token, nil
}

func (UserService) Logout() {

	// 删除redis中token
	fmt.Println("退出登录了")

}

func (UserService) CreateUser(u model.User) error {
	user := model.User{}
	// 先查询是否有此用户名
	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("this user name already exists")
	}

	//加密
	u.Password = utils.BcryptHash(u.Password)

	// 创建失败
	if err := global.DB.Create(&u).Error; err != nil {
		return err
	}

	return nil
}

func (UserService) GetUserById(id uint) (model.User, error) {
	var user model.User
	err := global.DB.Where("id = ?", id).Preload("Role").First(&user).Error

	return user, err
}
