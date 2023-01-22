package entity

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Quantity    int
	Total       float64
	Employee_ID *uint
	Employee    Employee
	Member_ID   *uint
	Member      Member
	Shelving_ID			*uint
	Shelving			Shelving
	Payment     []Payment 	`gorm:"foreignKey:Shopping_Cart_ID"`
}
