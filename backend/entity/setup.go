package entity

import (
	// "time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("farm_mart.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema
	database.AutoMigrate(
		&Employee{}, &Kind{}, &Storage{}, &Stock{}, 
		&Label{}, &Shelving{}, 
		&Discount_Type{}, &Discount{}, 
		&Delivery{}, &Car{},
		&Shopping_Cart{}, &Order{}, 
		&Payment_method{}, &Payment{},
		&Review_Point{}, &Type_Comment{}, &Comment{}, 
		&Reason{}, &Separation{},
		&Section{}, &L_Type{}, &Leave{}, 
		&Gender{}, &Level{}, &Member{},
		&Duty{}, &Overtime{}, &Working_time{}, &Employee_attendance{}, &Record_employee_leave{},
	)

	password1, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password2, _ := bcrypt.GenerateFromPassword([]byte("654321"), 14)
	password3, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	db = database

	Em1 := Employee{
		Name:     "นายคณาการ เชิดในเมือง",
		Email:    "k@gmail.com",
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

	// //Kind
	// Kind1 := Kind{
	// 	Name: "Meat",
	// }
	// db.Model(&Kind{}).Create(&Kind1)

	// Kind2 := Kind{
	// 	Name: "Fresh",
	// }
	// db.Model(&Kind{}).Create(&Kind2)

	// Kind3 := Kind{
	// 	Name: "Dairy",
	// }
	// db.Model(&Kind{}).Create(&Kind3)

	// //Storage
	// Storage1 := Storage{
	// 	Name: "Storage1",
	// }
	// db.Model(&Storage{}).Create(&Storage1)

	// Storage2 := Storage{
	// 	Name: "Storage2",
	// }
	// db.Model(&Storage{}).Create(&Storage2)

	// Storage3 := Storage{
	// 	Name: "Storage3",
	// }
	// db.Model(&Storage{}).Create(&Storage3)

	//Stock
	// Stock1 := Stock{
	// 	Name:     "Banana",
	// 	Quantity: 20,
	// 	Price:    20.00,
	// 	Employee: Em1,
	// 	Kind:     Kind2,
	// 	Storage:  Storage2,
	// 	DateTime: time.Now(),
	// }
	// db.Model(&Stock{}).Create(&Stock1)

	// Stock2 := Stock{
	// 	Name:     "Pork",
	// 	Quantity: 20,
	// 	Price:    80.00,
	// 	Employee: Em1,
	// 	Kind:     Kind1,
	// 	Storage:  Storage1,
	// 	DateTime: time.Now(),
	// }
	// db.Model(&Stock{}).Create(&Stock2)

	// Stock3 := Stock{
	// 	Name:     "Milk",
	// 	Quantity: 20,
	// 	Price:    20.00,
	// 	Employee: Em1,
	// 	Kind:     Kind3,
	// 	Storage:  Storage3,
	// 	DateTime: time.Now(),
	// }
	// db.Model(&Stock{}).Create(&Stock3)

	// //Label
	// Label1 := Label{
	// 	Name: "Meat",
	// }
	// db.Model(&Label{}).Create(&Label1)

	// Label2 := Label{
	// 	Name: "Fresh",
	// }
	// db.Model(&Label{}).Create(&Label2)

	// Label3 := Label{
	// 	Name: "Dairy",
	// }
	// db.Model(&Label{}).Create(&Label3)

	// //Shelving
	// Shelving1 := Shelving{
	// 	Employee: Em1,
	// 	Label:    Label1,
	// 	Stock:    Stock2,
	// 	Quantity: 20,
	// }
	// db.Model(&Shelving{}).Create(&Shelving1)

	// Shelving2 := Shelving{
	// 	Employee: Em1,
	// 	Label:    Label2,
	// 	Stock:    Stock1,
	// 	Quantity: 20,
	// }
	// db.Model(&Shelving{}).Create(&Shelving2)

	// Shelving3 := &Shelving{
	// 	Employee: Em1,
	// 	Label:    Label3,
	// 	Stock:    Stock3,
	// 	Quantity: 20,
	// }
	// db.Model(&Shelving{}).Create(&Shelving3)

	// Reason1 := &Reason{
	// 	Cuase: "สินค้าหมดอายุ",
	// }
	// db.Model(&Reason{}).Create(&Reason1)

	// Reason2 := &Reason{
	// 	Cuase: "สินค้าเสียหาย",
	// }
	// db.Model(&Reason{}).Create(&Reason2)

	// ReviewP1 := &Review_Point{
	// 	Point: 1,
	// }
	// db.Model(&Review_Point{}).Create(&ReviewP1)

	// ReviewP2 := &Review_Point{
	// 	Point: 2,
	// }
	// db.Model(&Review_Point{}).Create(&ReviewP2)

	// ReviewP3 := &Review_Point{
	// 	Point: 3,
	// }
	// db.Model(&Review_Point{}).Create(&ReviewP3)

	// ReviewP4 := &Review_Point{
	// 	Point: 4,
	// }
	// db.Model(&Review_Point{}).Create(&ReviewP4)

	// ReviewP5 := &Review_Point{
	// 	Point: 5,
	// }
	// db.Model(&Review_Point{}).Create(&ReviewP5)

	// TypeCom1 := &Type_Comment{
	// 	Type_Com_Name: "เสนอความคิดเห็น",
	// }
	// db.Model(&Type_Comment{}).Create(&TypeCom1)

	// TypeCom2 := &Type_Comment{
	// 	Type_Com_Name: "ความรู้สึกหลังใช้สินค้า",
	// }
	// db.Model(&Type_Comment{}).Create(&TypeCom2)

	// TypeCom3 := &Type_Comment{
	// 	Type_Com_Name: "ความรู้สึกเกี่ยวกับการบริการในร้าน",
	// }
	// db.Model(&Type_Comment{}).Create(&TypeCom3)

	// TypeCom4 := &Type_Comment{
	// 	Type_Com_Name: "แจ้งปัญหาหรือข้อบกพร่อง",
	// }
	// db.Model(&Type_Comment{}).Create(&TypeCom4)

	//Payment_method
	// Pay1 := Kind{
	// 	Name: "Cash",
	// }
	// db.Model(&Payment_method{}).Create(&Pay1)

	// Pay2 := Kind{
	// 	Name: "Banking",
	// }
	// db.Model(&Payment_method{}).Create(&Pay2)

	// Pay3 := Kind{
	// 	Name: "Prompt Pay",
	// }
	// db.Model(&Payment_method{}).Create(&Pay3)

	// Pay4 := Kind{
	// 	Name: "Debit/Credit Card",
	// }
	// db.Model(&Payment_method{}).Create(&Pay4)
}


