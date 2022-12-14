package entity

import (
	"time"

	"gorm.io/gorm"
)


type Review_Point struct {
	gorm.Model
	Point        int
	Point_Name   string
	Comment      []Comment `gorm:"foreignKey:Review_point_ID"`
}

type Type_Comment struct {
	gorm.Model
	Type_Com_Name        string
	Comment      []Comment `gorm:"foreignKey:Type_Com_ID"`
}

type Comment struct {
	gorm.Model
	Comments    string

	Review_point_ID  *uint
	Review_point      Review_Point

	Payment_ID *uint
	Payment     Payment

	Type_Com_ID *uint
	Type_Com    Type_Comment

	Date_Now    time.Time
	Bought_now  int
}
