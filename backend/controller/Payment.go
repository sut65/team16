package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /Payment
func CreatePayment(c *gin.Context) {

	var payment entity.Payment
	var cart entity.Shopping_Cart
	var method entity.Payment_method
	var employee entity.Employee

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา cart ด้วย id
	if tx := entity.DB().Where("id = ?", payment.Shopping_Cart_ID).First(&cart); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}

	// 11: ค้นหา Payment_method ด้วย id
	if tx := entity.DB().Where("id = ?", payment.Payment_method_ID).First(&method); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกช่องทางการชำระ"})
		return
	}

	// 12: ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", payment.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 12: สร้าง Payment
	sc := entity.Payment{
		Paytotal:       payment.Paytotal,
		Note: 			payment.Note,
		Time:           payment.Time,
		Shopping_Cart:  cart,     // โยงความสัมพันธ์กับ Entity Shopping_Cart
		Payment_method: method,   // โยงความสัมพันธ์กับ Entity Payment_method
		Employee:       employee, // โยงความสัมพันธ์กับ Entity Employee
	}
	payment.Time = payment.Time.Local()
	if _, err := govalidator.ValidateStruct(sc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
	if err := entity.DB().Preload("Payment_method").Preload("Shopping_Cart").Preload("Employee").Raw("SELECT * FROM payments WHERE id = ?", id).Find(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// GET /payment
func ListPayment(c *gin.Context) {
	var payment []entity.Payment
	if err := entity.DB().Preload("Payment_method").Preload("Shopping_Cart").Preload("Employee").Raw("SELECT * FROM payments").Find(&payment).Error; err != nil {
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
	var cart entity.Shopping_Cart
	var method entity.Payment_method
	var employee entity.Employee
	id := c.Param("id")

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", payment.Shopping_Cart_ID).First(&cart); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Shopping_Cart_ID not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", payment.Payment_method_ID).First(&method); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบช่องทางการชำระ"})
		return
	}
	if tx := entity.DB().Where("id = ?", payment.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found"})
		return
	}

	sc := entity.Payment{
		Paytotal:       payment.Paytotal,
		Time:           payment.Time,
		Note: 			payment.Note,
		Shopping_Cart:  cart,     
		Payment_method: method,   
		Employee:       employee, 
	}
	payment.Time = payment.Time.Local()
	if _, err := govalidator.ValidateStruct(sc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := entity.DB().Where("id = ?", id).Updates(&sc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment})
}
