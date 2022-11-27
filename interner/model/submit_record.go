package model

import "time"

type SubmitRecord struct {
	Date  time.Time `json:"date"`
	Count int       `json:"count"`
}
