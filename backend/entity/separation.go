package entity

import (
	"time"

	"gorm.io/gorm"
)

type Reason struct {
	gorm.Model
	Cuase      string
	Separation []Separation `gorm:"foreignKey:Reason_ID"`
}

type Separation struct {
	gorm.Model
	Reason_ID   *uint
	Reason      Reason     `valid:"-"`
	Employee_ID *uint  
	Employee    Employee   `valid:"-"`
	Shelving_ID *uint
	Shelving    Shelving   `valid:"-"`
	Date_Out    time.Time  `valid:"Past~วันที่ต้องไม่เป็นอดีต"`
	Amount      int        `valid:"required~จำนวนต้องไม่เป็นค่าว่าง, range(0|9223372036854775807)~กรุณากรอกจำนวนเต็มบวกเท่านั้น"`
	Status      string     `valid:"required~สถานะต้องไม่เป็นค่าว่าง, in(+|-)~กรุณากรอก + หรือ - เท่านั้น"`
}

