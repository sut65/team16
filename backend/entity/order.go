package entity

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Quantity    int
	Prices		float64
	Shelving_ID			*uint
	Shelving			Shelving
	Shopping_Cart_ID		*uint
	Shopping_Cart			Shopping_Cart
}

type Status struct {
	gorm.Model
	Status    	string

	Shopping_Cart     []Shopping_Cart 	`gorm:"foreignKey:Status_ID"`
}

type Shopping_Cart struct {
	gorm.Model
	Total       float64

	Employee_ID *uint
	Employee    Employee
	Member_ID   *uint
	Member      Member
	Status_ID   *uint
	Status      Status
	Order		[]Order 	`gorm:"foreignKey:Shopping_Cart_ID"`
	Payment     []Payment 	`gorm:"foreignKey:Shopping_Cart_ID"`
}
