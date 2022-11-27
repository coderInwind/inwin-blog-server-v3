package model

type Menu struct {
	BasicModel
	Path      string  `json:"path"`
	Name      string  `json:"name"`
	Title     string  `json:"title"`
	Icon      string  `json:"icon"`
	Component string  `json:"component"`
	Hidden    int     `json:"hidden"`
	Pid       uint    `json:"pid"`
	Redirect  string  `json:"redirect"`
	Power     int     `json:"power"`
	Child     []*Menu `gorm:"-" json:"child"`
}
