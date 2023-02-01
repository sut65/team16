package entity

import (
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	Gender_Name string
	Member      []Member `gorm:"foreignkey:Gender_ID"`
}

type Level struct {
	gorm.Model
	Level_Name    string
	Level_Pay     int
	Level_Benefit string
	Member        []Member `gorm:"foreignkey:Level_ID"`
}

type Member struct {
	gorm.Model
	Mem_Name string `json:"Mem_Name" valid:"required~กรุณากรอกชื่อ - นามสกุล"`
	Mem_Age  int `json:"Mem_Age" valid:"range(15|100)~โปรดระบุอายุที่มากกว่า 15 ปีขึ้นไป"`
	Mem_Tel  string `gorm:"uniqueIndex" json:"Mem_Tel" valid:"required~กรุณากรอกเบอร์มือถือ"`

	Gender_ID   *uint
	Gender      Gender
	Level_ID    *uint
	Level       Level
	Employee_ID *uint
	Employee    Employee

	Shopping_Cart []Shopping_Cart `gorm:"foreignkey:Member_ID"`
}

