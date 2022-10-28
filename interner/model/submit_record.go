package model

import "time"

type SubmitRecord struct {
	Date        time.Time
	IsSubmitted int
}
