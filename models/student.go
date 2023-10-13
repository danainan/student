package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID    int    `gorm:"uniqueIndex;not null" json:"id"`
	Fname string `gorm:"not null" json:"fname"`
	Lname string `gorm:"not null" json:"lname"`
}
