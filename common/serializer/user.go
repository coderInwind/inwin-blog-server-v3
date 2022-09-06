package serializer

import (
	"inwind-blog-server-v3/interner/model"
	"time"
)

type User struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func BuildUser(item model.User) User {
	return User{
		Id:        item.ID,
		Username:  item.Username,
		Password:  item.Password,
		Role:      item.Role.Role,
		CreatedAt: item.BasicModel.CreatedAt,
		UpdatedAt: item.BasicModel.UpdatedAt,
	}
}

func BuildUsers(list []model.User) (users []User) {
	for _, item := range list {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
