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

type Type_Comment struct {
	gorm.Model
	Type_Com_Name        string
	// Comment      []Comment `gorm:"foreignKey:Type_Com_ID"`
}