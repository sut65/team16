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
	Member         []Member `gorm:"foreignkey:Level_ID"`
}

type Member struct {
	gorm.Model
	Mem_Name string
	Mem_Age  int
	Mem_Tel  string	

	Gender_ID   *uint
	Gender      Gender
	Level_ID    *uint
	Level       Level
	Employee_ID *uint
	Employee    Employee

	Shopping_Cart	[]Shopping_Cart `gorm:"foreignkey:Member_ID"`

}
