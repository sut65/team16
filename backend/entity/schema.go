package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
<<<<<<< HEAD

	Name string

	Email string `gorm:"uniqueIndex" valid:"email"`

=======
	Name     string
	Email    string `gorm:"uniqueIndex" valid:"email"`
>>>>>>> main
	Password string
	Position string

	Stock     []Stock     `gorm:"foreignKey:Employee_ID"`
	Inventory []Inventory `gorm:"foreignKey:Employee_ID"`
<<<<<<< HEAD

	Product []Product `gorm:"foreignKey:Employee_ID"`

	Shelving []Shelving `gorm:"foreignKey:Employee_ID"`
=======
	Member    []Member    `gorm:"foreignKey:Employee_ID"`
	Leave     []Leave     `gorm:"foreignKey:Employee_ID"`
>>>>>>> main
}
