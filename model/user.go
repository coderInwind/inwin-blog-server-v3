package model

import "time"

type Role struct {
	Id    int `gorm:"primary_key"`
	Role  string
	Power int
}

type User struct {
	Id        int `gorm:"primary_key"`
	Username  string
	Password  string
	RoleId    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Role      Role
}
