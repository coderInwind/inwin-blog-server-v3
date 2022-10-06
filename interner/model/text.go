package model

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	student []Student
}

type Student struct {
	gorm.Model
	name string
}
