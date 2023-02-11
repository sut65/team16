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

	Number int `valid:"required~Number must be in the range 1-20,range(1|20)~Number must be in the range 1-20"`

	Cost float64 `valid:"required~Cost must be in the range 1-1000,range(1|1000)~Cost must be in the range 1-1000"`
	Date_Time  time.Time    `valid:"Past~DateTime must not be in the past,Future~DateTime must not be in the future"`

	Separation []Separation `gorm:"foreignKey:Shelving_ID"`
	Order      []Order      `gorm:"foreignKey:Shelving_ID"`
	Discount   []Discount 	`gorm:"foreignKey:Shelving_ID"`
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
