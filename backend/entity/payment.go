package entity

import (
	"time"

	"gorm.io/gorm"
)

type Payment_method struct {
	gorm.Model
	Method  string
	Payment []Payment `gorm:"foreignKey:Payment_method_ID"`
}

type Payment struct {
	gorm.Model
	Time              time.Time			`valid:"Past~วันที่ห้ามเป็นอดีต,Future~วันที่ห้ามเป็นอนาคต"`
	Paytotal          float64			`valid:"required~ระบุราคา,range(1|1000000)~ยอดรวมอยู่ในช่วง (1-1000000)"` 
	Note			  string			`valid:"stringlength(0|64)~อักขระไม่เกิน 64 ตัว"`
	Shopping_Cart_ID  *uint				`valid:"-"`
	Shopping_Cart     Shopping_Cart		`gorm:"references:id" valid:"-"`
	Payment_method_ID *uint				`valid:"-"`
	Payment_method    Payment_method	`gorm:"references:id" valid:"-"`
	Employee_ID       *uint				`valid:"-"`
	Employee          Employee			`gorm:"references:id" valid:"-"`
	Delivery          []Delivery		`gorm:"foreignKey:Payment_ID"`
}
