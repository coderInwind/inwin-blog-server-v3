package model

type Blog struct {
	BasicModel
	Title    string `json:"title"`
	Content  string `json:"content"`
	Src      string `json:"src"`
	Tag      string `json:"tag"`
	Overview string `json:"overview"`
	Pv       int    `json:"pv"`
	Like     int    `json:"like"`
}
