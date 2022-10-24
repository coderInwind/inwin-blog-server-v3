package model

type User struct {
	BasicModel
	Username string `form:"username"`
	Password string `form:"password"`
	RoleId   int    `form:"roleId"`
	Role     Role   `gorm:"ForeignKey:RoleId"`
	Author   string `form:"author"`
}
