package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("stock_farm_mart.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(&Employee{}, &Kind{}, &Inventory{}, &Stock{}, &Label{}, &Product{}, &Shelving{})

	db = database

	database.AutoMigrate(&Employee{}, &Kind{}, &Inventory{}, &Stock{})

	password1, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password2, _ := bcrypt.GenerateFromPassword([]byte("654321"), 14)
	password3, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	Em1 := Employee{
		Name:     "นายเมธัส ภาคภูมิพงศ์",
		Email:    "m@gmail.com",
		Password: string(password1),
		Position: "Full-Time",
	}
	db.Model(&Employee{}).Create(&Em1)

	Em2 := Employee{
		Name:     "นางญาตา ประสานวงศ์",
		Email:    "y@gmail.com",
		Password: string(password2),
		Position: "Part-Time",
	}
	db.Model(&Employee{}).Create(&Em2)

	Em3 := Employee{
		Name:     "นางรสนันท์ กลับเพชร",
		Email:    "r@gmail.com",
		Password: string(password3),
		Position: "Trainee",
	}
	db.Model(&Employee{}).Create(&Em3)

	//Kind
	Kind1 := Kind{
		Name: "Meat",
	}
	db.Model(&Kind{}).Create(&Kind1)

	Kind2 := Kind{
		Name: "Fresh",
	}
	db.Model(&Kind{}).Create(&Kind2)

	Kind3 := Kind{
		Name: "Dairy",
	}
	db.Model(&Kind{}).Create(&Kind3)

	//Storage
	Storage1 := Storage{
		Name: "Storage1",
	}
	db.Model(&Storage{}).Create(&Storage1)

	Storage2 := Storage{
		Name: "Storage2",
	}
	db.Model(&Storage{}).Create(&Storage2)

	Storage3 := Storage{
		Name: "Storage3",
	}
	db.Model(&Storage{}).Create(&Storage3)

	//Inventory
	Inventory1 := Inventory{
		Name:     "Banana",
		Quantity: 20,
		Price:    20,
		Employee: Em1,
		Kind:     Kind2,
		Storage:  Storage2,
	}
	db.Model(&Inventory{}).Create(&Inventory1)

	Inventory2 := Inventory{
		Name:     "Pork",
		Quantity: 20,
		Price:    80,
		Employee: Em1,
		Kind:     Kind1,
		Storage:  Storage1,
	}
	db.Model(&Inventory{}).Create(&Inventory2)

	Inventory3 := Inventory{
		Name:     "Milk",
		Quantity: 20,
		Price:    20,
		Employee: Em1,
		Kind:     Kind3,
		Storage:  Storage3,
	}
	db.Model(&Inventory{}).Create(&Inventory3)

	//Stock
	Stock1 := Stock{
		Employee:  Em1,
		Kind:      Kind1,
		Storage:   Storage1,
		Inventory: Inventory2,

		DateTime: time.Now(),
	}
	db.Model(&Stock{}).Create(&Stock1)

	Stock2 := Stock{
		Employee:  Em1,
		Kind:      Kind2,
		Storage:   Storage2,
		Inventory: Inventory1,

		DateTime: time.Now(),
	}
	db.Model(&Stock{}).Create(&Stock2)

	Stock3 := Stock{
		Employee:  Em1,
		Kind:      Kind3,
		Storage:   Storage3,
		Inventory: Inventory3,

		DateTime: time.Now(),
	}
	db.Model(&Stock{}).Create(&Stock3)

}
