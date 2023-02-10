package entity

import (
	"time"
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Car_Model			string
	Registation_Number	string
	Delivery  []Delivery 	`gorm:"foreignKey:Car_ID"`
}

type Delivery struct {
	gorm.Model
	Location string			`json:"location" valid:"required~กรุณากรอกสถานที่"`
	Customer_name string	`json:"customer_name" valid:"required~กรุณากรอกชื่อลูกค้า"`
	Delivery_date time.Time	`valid:"Future~วันที่ส่งสินค้าต้องไม่เป็นวันที่ในอนาคต"`

	Car_ID	*uint			`valid:"-"`
	Car	Car					`valid:"-"`
	Employee_ID *uint		`valid:"-"`
	Employee    Employee	`valid:"-"`
	Payment_ID		*uint	`valid:"-"`
	Payment	Payment			`valid:"-"`
}
