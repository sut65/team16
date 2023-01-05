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

	Invetory []Invetory `gorm:"foreignKey:Employee_ID"`
}

type Kind struct {
	gorm.Model
	Name     string
	Invetory []Invetory `gorm:"foreignKey:Kind_ID"`
	Stock    []Stock    `gorm:"foreignKey:Kind_ID"`
}

type Storage struct {
	gorm.Model
	Name     string
	Invetory []Invetory `gorm:"foreignKey:Storage_ID"`
}

type Invetory struct {
	gorm.Model
	Name     string
	Quantity int
	Price    int
	Employee Employee `gorm:"reference:id"`
	Kind     Kind     `gorm:"references:id"`
	Storage  Storage  `gorm:"reference:id"`
	Stock    []Stock  `gorm:"references:Invetory_ID"`
}
type Stock struct {
	gorm.Model

	Employee_ID *uint
	Employee    Employee `gorm:"reference:id"`

	Kind_ID *uint
	Kind    Kind `gorm:"references:id"`

	Invetory_ID *uint
	Invetory    Invetory `gorm:"references:id"`

	Storage_ID *uint
	Storage    Storage `gorm:"references:id"`

	DateTime time.Time
}
