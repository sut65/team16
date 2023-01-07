package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name string
	Email string `gorm:"uniqueIndex" valid:"email"`
	Password string
	Position string

	Stock []Stock `gorm:"foreignKey:Employee_ID"`
	Inventory []Inventory `gorm:"foreignKey:Employee_ID"`
	Member []Member `gorm:"foreignKey:Employee_ID"`
}
