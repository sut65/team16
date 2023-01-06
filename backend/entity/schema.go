package entity

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model

	FirstName string

	LastName string

	Email string `gorm:"uniqueIndex" valid:"email"`

	Password string

	Position string

	Stock []Stock `gorm:"foreignKey:Employee_ID"`

	Inventory []Inventory `gorm:"foreignKey:Employee_ID"`
}

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
