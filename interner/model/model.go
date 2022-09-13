package model

import "time"

type BasicModel struct {
	ID        uint      `form:"id" gorm:"PrimaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
