package entity

import (
	"time"

	"gorm.io/gorm"
)

type Discount struct {
	gorm.Model
	Discount_Price float64		`valid:"required~กรุณากรอกราคาที่ลด,range(1|1000)~กรุณากรอกราคาที่ลดอยู่ในช่วง 1-1000"`
	Discount_s     time.Time	`valid:"Past~วันที่เริ่มลดราคาต้องไม่เป็นวันที่ผ่านมาแล้ว, Future~วันที่เริ่มลดราคาต้องไม่เป็นวันที่ในอนาคต"`
	Discount_e     time.Time	`valid:"Past~วันที่สิ้นสุดการลดราคาต้องไม่เป็นวันที่ผ่านมาแล้ว"`

	Discount_Type_ID *uint 			`valid:"-"`
	Discount_Type    Discount_Type 	`valid:"-"`
	Employee_ID      *uint 			`valid:"-"`
	Employee         Employee 		`valid:"-"`
	Stock_ID     *uint 				`valid:"-"`
	Stock        Stock 				`valid:"-"`
}

type Discount_Type struct {
	gorm.Model
	Type_Name string
	Discount  []Discount `gorm:"foreignKey:Discount_Type_ID"`
}
