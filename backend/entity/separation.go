package entity

import (
	"time"

	"gorm.io/gorm"
)


type Reason struct {
	gorm.Model
	cuase        string
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
	Date_Out    time.Time
	Amount      int
	Status      string
}
