package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /Order
func CreateOrder(c *gin.Context) {

	var order entity.Order
	var member entity.Member
	var employee entity.Employee
	var shelv entity.Shelving

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// 9: ค้นหา Member ด้วย id
	if tx := entity.DB().Where("id = ?", order.Member_ID).First(&member); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Member not found"})
		return
	}

	// 10: ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", order.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 11: ค้นหา shelv ด้วย id
	if tx := entity.DB().Where("id = ?", order.Shelving_ID).First(&shelv); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shelving not found"})
		return
	}
	// 12: สร้าง Order
	sc := entity.Order{
		Employee: 	employee, // โยงความสัมพันธ์กับ Entity Employee
		Member:   	member,   // โยงความสัมพันธ์กับ Entity shelving
		Shelving:   shelv,   // โยงความสัมพันธ์กับ Entity shelving
	}

	// 13: บันทึก
	if err := entity.DB().Create(&sc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": order})
}

// GET /Order/:id
func GetOrder(c *gin.Context) {
	var order entity.Order
	id := c.Param("id")
	if err := entity.DB().Preload("Shelving").Preload("Shopping_Cart").Raw("SELECT * FROM orders WHERE id = ?", id).Find(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}

// GET /Order
func ListOrder(c *gin.Context) {
	var order []entity.Order
	if err := entity.DB().Preload("Shelving").Preload("Shopping_Cart").Raw("SELECT * FROM orders").Find(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// DELETE /Order/:id
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM orders WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Order
func UpdateOrder(c *gin.Context) {
	var order entity.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", order.ID).First(&order); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order not found"})
		return
	}

	if err := entity.DB().Save(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}
