package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model

	ID          string    `gorm:"primary_key" json:"id"`
	StudentID   int       `gorm:"column:student_id" json:"student_id"`
	Fname       string    `gorm:"column:fname" json:"fname"`
	Lname       string    `gorm:"column:lname" json:"lname"`
	CreatedDate time.Time `gorm:"column:created_date" json:"created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date" json:"updated_date"`
}
