package entity

import (

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Quantity         int			`valid:"required~ระบุจำนวน,range(1|100)~จำนวนอยู่ในช่วง (1-100),int~จำนวนต้องเป็นจำนวนเต็ม"` 
	Prices           float64		`valid:"required~ระบุราคา,range(1|1000000)~ราคาอยู่ในช่วง (1-1000000)"` 
	Shelving_ID      *uint			`valid:"-"`
	Shelving         Shelving		`gorm:"references:id" valid:"-"`
	Shopping_Cart_ID *uint			`valid:"-"`
	Shopping_Cart    Shopping_Cart	`gorm:"references:id" valid:"-"`
}

type Status struct {
	gorm.Model
	Status string

	Shopping_Cart []Shopping_Cart `gorm:"foreignKey:Status_ID"`
}
