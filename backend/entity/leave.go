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
	Type_Name      string
	Type_Condition string
	Type_NTime     int
	Leave         []Leave `gorm:"foreignkey:L_Type_ID"`
}

type Leave struct {
	gorm.Model
	Doc_Reason string `json:"Doc_Reason" valid:"required~กรุณากรอกเหตุผลการลา / รายละเอียด"`
	Doc_DateS  time.Time `valid:"Past~วันที่เริ่มลาต้องไม่เป็นวันที่ผ่านมาแล้ว"`
	Doc_DateE  time.Time `valid:"Past~วันสิ้นสุดการลาต้องไม่เป็นอดีต"`
	Doc_Cont   string `json:"Doc_Cont" valid:"required~กรุณากรอกช่องทางการติดต่อ"`

	Section_ID  *uint `valid:"-"`
	Section     Section `valid:"-"`
	L_Type_ID   *uint `valid:"-"`
	L_Type      L_Type `valid:"-"`
	Employee_ID *uint `valid:"-"`
	Employee    Employee `valid:"-"`
}