package entity

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Car_Model			string
	Registation_Number	string
}

//ยังไม่เสร็จต้องทำต่อ
type Delivery struct {
	gorm.Model
	Location string
	Customer_name string
	Delivery_date time.Time

	
	Employee_ID *uint
	Employee    Employee
}

