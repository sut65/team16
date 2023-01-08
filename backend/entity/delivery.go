package entity

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Car_Model			string
	Registation_Number	string
	Delivery  []Delivery `gorm:"foreignKey:Car_ID"`
}

//ยังไม่เชื่อม FK กับ Employee และ Playmet
type Delivery struct {
	gorm.Model
	Location string
	Customer_name string
	Delivery_date time.Time

	Car_ID	*uint
	Car	Car
	Employee_ID *uint
	Employee    Employee
	Playment_ID		*uint
	// Playment	Playment
}

