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

	Amount int

	Separation   []Separation `gorm:"foreignKey:Shelving_ID"`
	Order   []Order `gorm:"foreignKey:Shelving_ID"`
}
