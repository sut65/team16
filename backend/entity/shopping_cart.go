package entity

import (
	"gorm.io/gorm"
)

type Shopping_Cart struct {
	gorm.Model
	Total       float64
	Employee_ID *uint
	Employee    Employee
	Member_ID   *uint
	Member      Member
	Status_ID		*uint
	Status      Status
	Oder        []Order   `gorm:"foreignKey:Shopping_Cart_ID"`
	Payment     []Payment `gorm:"foreignKey:Shopping_Cart_ID"`
}
