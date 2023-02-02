package entity

import (

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Quantity         int		`valid:"required~ระบุจำนวน,range(1|100)~จำนวนอยู่ในช่วง (1-100)"` 
	Prices           float64	`valid:"required~ระบุราคา,range(1|1000000)~ราคาอยู่ในช่วง (1-1000000)"` 
	Shelving_ID      *uint
	Shelving         Shelving	
	Shopping_Cart_ID *uint		
	Shopping_Cart    Shopping_Cart
}

type Status struct {
	gorm.Model
	Status string

	Shopping_Cart []Shopping_Cart `gorm:"foreignKey:Status_ID"`
}
