package entity

import (
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

	Number int `valid:"required~Number must be in the range 1-20,range(1|20)~Number must be in the range 1-20,~Number should not start with 0"`

	Cost float64 `valid:"required~Cost must be in the range 1-1000,range(1|1000)~Cost must be in the range 1-1000"`

	Separation []Separation `gorm:"foreignKey:Shelving_ID"`
	Order      []Order      `gorm:"foreignKey:Shelving_ID"`
}

func init() {
	govalidator.TagMap["image_valid"] = govalidator.Validator(func(str string) bool {
		return govalidator.Matches(str, "^(data:image(.+);base64,.+)$")
	})
}
