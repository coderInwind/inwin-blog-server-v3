package model

type User struct {
	BasicModel
	Username string
	Password string
	RoleId   int
	Role     Role `gorm:"ForeignKey:RoleId"`
}
