package entity

import (
	"time"
	"gorm.io/gorm"
)


type Reason struct {
	gorm.Model
	Cuase        string
	Separation   []Separation `gorm:"foreignKey:Reason_ID"`
}

type Separation struct {
	gorm.Model
	Reason_ID  *uint
	Reason      Reason
	Employee_ID *uint
	Employee    Employee
	Shelving_ID *uint
	Shelving    Shelving
	Date_Out    time.Time   //`valid:"required~Date out cannot be blank"`
	Amount      int         `valid:"required~จำนวนต้องไม่เป็นค่าว่าง"`
	Status      string		`valid:"required~สถานะต้องไม่เป็นค่าว่าง"`
}
