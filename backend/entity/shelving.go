package entity

import (
	"gorm.io/gorm"
)

type Label struct {
	gorm.Model
	Name     string
	Product  []Product  `gorm:"foreignKey:Label_ID"`
	Shelving []Shelving `gorm:"foreignKey:Label_ID"`
}

type Product struct {
	gorm.Model
	Name        string
	Price       float64
	Employee_ID *uint
	Employee    Employee
	Label_ID    *uint
	Label       Label
}
type Shelving struct {
	gorm.Model

	Employee_ID *uint
	Employee    Employee

	Label_ID *uint
	Label    Label

	Product_ID *uint
	Product    Product

	Quantity int
}
