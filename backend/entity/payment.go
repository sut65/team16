package entity

import (
	"time"
	"gorm.io/gorm"
)

type Payment_method struct {
	gorm.Model
	Method	string
	Payment  []Payment `gorm:"foreignKey:Payment_method_ID"`
}

type Payment struct {
	gorm.Model
	Time				time.Time
	Price       		float64
	Order_ID			*uint
	Order				Order
	Payment_method_ID	*uint
	Payment_method		Payment_method
	Employee_ID 		*uint
	Employee    		Employee
	Delivery  			[]Delivery 			  `gorm:"foreignKey:Payment_ID"`
}
