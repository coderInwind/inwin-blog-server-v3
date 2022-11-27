package model

type Role struct {
	BasicModel
	Name  string `json:"name"`
	Power int    `json:"power"`
}
