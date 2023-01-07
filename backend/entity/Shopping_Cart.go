package entity

import (
	"gorm.io/gorm"
)

type Shopping_Cart struct {
	gorm.Model
	Total float64
	Oder  []Oder `gorm:"foreignKey:Kind_ID"`
}

type Oder struct {
	gorm.Model
	Quantity    int
	Price       int
	Employee_ID *uint
	Employee    Employee
	Payment  []Payment `gorm:"foreignKey:Oder_ID"`

	//Mem_tal    `db:"member_tal"`
	//Member     Member
}
