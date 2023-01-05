package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model

	FirstName string

	LastName string

	Email string `gorm:"uniqueIndex" valid:"email"`

	Password string

	Position string

	Stock []Stock `gorm:"uniqueIndex" valid:"EmployeeID"`
}

type Kind struct {
	gorm.Model
	Name    string
	Product []Product `gorm:"foreignKey:KindID"`
	Stock   []Stock   `gorm:"foreignKey:KindID"`
}

type Product struct {
	gorm.Model
	Name     string
	Quantity int
	Price    int
	Kind     Kind    `gorm:"references:id"`
	Stock    []Stock `gorm:"references:ProductID"`
}
type Stock struct {
	gorm.Model

	Employee_ID *uint
	Employee    Employee `gorm:"references:id"`

	Kind_ID *uint
	Kind    Kind `gorm:"references:id"`

	Product_ID *uint
	Product    Product `gorm:"references:id"`

	Section string
}
