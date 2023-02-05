package entity

import (
	"gorm.io/gorm"
)

type Label struct {
	gorm.Model
	Name     string
	Shelving []Shelving `gorm:"foreignKey:Label_ID"`
}

type Shelving struct {
	gorm.Model

	Employee_ID *uint
	Employee    Employee

	Label_ID *uint
	Label    Label

	Stock_ID *uint
	Stock    Stock

	Amount int `valid:"required~A number of goods must be in the range 1-20,range(1|20)~A number of goods must be in the range 1-20"`

	Separation []Separation `gorm:"foreignKey:Shelving_ID"`
	Order      []Order      `gorm:"foreignKey:Shelving_ID"`
}
