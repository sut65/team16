package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /Payment
func CreatePayment(c *gin.Context) {

	var payment entity.Payment
	var order entity.Order
	var method entity.Payment_method
	var employee entity.Employee

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา cart ด้วย id
	if tx := entity.DB().Where("id = ?", payment.Order_ID).First(&order); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}

	// 11: ค้นหา Payment_method ด้วย id
	if tx := entity.DB().Where("Mem_Tel = ?", payment.Payment_method_ID).First(&method); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payment_method not found"})
		return
	}

	// 12: ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", payment.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 12: สร้าง Payment
	sc := entity.Payment{
		Order: order, 				// โยงความสัมพันธ์กับ Entity Order
		Payment_method:   method,   // โยงความสัมพันธ์กับ Entity Payment_method
		Employee: employee, 		// โยงความสัมพันธ์กับ Entity Employee
	}

	// 13: บันทึก
	if err := entity.DB().Create(&sc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": payment})
}

// GET /Payment/:id
func GetPayment(c *gin.Context) {
	var payment entity.Payment
	id := c.Param("id")
	if err := entity.DB().Preload("Shelving").Preload("Shopping_Cart").Raw("SELECT * FROM payments WHERE id = ?", id).Find(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// GET /payment
func ListPayment(c *gin.Context) {
	var payment []entity.Payment
	if err := entity.DB().Preload("Shelving").Preload("Shopping_Cart").Raw("SELECT * FROM payments").Find(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// DELETE /Payment/:id
func DeletePayment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM payments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Payment
func UpdatePayment(c *gin.Context) {
	var payment entity.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", payment.ID).First(&payment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment not found"})
		return
	}

	if err := entity.DB().Save(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment})
}
