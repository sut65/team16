package entity

import (
	"time"

	"gorm.io/gorm"
)

type Discount struct {
	gorm.Model
	Discount_Price float64
	Discount_s     time.Time
	Discount_e     time.Time

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
