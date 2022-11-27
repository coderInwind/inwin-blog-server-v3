package model

type User struct {
	BasicModel
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Role     Role   `json:"role"`
	RoleId   int    `form:"roleId" json:"roleId"`
}
