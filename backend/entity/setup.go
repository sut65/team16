package entity

import (
	"time"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("se-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(&Employee{}, &Kind{}, &Invetory{}, &Stock{})

	db = database

	Em1 := &Employee{
		FirstName: "First",
		LastName:  "Employ1",
		Email:     "Employee1@gmail.com",
		Password:  "Em1Password",
		Position:  "Full-Time",
	}
	db.Model(&Employee{}).Create(&Em1)

	Em2 := &Employee{
		FirstName: "First",
		LastName:  "Employ2",
		Email:     "Employee2@gmail.com",
		Password:  "Em2Password",
		Position:  "Part-Time",
	}
	db.Model(&Employee{}).Create(&Em2)

	Em3 := &Employee{
		FirstName: "First",
		LastName:  "Employ3",
		Email:     "Employee3@gmail.com",
		Password:  "Em3Password",
		Position:  "Casual",
	}
	db.Model(&Employee{}).Create(&Em3)

	//Kind
	Kind1 := &Kind{
		Name: "Meat",
	}
	db.Model(&Kind{}).Create(&Kind1)

	Kind2 := &Kind{
		Name: "Fresh",
	}
	db.Model(&Kind{}).Create(&Kind2)

	Kind3 := &Kind{
		Name: "Dairy",
	}
	db.Model(&Kind{}).Create(&Kind3)

	//Storage
	Storage1 := &Storage{
		Name: "Storage1",
	}
	db.Model(&Storage{}).Create(&Storage1)

	Storage2 := &Storage{
		Name: "Storage2",
	}
	db.Model(&Storage{}).Create(&Storage2)

	Storage3 := &Storage{
		Name: "Storage3",
	}
	db.Model(&Storage{}).Create(&Storage3)

	//Invetory
	Invetory1 := &Invetory{
		Name:     "Banana",
		Quantity: 20,
		Price:    20,
		Kind:     *Kind2,
		Storage:  *Storage2,
	}
	db.Model(&Invetory{}).Create(&Invetory1)

	Invetory2 := &Invetory{
		Name:     "Pork",
		Quantity: 20,
		Price:    80,
		Kind:     *Kind1,
		Storage:  *Storage1,
	}
	db.Model(&Invetory{}).Create(&Invetory2)

	Invetory3 := &Invetory{
		Name:     "Milk",
		Quantity: 20,
		Price:    20,
		Kind:     *Kind3,
		Storage:  *Storage3,
	}
	db.Model(&Invetory{}).Create(&Invetory3)

	//Stock
	Stock1 := &Stock{
		Employee: *Em1,
		Kind:     *Kind1,
		Storage:  *Storage1,
		Invetory: *Invetory2,

		DateTime: time.Now(),
	}
	db.Model(&Stock{}).Create(&Stock1)

	Stock2 := &Stock{
		Employee: *Em1,
		Kind:     *Kind2,
		Storage:  *Storage2,
		Invetory: *Invetory1,

		DateTime: time.Now(),
	}
	db.Model(&Stock{}).Create(&Stock2)

	Stock3 := &Stock{
		Employee: *Em1,
		Kind:     *Kind3,
		Storage:  *Storage3,
		Invetory: *Invetory3,

		DateTime: time.Now(),
	}
	db.Model(&Stock{}).Create(&Stock3)

}
