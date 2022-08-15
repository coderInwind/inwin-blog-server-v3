package model

import "time"

type Role struct {
	Id    uint `gorm:"primary_key"`
	Role  string
	Power int
}

type User struct {
	Id        uint `gorm:"primary_key"`
	Username  string
	Password  string
	RoleId    uint
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
