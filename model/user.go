package model

type Role struct {
	BasicModel
	Role  string `json:"role"`
	Power int    `json:"power"`
}

type User struct {
	BasicModel
	Username string
	Password string
	RoleId   int
	Role     Role `gorm:"ForeignKey:RoleId"`
}
