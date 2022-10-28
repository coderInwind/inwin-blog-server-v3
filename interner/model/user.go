package model

type User struct {
	BasicModel
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	RoleId   int    `form:"roleId" json:"roleId"`
	Role     Role   `json:"role"`
}
