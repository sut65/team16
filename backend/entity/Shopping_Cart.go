package entity

import (
	"gorm.io/gorm"
)

type Shopping_Cart struct {
	gorm.Model
	Total       float64
	Employee_ID *uint
	Employee    Employee
	Mem_Tel     string
	Member      Member
	Oder        []Order   `gorm:"foreignKey:Shopping_Cart_ID"`
	Payment     []Payment `gorm:"foreignKey:Shopping_Cart_ID"`
}

type Order struct {
gorm.Model
	Quantity    		int
	
	Shelving_ID			*uint
	Shelving			Shelving
	Shopping_Cart_ID	*uint
	Shopping_Cart		Shopping_Cart
}
