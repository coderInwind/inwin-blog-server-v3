package model

type Blog struct {
	BasicModel
	Title    string `json:"title"`
	Content  string `json:"content"`
	Src      string `json:"src"`
	TagId    int    `json:"tagId"`
	Tag      Tag    `gorm:"ForeignKey:TagId"`
	Overview string `json:"overview"`
	Pv       int    `json:"pv"`
	Like     int    `json:"like"`
}
