package entity

import (
	// "time"

	"gorm.io/gorm"
)


type Review_Point struct {
	gorm.Model
	Point        int
	Point_Name   string
	// Comment      []Comment `gorm:"foreignKey:Review_point_ID"`
}