package entity

import (
	"gorm.io/gorm"
)
type Gender struct{
	gorm.Model
	Gender_Name string
	Member []Member `gorm:"foreignkey:GenderID"`
}

type Level struct {
	gorm.Model
	Level_Name string
	Level_Pay int
	Level_Benefit string
	Level []Level `gorm:"foreignkey:LevelID"`
}

type Member struct {
	gorm.Model
	Mem_Name string
	Mem_Age int
	Mem_Tel string

	GenderID *uint
	Gender Gender
	LevelID *uint
	Level Level
	Employee_ID *uint
	Employee Employee
}