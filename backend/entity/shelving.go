package entity

import (
	"gorm.io/gorm"
)

type Label struct {
	gorm.Model
	Name     string
	Stock    []Stock    `gorm:"foreignKey:label_ID"`
	Shelving []Shelving `gorm:"foreignKey:label_ID"`
}

type Shelving struct {
	gorm.Model

	Employee_ID *uint
	Employee    Employee

	Label_ID *uint
	Label    Label

	Stock_ID *uint
	Stock    Stock

	Quantity int
}
