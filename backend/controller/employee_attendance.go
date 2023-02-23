package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /Shopping_Cart
func CreateEmployee_attendance(c *gin.Context) {
	var Em_IN entity.Employee_attendance
	var employee entity.Employee
	var duty entity.Duty
	var working_time entity.Working_time
	var overtime entity.Overtime
	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร Employee_attendance
	if err := c.ShouldBindJSON(&Em_IN); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 9: ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", Em_IN.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเข้าสู่ระบบ"})
		return
	}
	// 10: ค้นหา Duty ด้วย id
	if tx := entity.DB().Where("id = ?", Em_IN.Duty_ID).First(&duty); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกหน้าที่ในการทำงาน"})
		return
	}
	// 11: ค้นหา working_time ด้วย id
	if tx := entity.DB().Where("id = ?", Em_IN.Working_time_ID).First(&working_time); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกเวลาในการทำงาน"})
		return
	}
	// 12 ค้นหา overtime ด้วย id
	if tx := entity.DB().Where("id = ?", Em_IN.Overtime_ID).First(&overtime); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกการทำงานล่วงเวลา"})
		return
	}
	Em_IN.Time_IN = Em_IN.Time_IN.Local()
	// 13: สร้าง Employee_attendance
	sc := entity.Employee_attendance{
		Employee:     employee,        // โยงความสัมพันธ์กับ Entity Employee
		Duty:         duty,            // โยงความสัมพันธ์กับ Entity Duty
		Working_time: working_time,    // โยงความสัมพันธ์กับ Entity Working_time
		Overtime:     overtime,        // โยงความสัมพันธ์กับ Entity Overtime
		Time_IN:      Em_IN.Time_IN,   // ตั้งค่าฟิลด์ Time_IN                
		Number_Em:    Em_IN.Number_Em, // ตั้งค่าฟิลด์ Number_Em
	}

	// ขั้นตอนการ validate ที่นำมาจาก unit test

	if _, err := govalidator.ValidateStruct(sc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 14: บันทึก
	if err := entity.DB().Create(&sc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": sc})
}

// GET /Employee_attendance
func ListEmployee_attendance(c *gin.Context) {
	var employee_attendance []entity.Employee_attendance
	if err := entity.DB().Preload("Employee").Preload("Duty").Preload("Working_time").Preload("Overtime").Raw("SELECT * FROM employee_attendances").Find(&employee_attendance).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employee_attendance})
}

// DELETE /Employee_attendance/:id
func DeleteEmployee_attendance(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Employee_attendances WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Employee_attendance
func UpdateEmployee_attendance(c *gin.Context) {
	var Em_IN entity.Employee_attendance
	id := c.Param("id")
	var employee entity.Employee
	var duty entity.Duty
	var working_time entity.Working_time
	var overtime entity.Overtime

	if err := c.ShouldBindJSON(&Em_IN); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 9: ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", Em_IN.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเข้าสู่ระบบ"})
		return
	}
	// 10: ค้นหา Duty ด้วย id
	if tx := entity.DB().Where("id = ?", Em_IN.Duty_ID).First(&duty); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกหน้าที่ในการทำงาน"})
		return
	}
	// 11: ค้นหา working_time ด้วย id
	if tx := entity.DB().Where("id = ?", Em_IN.Working_time_ID).First(&working_time); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกเวลาในการทำงาน"})
		return
	}
	// 12 ค้นหา overtime ด้วย id
	if tx := entity.DB().Where("id = ?", Em_IN.Overtime_ID).First(&overtime); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกการทำงานล่วงเวลา"})
		return
	}
	
	// 13: สร้าง Employee_attendance
	sc := entity.Employee_attendance{
		Employee:     employee,        // โยงความสัมพันธ์กับ Entity Employee
		Duty:         duty,            // โยงความสัมพันธ์กับ Entity Duty
		Working_time: working_time,    // โยงความสัมพันธ์กับ Entity Working_time
		Overtime:     overtime,        // โยงความสัมพันธ์กับ Entity Overtime
		Time_IN:      Em_IN.Time_IN,   // ตั้งค่าฟิลด์ Time_IN                 
		Number_Em:    Em_IN.Number_Em, // ตั้งค่าฟิลด์ Number_Em
	}

	// ขั้นตอนการ validate ที่นำมาจาก unit test

	if _, err := govalidator.ValidateStruct(sc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Where("id = ?", id).Updates(&sc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Em_IN})
}

// GET /Employee_attendance/:id

func GetEmployee_attendance(c *gin.Context) {

	var employee_attendance entity.Employee_attendance

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM employee_attendances WHERE id = ?", id).Scan(&employee_attendance).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": employee_attendance})

}
