package controller

import (
	"net/http"
	"time"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /Shopping_Cart
func CreateRecord_employee_leave(c *gin.Context) {

	var record_employee_leave entity.Record_employee_leave
	var employee entity.Employee
	var duty entity.Duty
	var working_time entity.Working_time
	var overtime entity.Overtime

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร Record_employee_leave
	if err := c.ShouldBindJSON(&record_employee_leave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", record_employee_leave.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 10: ค้นหา Duty ด้วย id
	if tx := entity.DB().Where("id = ?", record_employee_leave.Duty_ID).First(&duty); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Duty not found"})
		return
	}

	// 11: ค้นหา working_time ด้วย id
	if tx := entity.DB().Where("id = ?", record_employee_leave.Working_time_ID).First(&working_time); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "working_time not found"})
		return
	}

	// 12 ค้นหา overtime ด้วย id
	if tx := entity.DB().Where("id = ?", record_employee_leave.Overtime_ID).First(&overtime); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "overtime not found"})
		return
	}

	// 13: สร้าง Record_employee_leave
	sc := entity.Record_employee_leave{
		Employee:     employee,                        // โยงความสัมพันธ์กับ Entity Employee
		Duty:         duty,                            // โยงความสัมพันธ์กับ Entity Duty
		Working_time: working_time,                    // โยงความสัมพันธ์กับ Entity Working_time
		Overtime:     overtime,                        // โยงความสัมพันธ์กับ Entity Overtime
		Time_OUT:     time.Time{},                     // ตั้งค่าฟิลด์ Time_OUT
		Number_Em:    record_employee_leave.Number_Em, // ตั้งค่าฟิลด์ Number_Em
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
	c.JSON(http.StatusCreated, gin.H{"data": record_employee_leave})
}

// GET /Record_employee_leave
func ListRecord_employee_leave(c *gin.Context) {
	var record_employee_leave []entity.Record_employee_leave
	if err := entity.DB().Preload("Employee").Preload("Duty").Preload("Working_time").Preload("Overtime").Raw("SELECT * FROM record_employee_leaves").Find(&record_employee_leave).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": record_employee_leave})
}

// DELETE /Record_employee_leave/:id
func DeleteRecord_employee_leave(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM record_employee_leaves WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Record_employee_leave
func UpdateRecord_employee_leave(c *gin.Context) {
	var Em_Out entity.Record_employee_leave
	id := c.Param("id")
	var employee entity.Employee
	var duty entity.Duty
	var working_time entity.Working_time
	var overtime entity.Overtime

	if err := c.ShouldBindJSON(&Em_Out); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", Em_Out.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	// 10: ค้นหา Duty ด้วย id
	if tx := entity.DB().Where("id = ?", Em_Out.Duty_ID).First(&duty); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Duty not found"})
		return
	}
	// 11: ค้นหา working_time ด้วย id
	if tx := entity.DB().Where("id = ?", Em_Out.Working_time_ID).First(&working_time); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "working_time not found"})
		return
	}
	// 12 ค้นหา overtime ด้วย id
	if tx := entity.DB().Where("id = ?", Em_Out.Overtime_ID).First(&overtime); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "overtime not found"})
		return
	}
	Em_Out.Time_OUT = Em_Out.Time_OUT.Local()
	// 13: สร้าง Record_employee_leave
	sc := entity.Record_employee_leave{
		Employee:     employee,                      // โยงความสัมพันธ์กับ Entity Employee
		Duty:         duty,                          // โยงความสัมพันธ์กับ Entity Duty
		Working_time: working_time,                  // โยงความสัมพันธ์กับ Entity Working_time
		Overtime:     overtime,                      // โยงความสัมพันธ์กับ Entity Overtime
		Time_OUT:      Em_Out.Time_OUT,   			// ตั้งค่าฟิลด์ Time_IN                         // ตั้งค่าฟิลด์ Status_ID
		Number_Em:    Em_Out.Number_Em, 				// ตั้งค่าฟิลด์ Number_Em
	}
	if err := entity.DB().Where("id = ?", id).Updates(&sc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Em_Out})
}

// GET /Record_employee_leave/:id

func GetRecord_employee_leave(c *gin.Context) {

	var record_employee_leave entity.Record_employee_leave

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM record_employee_leaves WHERE id = ?", id).Scan(&record_employee_leave).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": record_employee_leave})

}
