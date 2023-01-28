package entity

import (
	"gorm.io/gorm"
)

// struct entity
type Employee struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex" valid:"email"`
	Password string
	Position string

	Stock                 []Stock                 `gorm:"foreignKey:Employee_ID"`
	Member                []Member                `gorm:"foreignKey:Employee_ID"`
	Leave                 []Leave                 `gorm:"foreignKey:Employee_ID"`
	Employee_attendance   []Employee_attendance   `gorm:"foreignKey:Employee_ID"`
	Record_employee_leave []Record_employee_leave `gorm:"foreignKey:Employee_ID"`
<<<<<<< HEAD
	Discount			  []Discount 			  `gorm:"foreignKey:Employee_ID"`
	Delivery  			  []Delivery 			  `gorm:"foreignKey:Employee_ID"`
	Payment  			  []Payment 			  `gorm:"foreignKey:Employee_ID"`
	Separation			  []Separation 			  `gorm:"foreignKey:Employee_ID"`
	Shelving			  []Shelving 			  `gorm:"foreignKey:Employee_ID"`
	Shopping_Cart		  []Shopping_Cart		  `gorm:"foreignKey:Employee_ID"`
}
=======
	Discount              []Discount              `gorm:"foreignKey:Employee_ID"`
	Delivery              []Delivery              `gorm:"foreignKey:Employee_ID"`
	Payment               []Payment               `gorm:"foreignKey:Employee_ID"`
	Separation            []Separation            `gorm:"foreignKey:Employee_ID"`
	Shelving              []Shelving              `gorm:"foreignKey:Employee_ID"`
	Shopping_Cart         []Shopping_Cart         `gorm:"foreignKey:Employee_ID"`
}
>>>>>>> issue-126
