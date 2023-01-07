package entity

import (
	"time"

	"gorm.io/gorm"
)

type Duty struct {
	gorm.Model
	Name      string
	// Employee_attendance  []Employee_attendance  `gorm:"foreignKey:Duty_ID"`
	// record_employee_leave  []record_employee_leave  `gorm:"foreignKey:Duty_ID"`
}

type Overtime struct {
	gorm.Model
	Name      string
	// Employee_attendance  []Employee_attendance  `gorm:"foreignKey:Duty_ID"`
	// record_employee_leave  []record_employee_leave  `gorm:"foreignKey:Duty_ID"`
}

