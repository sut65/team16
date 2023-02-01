package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Kind struct {
	gorm.Model
	Name  string
	Stock []Stock `gorm:"foreignKey:Kind_ID"`
}

type Storage struct {
	gorm.Model
	Name  string
	Stock []Stock `gorm:"foreignKey:Storage_ID"`
}

type Stock struct {
	gorm.Model
	Name        string  `valid:"required~Name cannot be blank"`
	Amount      int     `valid:"required~Amount must be in the range 1-1000,range(1|1000)~Amount must be in the range 1-1000"`
	Price       float64 `valid:"required~Price must be in the range 1-1000,range(1|1000)~Price must be in the range 1-1000"`
	Employee_ID *uint
	Employee    Employee
	Kind_ID     *uint
	Kind        Kind
	Storage_ID  *uint
	Storage     Storage
	Shelving    []Shelving `gorm:"foreignKey:Stock_ID"`
	Discount    []Discount `gorm:"foreignKey:Stock_ID"`
	DateTime    time.Time  `valid:"Past~DateTime must not be in the past,Future~DateTime must not be in the future"`
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
