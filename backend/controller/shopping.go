package controller

import (
	"net/http"

	"github.com/Team16/farm_mart/entity"
	"github.com/gin-gonic/gin"
)

// POST /Cart
func CreateCart(c *gin.Context) {

	var cart entity.Shopping_Cart
	var member entity.Member
	var employee entity.Employee
	var status entity.Status

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// 9: ค้นหา Member ด้วย id
	if cart.Member_ID != nil {
		if tx := entity.DB().Where("id = ?", cart.Member_ID).First(&member); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Member not found"})
			return
		}
	} else { 
		// ทำงานเมื่อ Member_ID เป็น nil
	}
	  
	// if tx := entity.DB().Where("id = ?", cart.Member_ID).First(&member); tx.RowsAffected == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Member not found"})
	// 	return
	// }

	// 10: ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", cart.Employee_ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 11: ค้นหา status ด้วย id
	if tx := entity.DB().Where("id = ?", cart.Status_ID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}
	// 12: สร้าง cart
	
	sc := entity.Shopping_Cart{
		Member:   	member,   // โยงความสัมพันธ์กับ Entity member
		Employee: 	employee, // โยงความสัมพันธ์กับ Entity Employee
		Status: 	status, // โยงความสัมพันธ์กับ Entity Status
	}

	// 13: บันทึก
	if err := entity.DB().Create(&sc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": cart})
}

// GET /cart/:id
func GetCart(c *gin.Context) {
	var cart entity.Shopping_Cart
	id := c.Param("id")
	if err := entity.DB().Preload("Member").Preload("Employee").Preload("Status").Raw("SELECT * FROM shopping_carts WHERE id = ?", id).Find(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// GET /shopping_cart
func ListCart(c *gin.Context) {
	var cart []entity.Shopping_Cart
	if err := entity.DB().Preload("Member").Preload("Employee").Preload("Status").Raw("SELECT * FROM shopping_carts").Find(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// DELETE /cart/:id
func DeleteCart(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM shopping_carts WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shopping_cart not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Cart
func UpdateCart(c *gin.Context) {
	var cart entity.Shopping_Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", cart.ID).First(&cart); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}

	if err := entity.DB().Save(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})
}
