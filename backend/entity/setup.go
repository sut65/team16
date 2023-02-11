package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
		&Order{}, Shopping_Cart{}, Status{},
		&Payment_method{}, &Payment{},
		&Review_Point{}, &Type_Comment{}, &Comment{},
		&Reason{}, &Separation{},
		&Section{}, &L_Type{}, &Leave{},
		&Gender{}, &Level{}, &Member{},
		&Duty{}, &Overtime{}, &Working_time{}, &Employee_attendance{}, &Record_employee_leave{},
	)
	db = database // ห้าม comment บรรทัดนี้

// 	password1, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
// 	password2, _ := bcrypt.GenerateFromPassword([]byte("654321"), 14)
// 	password3, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)

// 	Em1 := Employee{
// 		Name:     "นายคณาการ เชิดในเมือง",
// 		Email:    "k@gmail.com",
// 		Password: string(password1),
// 		Position: "Full-Time",
// 	}
// 	db.Model(&Employee{}).Create(&Em1)

// 	Em2 := Employee{
// 		Name:     "นางญาตา ประสานวงศ์",
// 		Email:    "y@gmail.com",
// 		Password: string(password2),
// 		Position: "Part-Time",
// 	}
// 	db.Model(&Employee{}).Create(&Em2)

// 	Em3 := Employee{
// 		Name:     "นางรสนันท์ กลับเพชร",
// 		Email:    "r@gmail.com",
// 		Password: string(password3),
// 		Position: "Trainee",
// 	}
// 	db.Model(&Employee{}).Create(&Em3)

// 	//Kind
// 	Kind1 := Kind{
// 		Name: "เนื้อ",
// 	}
// 	db.Model(&Kind{}).Create(&Kind1)

// 	Kind2 := Kind{
// 		Name: "ผักและผลไม้",
// 	}
// 	db.Model(&Kind{}).Create(&Kind2)

// 	Kind3 := Kind{
// 		Name: "ผลิตภัณฑ์นม",
// 	}
// 	db.Model(&Kind{}).Create(&Kind3)

// 	//Storage
// 	Storage1 := Storage{
// 		Name: "คลังเก็บเนื้อ",
// 	}
// 	db.Model(&Storage{}).Create(&Storage1)

// 	Storage2 := Storage{
// 		Name: "คลังเก็บผักและผลไม้",
// 	}
// 	db.Model(&Storage{}).Create(&Storage2)

// 	Storage3 := Storage{
// 		Name: "คลังเก็บผลิตภัณฑ์นม",
// 	}
// 	db.Model(&Storage{}).Create(&Storage3)

// 	//Stock
// 	Stock1 := Stock{
// 		Name:     "กล้วย",
// 		Amount:   100,
// 		Price:    9.00,
// 		Employee: Em1,
// 		Kind:     Kind2,
// 		Storage:  Storage2,
// 		DateTime: time.Now(),
// 	}
// 	db.Model(&Stock{}).Create(&Stock1)

// 	Stock2 := Stock{
// 		Name:     "เนื้อหมู",
// 		Amount:   100,
// 		Price:    80.00,
// 		Employee: Em1,
// 		Kind:     Kind1,
// 		Storage:  Storage1,
// 		DateTime: time.Now(),
// 	}
// 	db.Model(&Stock{}).Create(&Stock2)

// 	Stock3 := Stock{
// 		Name:     "นม",
// 		Amount:   100,
// 		Price:    24.00,
// 		Employee: Em1,
// 		Kind:     Kind3,
// 		Storage:  Storage3,
// 		DateTime: time.Now(),
// 	}
// 	db.Model(&Stock{}).Create(&Stock3)

// 	//Label
// 	Label1 := Label{
// 		Name: "เนื้อ",
// 	}
// 	db.Model(&Label{}).Create(&Label1)

// 	Label2 := Label{
// 		Name: "ผักและผลไม้",
// 	}
// 	db.Model(&Label{}).Create(&Label2)

// 	Label3 := Label{
// 		Name: "ผลิตภัณฑ์นม",
// 	}
// 	db.Model(&Label{}).Create(&Label3)

// 	Reason1 := &Reason{
// 		Cuase: "สินค้าหมดอายุ",
// 	}
// 	db.Model(&Reason{}).Create(&Reason1)

// 	Reason2 := &Reason{
// 		Cuase: "สินค้าเสียหาย",
// 	}
// 	db.Model(&Reason{}).Create(&Reason2)

// 	ReviewP1 := &Review_Point{
// 		Point: 1,
// 	 Point_Name: "ไม่พอใจ",
// 	}
// 	db.Model(&Review_Point{}).Create(&ReviewP1)

// 	ReviewP2 := &Review_Point{
// 		Point: 2,
// 	 Point_Name: "พอใจน้อย",

// 	}
// 	db.Model(&Review_Point{}).Create(&ReviewP2)

// 	ReviewP3 := &Review_Point{
// 		Point: 3,
// 	 Point_Name: "พอใช้",

// 	}
// 	db.Model(&Review_Point{}).Create(&ReviewP3)

// 	ReviewP4 := &Review_Point{
// 		Point: 4,
// 	 Point_Name: "พอใจมาก",

// 	}
// 	db.Model(&Review_Point{}).Create(&ReviewP4)

// 	ReviewP5 := &Review_Point{
// 		Point: 5,
// 	 Point_Name: "พอใจมากที่สุด",

// 	}
// 	db.Model(&Review_Point{}).Create(&ReviewP5)

// 	TypeCom1 := &Type_Comment{
// 		Type_Com_Name: "เสนอความคิดเห็น",
// 	}
// 	db.Model(&Type_Comment{}).Create(&TypeCom1)

// 	TypeCom2 := &Type_Comment{
// 		Type_Com_Name: "ความรู้สึกหลังใช้สินค้า",
// 	}
// 	db.Model(&Type_Comment{}).Create(&TypeCom2)

// 	TypeCom3 := &Type_Comment{
// 		Type_Com_Name: "ความรู้สึกเกี่ยวกับการบริการในร้าน",
// 	}
// 	db.Model(&Type_Comment{}).Create(&TypeCom3)

// 	TypeCom4 := &Type_Comment{
// 		Type_Com_Name: "แจ้งปัญหาหรือข้อบกพร่อง",
// 	}
// 	db.Model(&Type_Comment{}).Create(&TypeCom4)

// 	//Payment_method
// 	Pay1 := Payment_method{
// 		Method: "เงินสด",
// 	}
// 	db.Model(&Payment_method{}).Create(&Pay1)

// 	Pay2 := Payment_method{
// 		Method: "Prompt Pay",
// 	}
// 	db.Model(&Payment_method{}).Create(&Pay2)

// 	//Gender
// 	Gender1 := Gender{
// 		Gender_Name: "ชาย",
// 	}
// 	db.Model(&Gender{}).Create(&Gender1)

// 	Gender2 := Gender{
// 		Gender_Name: "หญิง",
// 	}
// 	db.Model(&Gender{}).Create(&Gender2)

// 	//Level
// 	Level1 := Level{
// 		Level_Name:    "ปกติ",
// 		Level_Pay:     59,
// 		Level_Benefit: "ได้รับโปรโมชั้นขั้นพื้นฐานทั่วไป",
// 	}
// 	db.Model(&Level{}).Create(&Level1)
// 	Level2 := Level{
// 		Level_Name:    "มังสาวิรัติ",
// 		Level_Pay:     119,
// 		Level_Benefit: "ได้รับโปรโมชั่นพิเศษสำหรับสินค้าประเภทผัก ผลไม้",
// 	}
// 	db.Model(&Level{}).Create(&Level2)
// 	Level3 := Level{
// 		Level_Name:    "เนื้อล้วนๆ",
// 		Level_Pay:     119,
// 		Level_Benefit: "ได้รับโปรโมชั่นพิเศษสำหรับสินค้าประเภทเนื้อสัตว์ นม ไข้",
// 	}
// 	db.Model(&Level{}).Create(&Level3)
// 	Level4 := Level{
// 		Level_Name:    "พระเจ้า",
// 		Level_Pay:     199,
// 		Level_Benefit: "ได้รับโปรโมชั่นพิเศษสำหรับสินค้าทุกประเภท",
// 	}
// 	db.Model(&Level{}).Create(&Level4)

// 	//Member
// 	Member1 := Member{
// 		Mem_Name: "บรรเจิด เลิศเลอ",
// 		Mem_Age:  33,
// 		Mem_Tel:  "0987654321",
// 		Gender:   Gender1,
// 		Level:    Level4,
// 		Employee: Em1,
// 	}
// 	db.Model(&Member{}).Create(&Member1)

// 	//L_Type
// 	L_Type1 := L_Type{
// 		Type_Name:      "ลาป่วย",
// 		Type_Condition: "ลาตั้งแต่ 3 วันขึ้นไป ต้องมีใบรับรองแพทย์",
// 		Type_NTime:     15,
// 	}
// 	db.Model(&L_Type{}).Create(&L_Type1)

// 	L_Type2 := L_Type{
// 		Type_Name:      "ลากิจส่วนตัว",
// 		Type_Condition: "ต้องยื่นเอกสารล่วงหน้าก่อน 3 วัน",
// 		Type_NTime:     10,
// 	}
// 	db.Model(&L_Type{}).Create(&L_Type2)

// 	L_Type3 := L_Type{
// 		Type_Name:      "ลาคลอดบุตร",
// 		Type_Condition: "ต้องมีสูติบัตรของบุตร",
// 		Type_NTime:     60,
// 	}
// 	db.Model(&L_Type{}).Create(&L_Type3)

// 	L_Type4 := L_Type{
// 		Type_Name:      "ลาไปช่วยภริยาคลอดบุตร",
// 		Type_Condition: "เป็นภริยาโดยชอบด้วยกฎหมาย และยื่นใบลาให้หัวหน้าแผนกอนุญาต",
// 		Type_NTime:     15,
// 	}
// 	db.Model(&L_Type{}).Create(&L_Type4)

// 	L_Type5 := L_Type{
// 		Type_Name:      "ลาเข้ารับการคัดเลือกทหาร",
// 		Type_Condition: "ลมีหมายเรียก (สด.35) และยื่นใบลาให้หัวหน้าแผนกอนุญาต",
// 		Type_NTime:     7,
// 	}
// 	db.Model(&L_Type{}).Create(&L_Type5)
// 	L_Type6 := L_Type{
// 		Type_Name:      "ไปศึกษา ฝึกอบรม หรือดูงาน ",
// 		Type_Condition: "ยื่นใบลาให้หัวหน้าแผนกอนุญาต",
// 		Type_NTime:     90,
// 	}
// 	db.Model(&L_Type{}).Create(&L_Type6)

// 	//Section
// 	Section1 := Section{
// 		Sec_Name:   "แผนกบริการลูกค้า",
// 		Sec_Salary: 11500,
// 	}
// 	db.Model(&Section{}).Create(&Section1)

// 	Section2 := Section{
// 		Sec_Name:   "แผนกตรวจสอบสินค้า",
// 		Sec_Salary: 14900,
// 	}
// 	db.Model(&Section{}).Create(&Section2)

// 	Section3 := Section{
// 		Sec_Name:   "แผนกขนส่ง",
// 		Sec_Salary: 13700,
// 	}
// 	db.Model(&Section{}).Create(&Section3)

// 	// Discount Type
// 	Discount_Type1 := Discount_Type{
// 		Type_Name: "ลดราคาทั่วไป",
// 	}
// 	db.Model(&Discount_Type{}).Create(&Discount_Type1)
// 	Discount_Type2 := Discount_Type{
// 		Type_Name: "ลดราคาช่วงเทศกาล",
// 	}
// 	db.Model(&Discount_Type{}).Create(&Discount_Type2)
// 	Discount_Type3 := Discount_Type{
// 		Type_Name: "ลดราคาทั่วไป",
// 	}
// 	db.Model(&Discount_Type{}).Create(&Discount_Type3)

// 	//Status
// 	Status1 := Status{
// 		Status: "ยังไม่ได้ชำระ",
// 	}
// 	db.Model(&Status{}).Create(&Status1)

// 	Status2 := Status{
// 		Status: "ชำระแล้ว",
// 	}
// 	db.Model(&Status{}).Create(&Status2)

// 	// Car
// 	Car1 := Car{
// 		Car_Model:          "Ranger Double Cab Raptor",
// 		Registation_Number: "รกช อุดรธานี 965",
// 	}
// 	db.Model(&Car{}).Create(&Car1)
// 	Car2 := Car{
// 		Car_Model:          "Ranger Double Cab Wildtrak",
// 		Registation_Number: "ดชช อุดรธานี 564",
// 	}
// 	db.Model(&Car{}).Create(&Car2)
// 	Car3 := Car{
// 		Car_Model:          "Ranger Double Cab Wildtrak",
// 		Registation_Number: "สวง สุโขทัย 122",
// 	}
// 	db.Model(&Car{}).Create(&Car3)
}
