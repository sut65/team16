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
	Name        string
	Amount      int
	Price       float64
	Employee_ID *uint
	Employee    Employee
	Kind_ID     *uint
	Kind        Kind
	Storage_ID  *uint
	Storage     Storage
	Shelving    []Shelving `gorm:"foreignKey:Stock_ID"`
	Discount    []Discount `gorm:"foreignKey:Stock_ID"`
	DateTime    time.Time
}

func init() {
	govalidator.CustomTypeTagMap.Set("past", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		now := time.Now()
		return now.After(t)
	})
	govalidator.CustomTypeTagMap.Set("future", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		now := time.Now()
		return now.Before(time.Time(t))
	})
}
