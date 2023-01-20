package entity

import (
	"time"

	"gorm.io/gorm"
)

type Section struct {
	gorm.Model
	Sec_Name   string
	Sec_Salary int
	Leave    []Leave `gorm:"foreignkey:Section_ID"`
}

type L_Type struct {
	gorm.Model
	Tpye_Name      string
	Type_Condition string
	Type_NTime     int
	Leave         []Leave `gorm:"foreignkey:L_Type_ID"`
}

type Leave struct {
	gorm.Model
	Doc_Reason string
	Doc_DateS  time.Time
	Doc_DateE  time.Time
	Doc_Cont   string

	Section_ID  *uint
	Section     Section
	L_Type_ID   *uint
	L_Type      L_Type
	Employee_ID *uint
	Employee    Employee
}
