package entity

import (
	"time"

	"gorm.io/gorm"
)

type Duty struct {
	gorm.Model
	Name string
	Employee_attendance  []Employee_attendance  `gorm:"foreignKey:Duty_ID"`
	Record_employee_leave  []Record_employee_leave  `gorm:"foreignKey:Duty_ID"`
}

type Overtime struct {
	gorm.Model
	Name string
	Employee_attendance  []Employee_attendance  `gorm:"foreignKey:Overtime_ID"`
	Record_employee_leave  []Record_employee_leave  `gorm:"foreignKey:Overtime_ID"`
}

type Working_time struct {
	gorm.Model
	Name    string
	WT_Time string
	Employee_attendance  []Employee_attendance  `gorm:"foreignKey:Working_time_ID"`
	Record_employee_leave  []Record_employee_leave  `gorm:"foreignKey:Working_time_ID"`
}

type Employee_attendance struct {
	gorm.Model
	Employee_ID     *uint
	Employee        Employee
	Duty_ID         *uint
	Duty            Duty
	Working_time_ID *uint
	Working_time    Working_time
	Overtime_ID     *uint
	Overtime        Overtime
	Time_IN         time.Time
	Status_ID       bool
	Number_Em       string
}

type Record_employee_leave struct {
	gorm.Model
	Employee_ID     *uint
	Employee        Employee
	Duty_ID         *uint
	Duty            Duty
	Working_time_ID *uint
	Working_time    Working_time
	Overtime_ID     *uint
	Overtime        Overtime
	Time_OUT         time.Time
	Number_Em       string
}
