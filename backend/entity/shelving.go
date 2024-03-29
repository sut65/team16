package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Label struct {
	gorm.Model
	Name     string
	Shelving []Shelving `gorm:"foreignKey:Label_ID"`
}

type Shelving struct {
	gorm.Model

	Employee_ID *uint    `valid:"-"`
	Employee    Employee `gorm:"references:id" valid:"-"`

	Label_ID *uint `valid:"-"`
	Label    Label `gorm:"references:id" valid:"-"`

	Stock_ID *uint `valid:"-"`
	Stock    Stock `gorm:"references:id" valid:"-"`

	Number int `valid:"required~จำนวนสินค้าต้องอยู่ในช่วง 1 - 20,range(1|20)~จำนวนสินค้าต้องอยู่ในช่วง 1 - 20"`

	Cost      float64   `valid:"required~ราคาสินค้าต้องอยู่ในช่วง 1 - 1000 บาท,range(1|1000)~ราคาสินค้าต้องอยู่ในช่วง 1 - 1000 บาท"`
	Date_Time time.Time `valid:"Past~เวลาต้องไม่เป็นอดีต,Future~เวลาต้องไม่เป็นอนาคต,required~เลือกเวลาที่ทำการบันทึก"`

	Separation []Separation `gorm:"foreignKey:Shelving_ID"`
	Order      []Order      `gorm:"foreignKey:Shelving_ID"`
	Discount   []Discount   `gorm:"foreignKey:Shelving_ID"`
}

// ฟังก์ชันที่จะใช่ในการ validation EntryTime
func init() {
	govalidator.CustomTypeTagMap.Set("Past", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now().Add(time.Minute*-2)) || t.Equal(time.Now())
		//return t.Before(time.Now())
	})

	govalidator.CustomTypeTagMap.Set("Future", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.Before(time.Now().Add(time.Minute*24)) || t.Equal(time.Now())

		// now := time.Now()
		// return now.Before(time.Time(t))
	})
}
