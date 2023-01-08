package controller

import (
	"net/http"
	"time"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /Shopping_Cart
func CreateEmployee_attendance(c *gin.Context) {

	var employee_attendance entity.Employee_attendance
	var employee entity.Employee
	var duty entity.Duty
	var working_time entity.Working_time
	var overtime entity.Overtime

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร Employee_attendance
	if err := c.ShouldBindJSON(&employee_attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", employee_attendance.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 10: ค้นหา Duty ด้วย id
	if tx := entity.DB().Where("id = ?", employee_attendance.Duty_ID).First(&duty); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Duty not found"})
		return
	}

	// 11: ค้นหา working_time ด้วย id
	if tx := entity.DB().Where("id = ?", employee_attendance.Working_time_ID).First(&working_time); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "working_time not found"})
		return
	}

	// 12 ค้นหา overtime ด้วย id
	if tx := entity.DB().Where("id = ?", employee_attendance.Overtime_ID).First(&overtime); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "overtime not found"})
		return
	}

	// 13: สร้าง Employee_attendance
	sc := entity.Employee_attendance{
		Employee:     employee,                      // โยงความสัมพันธ์กับ Entity Employee
		Duty:         duty,                          // โยงความสัมพันธ์กับ Entity Duty
		Working_time: working_time,                  // โยงความสัมพันธ์กับ Entity Working_time
		Overtime:     overtime,                      // โยงความสัมพันธ์กับ Entity Overtime
		Time_IN:      time.Time{},                   // ตั้งค่าฟิลด์ Time_IN
		Status_ID:    true,                          // ตั้งค่าฟิลด์ Status_ID
		Number_Em:    employee_attendance.Number_Em, // ตั้งค่าฟิลด์ Number_Em
	}

	// 14: บันทึก
	if err := entity.DB().Create(&sc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": employee_attendance})
}

// GET /Employee_attendance
func ListEmployee_attendance(c *gin.Context) {
	var employee_attendance []entity.Employee_attendance
	if err := entity.DB().Raw("SELECT * FROM Employee_attendances").Scan(&employee_attendance).Error; err != nil {
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
	var employee_attendance entity.Employee_attendance
	if err := c.ShouldBindJSON(&employee_attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", employee_attendance.ID).First(&employee_attendance); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}

	if err := entity.DB().Save(&employee_attendance).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employee_attendance})
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