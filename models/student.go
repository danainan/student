package models

import (
	"time"
)

type Student struct {
	ID          string    `gorm:"primary_key" json:"id"`
	StudentID   string    `json:"student_id"`
	Fname       string    `json:"fname"`
	Lname       string    `json:"lname"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}
