package entity

import (
	"gorm.io/gorm"
)

type Label struct {
	gorm.Model
	Name     string
<<<<<<< HEAD
	Product  []Product  `gorm:"foreignKey:Label_ID"`
	Shelving []Shelving `gorm:"foreignKey:Label_ID"`
=======
	Product  []Product  `gorm:"foreignKey:label_ID"`
	Shelving []Shelving `gorm:"foreignKey:label_ID"`
>>>>>>> main
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
<<<<<<< HEAD
=======

>>>>>>> main
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
