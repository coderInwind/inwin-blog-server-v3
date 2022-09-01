package model

type Role struct {
	BasicModel
	Role  string `json:"role"`
	Power int    `json:"power"`
}
