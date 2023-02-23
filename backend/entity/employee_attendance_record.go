package entity

import (
	"time"

	"gorm.io/gorm"
)

type Duty struct {
	gorm.Model
	Name                  string
	Employee_attendance   []Employee_attendance   `gorm:"foreignKey:Duty_ID"`
	Record_employee_leave []Record_employee_leave `gorm:"foreignKey:Duty_ID"`
}

type Overtime struct {
	gorm.Model
	Name                  string
	Employee_attendance   []Employee_attendance   `gorm:"foreignKey:Overtime_ID"`
	Record_employee_leave []Record_employee_leave `gorm:"foreignKey:Overtime_ID"`
}

type Working_time struct {
	gorm.Model
	Name                  string
	WT_Time               string
	Employee_attendance   []Employee_attendance   `gorm:"foreignKey:Working_time_ID"`
	Record_employee_leave []Record_employee_leave `gorm:"foreignKey:Working_time_ID"`
}

type Employee_attendance struct {
	gorm.Model
	Employee_ID     *uint        `valid:"-"`
	Employee        Employee     `valid:"-"`
	Duty_ID         *uint        `valid:"-"`
	Duty            Duty         `valid:"-"`
	Working_time_ID *uint        `valid:"-"`
	Working_time    Working_time `valid:"-"`
	Overtime_ID     *uint        `valid:"-"`
	Overtime        Overtime     `valid:"-"`
	Time_IN         time.Time    `valid:"-"`
	Number_Em       string       `valid:"required~โปรดใส่เบอร์โทร,minstringlength(10)~เบอร์โทรควรมีมากกว่า 9 ตัว,maxstringlength(10)~เบอร์โทรควรมีน้อยกว่า 11 ตัว"`
}

type Record_employee_leave struct {
	gorm.Model
	Employee_ID     *uint        `valid:"-"`
	Employee        Employee     `valid:"-"`
	Duty_ID         *uint        `valid:"-"`
	Duty            Duty         `valid:"-"`
	Working_time_ID *uint        `valid:"-"`
	Working_time    Working_time `valid:"-"`
	Overtime_ID     *uint        `valid:"-"`
	Overtime        Overtime     `valid:"-"`
	Time_OUT        time.Time    `valid:"-"`
	Number_Em       string       `valid:"required~โปรดใส่เบอร์โทร,minstringlength(10)~เบอร์โทรควรมีมากกว่า 9 ตัว,maxstringlength(10)~เบอร์โทรควรมีน้อยกว่า 11 ตัว"`
}
