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

	Number int `valid:"required~Number must be in the range 1-20,range(1|20)~Number must be in the range 1-20"`

	Cost float64 `valid:"required~Cost must be in the range 1-1000,range(1|1000)~Cost must be in the range 1-1000"`

	Separation []Separation `gorm:"foreignKey:Shelving_ID"`
	Order      []Order      `gorm:"foreignKey:Shelving_ID"`
}
