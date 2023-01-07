package entity

import (
	"time"

	"gorm.io/gorm"
)

type Kind struct {
	gorm.Model
	Name      string
	Inventory []Inventory `gorm:"foreignKey:Kind_ID"`
	Stock     []Stock     `gorm:"foreignKey:Kind_ID"`
}

type Storage struct {
	gorm.Model
	Name      string
	Inventory []Inventory `gorm:"foreignKey:Storage_ID"`
}

type Inventory struct {
	gorm.Model
	Name        string
	Quantity    int
	Price       int
	Employee_ID *uint
	Employee    Employee
	Kind_ID     *uint
	Kind        Kind
	Storage_ID  *uint
	Storage     Storage
}

type Stock struct {
	gorm.Model

	Employee_ID *uint
	Employee    Employee

	Kind_ID *uint
	Kind    Kind

	Inventory_ID *uint
	Inventory    Inventory

	Storage_ID *uint
	Storage    Storage

	DateTime time.Time
}

