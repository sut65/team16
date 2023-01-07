package entity

import (
	"gorm.io/gorm"
)

type Shopping_Cart struct {
	gorm.Model
	Total		float64
	Employee_ID *uint
	Employee    Employee
	Oder  		[]Oder 		`gorm:"foreignKey:Shopping_Cart_ID"`
	Payment  	[]Payment 	`gorm:"foreignKey:Shopping_Cart_ID"`
	Member_Tal  string
	Member		Member 

}

type Oder struct {
	gorm.Model
	Quantity    		int
	Shopping_Cart_ID	*uint
	Shopping_Cart		Shopping_Cart
}
