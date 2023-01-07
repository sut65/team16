package entity

import (
	"time"
	"gorm.io/gorm"
)

type Section struct {
	gorm.Model
	Sec_Name string
	Sec_Salary int
	Section []Section `gorm:"foreignkey:SectionID"`
}

type L_Type struct {
	gorm.Model
	Tpye_Name string
	Type_Condition string
	Type_NTime int
	L_Type []L_Type `gorm:"foreignkey:L_TypeID"`
}

type Leave struct {
	gorm.Model
	Doc_Reason string
	Doc_DateS time.Time
	Doc_DateE time.Time
	Doc_Cont string

	SectionID *uint
	Section Section
	L_TypeID *uint
	L_Type L_Type
	Employee_ID *uint
	Employee Employee
}