package entity

import (
	"gorm.io/gorm"
)

type Shopping_Cart struct {
	gorm.Model
	Total       float64
	Employee_ID *uint		`valid:"-"`
	Employee    Employee	`gorm:"references:id" valid:"-"`
	Member_ID   *uint		`valid:"-"`
	Member      Member		`gorm:"references:id" valid:"-"`
	Status_ID   *uint		`valid:"-"`
	Status      Status		`gorm:"references:id" valid:"-"`
	Oder        []Order   	`gorm:"foreignKey:Shopping_Cart_ID"`
	Payment     []Payment 	`gorm:"foreignKey:Shopping_Cart_ID"`
}


