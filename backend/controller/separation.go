package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /separations
func CreateSeparation(c *gin.Context) {

	var separation entity.Separation
	var reason entity.Reason
	var employee entity.Employee
	var shelving entity.Shelving

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร separation
	if err := c.ShouldBindJSON(&separation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา reason ด้วย id
	if tx := entity.DB().Where("id = ?", separation.Reason_ID).First(&reason); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบเหตุผลที่จำหน่าย"})
		return
	}

	// 10: ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", separation.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบพนักงาน"})
		return
	}

	// 11: ค้นหา shelving ด้วย id
	if tx := entity.DB().Where("id = ?", separation.Shelving_ID).First(&shelving); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบชั้นวาง"})
		return
	}

	// 12: สร้าง Separation
	wv := entity.Separation{
		Amount:      separation.Amount,		 // ตั้งค่าฟิลด์ Amount
		Status:      separation.Status,		 // ตั้งค่าฟิลด์ Status
		Reason:      reason,                 // โยงความสัมพันธ์กับ Entity Reason
		Shelving:    shelving,               // โยงความสัมพันธ์กับ Entity Shelving
		Employee:    employee,               // โยงความสัมพันธ์กับ Entity Employee
		Date_Out:    separation.Date_Out,    // ตั้งค่าฟิลด์ Date_Out
	}

	// 13, 14, 15: ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(wv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 16: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /separation/:id
func GetSeparation(c *gin.Context) {
	var separationS entity.Separation
	id := c.Param("id")
	if err := entity.DB().Preload("Employee").Preload("Reason").Preload("Shelving").Raw("SELECT * FROM separations WHERE id = ?", id).Find(&separationS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": separationS})
}

// GET /separations
func ListSeparations(c *gin.Context) {
	var separationS []entity.Separation
	if err := entity.DB().Preload("Employee").Preload("Reason").Preload("Shelving").Raw("SELECT * FROM separations").Find(&separationS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": separationS})
}

// DELETE /separations/:id
func DeleteSeparation(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM separations WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบการจำหน่าย"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /separations
func UpdateSeparation(c *gin.Context) {

	var separationS entity.Separation
	id := c.Param("id")
	var reason entity.Reason
	var employee entity.Employee
	var shelving entity.Shelving

	if err := c.ShouldBindJSON(&separationS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", separationS.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบพนักงาน"})
		return
	}
	if tx := entity.DB().Where("id = ?", separationS.Shelving_ID).First(&shelving); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบชั้นวาง"})
		return
	}
	if tx := entity.DB().Where("id = ?", separationS.Reason_ID).First(&reason); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบเหตุผลที่จำหน่าย"})
		return
	}
	dc := entity.Separation{           
		Date_Out:       separationS.Date_Out,   // ตั้งค่าฟิลด์ Date_Out
		Employee:	    separationS.Employee,   // โยงความสัมพันธ์กับ Entity Employee          
		Shelving:	    shelving,               // โยงความสัมพันธ์กับ Entity Shelving
		Reason:	    	reason,                 // โยงความสัมพันธ์กับ Entity Reason
		Amount:	        separationS.Amount,     // ตั้งค่าฟิลด์ Amount
		Status:		    separationS.Status,     // ตั้งค่าฟิลด์ Status
	}
	if _, err := govalidator.ValidateStruct(dc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Where("id = ?", id).Updates(&dc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": separationS})


}
