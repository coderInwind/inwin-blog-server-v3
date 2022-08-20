package model

import "time"

type BasicModel struct {
	ID        uint `gorm:"PrimaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
